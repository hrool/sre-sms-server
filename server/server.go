package server

import (
	"sre-sms-server/api"
	"sre-sms-server/middleware/auth"

	"github.com/gin-gonic/gin"
)

// Run start a http server use gin
func Run() {
	//Init()
	app := gin.Default()
	app.GET("/status", api.Ping)

	authorized := app.Group("/sms", auth.BasicAuth())
	authorized.POST("/send", api.SmsSend)

	err := app.Run(":4000")
	if err != nil {
		panic(err)
	}
}
