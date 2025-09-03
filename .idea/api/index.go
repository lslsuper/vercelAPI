// 文件路径: api/index.go

// Vercel 要求 Go serverless 函数的包名必须是 handler
package handler

import (
	"fmt"
	"net/http"
)

// Vercel 会自动寻找一个名为 Handler 的公开函数 (首字母大写)
// 它的函数签名必须是 (http.ResponseWriter, *http.Request)
// 这和 Go 标准库的 http.HandleFunc 签名完全一样，非常方便
func Handler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，告诉浏览器返回的是纯文本
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// 设置 HTTP 状态码为 200 OK
	w.WriteHeader(http.StatusOK)

	// 使用 fmt.Fprintf 向响应体写入我们的 "Hello World" 消息
	fmt.Fprintf(w, "Hello from Go on Vercel!")
}