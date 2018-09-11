// https://github.com/gin-gonic/gin

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j1cken/3scale-demo/api"
	_ "github.com/j1cken/3scale-demo/docs"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title 3scale-demo example API
// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host application-3scale-amp-demo.apps.46.4.143.210.xip.io:80
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", api.Ping)

	// This handler will match /user/john but will not match neither /user/ or /user
	r.GET("/user/:name", api.User)

	r.GET("/welcome", api.Welcome)

	r.POST("/postMessage", api.Message)

	r.Run() // listen and serve on 0.0.0.0:8080
}
