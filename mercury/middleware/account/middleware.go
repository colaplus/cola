package account

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		// 可以设置一些公共参数
		context.Set("example", "12345")
		// 等其他中间件先执行
		context.Next()
		// 获取耗时
		latency := time.Since(t)
		log.Printf("total cost time:%d us", latency/1000)
	}
}
