package response

import (
	"net/http"
	"runtime"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Body struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    any    `json:"data"`
}

func JSON(w http.ResponseWriter, data any, err error) {
	if err != nil {
		_, _, line, _ := runtime.Caller(3)
		httpx.OkJson(w, Body{
			Errcode: line,
			Errmsg:  extractErrorMessage(err),
		})
	} else {
		httpx.OkJson(w, Body{
			Errcode: 0,
			Errmsg:  "",
			Data:    data,
		})
	}
}

func OkJSON(w http.ResponseWriter, data Body) {
	httpx.OkJson(w, data)
}

// extractErrorMessage 从错误中提取友好的错误信息
func extractErrorMessage(err error) string {
	// 尝试转换为gRPC状态错误
	if s, ok := status.FromError(err); ok {
		return s.Message() // 直接获取gRPC错误的消息部分
	}

	// 如果不是gRPC错误，返回原始错误信息
	return err.Error()
}
