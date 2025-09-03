// 文件路径: api/stock/index.go

package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url" // 引入 url 包
)

// Handler Vercel的入口函数
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取所有查询参数
	queryParams := r.URL.Query()

	// 2. 校验必选参数
	token := queryParams.Get("token")
	exchangeCode := queryParams.Get("exchange_code")
	ticker := queryParams.Get("ticker")

	if token == "" || exchangeCode == "" || ticker == "" {
		errorMsg := "Missing required parameters. 'token', 'exchange_code', and 'ticker' are all required."
		http.Error(w, errorMsg, http.StatusBadRequest) // 400 Bad Request
		return
	}

	// 3. 动态构建目标 API 的基础 URL
	// 注意路径中的 %s 会被 exchangeCode 替换
	// 我们对 exchangeCode 进行 URL 编码，防止特殊字符注入
	baseURL := fmt.Sprintf(
		"https://www.tsanghi.com/api/fin/stock/%s/daily",
		url.PathEscape(exchangeCode),
	)

	// 4. ✨ 核心步骤：直接使用原始查询字符串进行参数透传
	// r.URL.RawQuery 包含了URL中 '?' 之后的所有内容，例如 "token=demo&ticker=600519&limit=10"
	// 这样我们就无需关心有哪些可选参数，实现了完美的代理
	externalAPIURL := fmt.Sprintf("%s?%s", baseURL, r.URL.RawQuery)
	fmt.Println("Forwarding request to:", externalAPIURL) // 在Vercel日志中打印请求地址，方便调试

	// 5. 发起请求到目标 API
	resp, err := http.Get(externalAPIURL)
	if err != nil {
		http.Error(w, "Failed to fetch data from external API.", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 6. ✨ 响应透传：将目标 API 的响应头、状态码和内容原样返回
	// 复制所有响应头 (特别是 Content-Type, Content-Length 等)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 写入目标 API 的状态码 (例如 200 OK, 404 Not Found 等)
	w.WriteHeader(resp.StatusCode)

	// 将目标 API 的响应体内容直接复制给我们的响应
	io.Copy(w, resp.Body)
}
