package main

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
)

func main() {
	// 初始化配置
	appkey := "xxx"
	appsecret := "yyy"
	timeoutMillis := 5000

	fmt.Println("=== 云信 Meeting 自定义API调用测试 ===")

	// 创建 HTTP 客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizCUSTOM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Endpoint("https://rtc-ai.yunxinapi.com").
		Build()

	// 构建请求参数
	request := make(map[string]interface{})
	request["cname"] = "xx"
	request["taskType"] = 7
	// 可以添加其他参数

	// 将请求转换为JSON字符串
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("请求JSON序列化失败: %v\n", err)
		return
	}

	// 执行JSON API请求
	response, err := client.ExecuteJson(http.POST, "/v1/api/task/create", nil, nil, string(requestJSON))
	if err != nil {
		fmt.Printf("API调用失败: %v\n", err)
		return
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Printf("响应JSON序列化失败: %v\n", err)
		fmt.Printf("原始响应: %+v\n", response)
	} else {
		fmt.Printf("响应结果:\n%s\n", string(responseJSON))
	}

	// 关闭客户端
	client.Shutdown()
	fmt.Println("\n=== 测试完成 ===")
}
