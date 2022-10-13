package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
)

// register global middleware
func Register(router *gin.Engine) *gin.Engine {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	router.Use(timeout.Timeout(
		timeout.WithTimeout(env.ServerSetting.ApiTimeout*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout),                   // optional
		timeout.WithDefaultMsg(`{"code": 408, "msg":"http: Handler timeout"}`), // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		}))) // optional

	return router
}
