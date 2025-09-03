// 文件路径: api/index.go

package stock

import (
	"fmt"
	"io/ioutil" // 用来读取 HTTP 响应体
	"net/http"  // Go 的 HTTP 核心库
)

// Handler 函数，Vercel 会调用它
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. 从用户请求的 URL 中获取股票代码(ticker)参数
	// 例如: /api/stock?ticker=600519, 我们要拿到 "600519"
	ticker := r.URL.Query().Get("ticker")

	// 2. 参数校验：如果用户没有提供 ticker，就返回错误信息
	if ticker == "" {
		// http.Error 是一个快捷函数，用来返回错误状态码和消息
		http.Error(w, "Error: 'ticker' query parameter is required.", http.StatusBadRequest) // 400 Bad Request
		return                                                                               // 终止函数执行
	}

	// 3. 构建我们要请求的外部 API 的 URL
	// 使用 fmt.Sprintf 来安全地拼接字符串
	// 注意：为了安全和演示，token 我们暂时硬编码为 "demo"
	externalAPIURL := fmt.Sprintf("https://www.tsanghi.com/api/fin/stock/XSHG/daily?token=996df0c0fea34fb691950b49e6a49c54&ticker=%s", ticker)

	// 4. 发起 HTTP GET 请求到外部 API
	resp, err := http.Get(externalAPIURL)
	if err != nil {
		// 如果请求失败（比如网络问题或对方服务器宕机），返回服务器内部错误
		http.Error(w, "Error: Failed to fetch data from external API.", http.StatusInternalServerError) // 500 Internal Server Error
		return
	}
	// defer 语句确保在函数结束时一定会执行 Body.Close()，这是 Go 的一个重要实践，防止资源泄露
	defer resp.Body.Close()

	// 5. 读取外部 API 返回的响应体数据
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 如果读取响应体失败，同样返回服务器内部错误
		http.Error(w, "Error: Failed to read response body.", http.StatusInternalServerError)
		return
	}

	// 6. 将获取到的数据返回给我们的用户
	// 首先，设置响应头，告诉浏览器我们返回的是 JSON 格式的数据
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 设置 HTTP 状态码为 200 OK
	w.WriteHeader(http.StatusOK)
	// 将从外部 API 获取到的原始数据（bodyBytes）直接写入响应
	w.Write(bodyBytes)
}
