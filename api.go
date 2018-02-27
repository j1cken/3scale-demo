// https://github.com/gin-gonic/gin

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// This handler will match /user/john but will not match neither /user/ or /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"hello": name})
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.JSON(http.StatusOK, gin.H{"firsname": firstname, "lastname": lastname})
	})

	r.POST("/postMessage", func(c *gin.Context) {
		message := c.PostForm("message")

		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
