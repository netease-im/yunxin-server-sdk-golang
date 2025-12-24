
## 其他

* 关于traceId
* 关于sdk的调度
* 关于超时
* 关于代理

### 关于traceId

* 默认每次请求会自动随机生产一个traceId,如果想要指定,则可以请求之前调用如下方法 `YunxinTraceId.Set(xxx)`
* 特别需要注意的,务必每次请求选择不同的traceId,因为云信服务器会根据traceId做请求的去重

```go
package main

import (
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/account"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).Build()

	// 创建服务
	services := im.NewYunxinV1ApiServices(client)

	// 创建请求
	request := &account.CreateAccountRequestV1{
		Accid: "zhangsan",
	}

	// 设置本次请求的traceId
	core.YunxinTraceId.Set("your_trace_id")

	// 执行请求
	result, err := services.GetAccountService().CreateAccount(request)
	if err != nil {
		fmt.Printf("register error: %v\n", err)
		return
	}

	if result.IsSuccess() {
		response := result.GetResponse()
		fmt.Printf("register success, accid=%s, token=%s, traceId=%s\n",
			response.Accid, response.Token, result.GetTraceId())
	} else {
		fmt.Printf("register fail, code=%d, msg=%s, traceId=%s\n",
			result.GetCode(), result.GetMsg(), result.GetTraceId())
	}
}
```


### 关于sdk的调度

* sdk初始化时会通过云信的调度服务器下发实际使用的api域名
* 调度域名默认使用全球化加速的域名
* 如果希望限制调度域名的dns解析结果到特定地区,可以初始化时增加region参数配置,如下

```go
package main

import (
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端,指定region
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Region("cn-east"). // 设置区域
		Build()

	// 使用client进行API调用
	_ = client
}
```


### 关于超时

* 默认超时时间是5s,如果希望修改,可以初始化时指定

```go
package main

import (
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"

	// 创建客户端,设置超时时间为10秒
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(10000 * time.Millisecond).
		Build()

	// 使用client进行API调用
	_ = client
}
```


### 关于代理

* 如果你使用了自定义的http-client,那么根据自己的http-client的规范来设置即可
* 如果你使用了sdk的默认http-client,你可以设置Proxy参数

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	coreHttp "github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建代理配置
	proxy := &coreHttp.Proxy{
		Addr: "10.0.0.0:1080",
	}

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Proxy(proxy).
		Build()

	// 执行自定义请求
	uri := "https://api.netease.im/nimserver/user/create.action"
	var header map[string]string
	var query map[string]string
	body := `{"xx":"xx"}`

	response, err := client.ExecuteJson(http.MethodPost, uri, header, query, body)
	if err != nil {
		fmt.Printf("request failed: %v\n", err)
		return
	}

	fmt.Println(response)
}
```


## 自定义Http拦截器

* 你可以自定义拦截器

```go
package main

import (
	"fmt"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
)

// 自定义拦截器
type customInterceptor struct{}

func (c *customInterceptor) Before(request *http.ApiRequest) {
	fmt.Println("before")
}

func (c *customInterceptor) After(request *http.ApiRequest, response *http.ApiResponse) {
	fmt.Println("after")
}

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建拦截器
	interceptor := &customInterceptor{}

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Interceptor(interceptor).
		Build()

	// 使用client进行API调用
	_ = client
}
```


## 自定义Http Client

* 你可以使用自定义的http client
* 自定义http client需要实现HttpClient接口

```go
package main

import (
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
)

// 自定义HttpClient
type customHttpClient struct{}

func (c *customHttpClient) Execute(request *http.HttpRequest) (*http.HttpResponse, error) {
	// 实现自定义的HTTP请求逻辑
	return nil, nil
}

func (c *customHttpClient) Close() error {
	// 实现资源清理逻辑
	return nil
}

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建自定义HttpClient
	httpClient := &customHttpClient{}

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		HttpClient(httpClient).
		Build()

	// 使用client进行API调用
	_ = client
}
```


## 使用底层HTTP客户端

* 如果封装的服务方法不满足你的要求,可以直接使用底层HTTP客户端
* 调用ExecuteJson或者ExecuteForm即可(json和form两种方式)

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
)

func main() {
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 执行自定义请求
	uri := "https://api.netease.im/nimserver/user/create.action"
	var header map[string]string
	var query map[string]string
	body := `{"xx":"xx"}`

	response, err := client.ExecuteJson(http.MethodPost, uri, header, query, body)
	if err != nil {
		fmt.Printf("request failed: %v\n", err)
		return
	}

	fmt.Println(response)
}
```
