package routers

import (
	"net/http"
	"time"
	"yuekao/bff/handler/service"
	"yuekao/bff/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	r.POST("Login", service.Login)
	r.POST("upload", service.Upload)
	r.POST("GetGoods", service.GetGoods)
	r.POST("Sx", service.Sx)
	r.POST("GoodsAdd", middleware.Reg(), middleware.Loggers(), service.GoodsAdd)
	return r
}
