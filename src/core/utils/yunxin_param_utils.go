package utils

import (
	"reflect"
	"strconv"
)

// YunxinParamUtils 云信参数工具类
type YunxinParamUtils struct{}

// Convert 将对象转换为参数Map
func Convert(obj interface{}) map[string]string {
	if obj == nil {
		return make(map[string]string)
	}

	paramMap := make(map[string]string)

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	// 如果是指针，获取其指向的值
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return paramMap
		}
		v = v.Elem()
		t = t.Elem()
	}

	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 检查字段是否可导出（首字母大写）
		if !field.CanInterface() {
			continue
		}

		// 获取JSON标签作为参数名
		paramName := ""
		if jsonTag := fieldType.Tag.Get("json"); jsonTag != "" {
			// 解析json标签，取第一个值（忽略omitempty等选项）
			if commaIndex := findComma(jsonTag); commaIndex != -1 {
				paramName = jsonTag[:commaIndex]
			} else {
				paramName = jsonTag
			}
		} else {
			// 如果没有json标签，使用字段名
			paramName = fieldType.Name
		}

		// 跳空值和omitempty字段
		if field.IsZero() {
			continue
		}

		// 转换字段值为字符串
		switch field.Kind() {
		case reflect.String:
			if str := field.String(); str != "" {
				paramMap[paramName] = str
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if num := field.Int(); num != 0 {
				paramMap[paramName] = strconv.FormatInt(num, 10)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if num := field.Uint(); num != 0 {
				paramMap[paramName] = strconv.FormatUint(num, 10)
			}
		case reflect.Float32, reflect.Float64:
			if num := field.Float(); num != 0 {
				paramMap[paramName] = strconv.FormatFloat(num, 'f', -1, 64)
			}
		case reflect.Bool:
			paramMap[paramName] = strconv.FormatBool(field.Bool())
		case reflect.Ptr:
			if !field.IsNil() {
				// 处理指针类型
				ptrValue := field.Elem()
				switch ptrValue.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					paramMap[paramName] = strconv.FormatInt(ptrValue.Int(), 10)
				case reflect.String:
					paramMap[paramName] = ptrValue.String()
				case reflect.Bool:
					paramMap[paramName] = strconv.FormatBool(ptrValue.Bool())
				}
			}
		case reflect.Slice:
			if field.Len() > 0 {
				// 对于slice类型，可以转换为JSON字符串
				paramMap[paramName] = field.String()
			}
		}
	}

	return paramMap
}

// findComma 查找逗号位置
func findComma(s string) int {
	for i, r := range s {
		if r == ',' {
			return i
		}
	}
	return -1
}
