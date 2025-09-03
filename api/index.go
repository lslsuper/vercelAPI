// 文件路径: api/index.go
package handler

import (
	"fmt"
	"net/http"
	"time" // 引入 time 包
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// 原来的代码:
	// fmt.Fprintf(w, "Hello from Go on Vercel!")

	// ✨ 新的代码，我们加上当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("Go Vercel API is updated! Current server time is: %s", currentTime)

	fmt.Fprintf(w, message)
}
