package http

import (
	"net/url"
	"strings"
)

// ParamBuilder 参数构建器
type ParamBuilder struct {
	params map[string]string
}

// NewParamBuilder 创建新的参数构建器
func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{
		params: make(map[string]string),
	}
}

// AddParam 添加参数
func (p *ParamBuilder) AddParam(key interface{}, value interface{}) *ParamBuilder {
	keyStr := ""
	valueStr := ""

	if key != nil {
		keyStr = key.(string)
	}
	if value != nil {
		switch v := value.(type) {
		case string:
			valueStr = v
		default:
			valueStr = v.(string)
		}
	}

	p.params[keyStr] = valueStr
	return p
}

// Build 构建URL编码的参数字符串
func (p *ParamBuilder) Build() string {
	var parts []string
	for key, value := range p.params {
		encodedKey := url.QueryEscape(key)
		encodedValue := url.QueryEscape(value)
		parts = append(parts, encodedKey+"="+encodedValue)
	}
	return strings.Join(parts, "&")
}
