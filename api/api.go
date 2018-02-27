package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ...
// @Summary Get a Pong back
// @Description get pong
// @ID ping
// @Produce  json
// @Success 200 {object} api.HelloResponse "ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// User ...
// @Summary Greet a user by name
// @Description greet user
// @ID user
// @Produce  json
// @Param   name     path    string     true        "Name of person"
// @Success 200 {object} api.HelloResponse "ok"
// @Router /user/{name} [get]
func User(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"hello": name})
}

// Welcome ...
// @Summary Welcome someone
// @Description Welcome guest unless given the name
// @ID welcome
// @Produce  json
// @Param   firstname     query    string     true        "first name"
// @Param   lastname     query    string     true        "last name"
// @Success 200 {object} api.UserResponse "ok"
// @Router /welcome [get]
func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.JSON(http.StatusOK, gin.H{"firsname": firstname, "lastname": lastname})
}

// Message ...
// @Summary Post a message
// @Description Post a message and return it.
// @ID postMessage
// @Accept text/plain
// @Produce  json
// @Param   message     body    string     true        "message=My Message"
// @Success 200 {object} api.MessageResponse "ok"
// @Router /postMessage [post]
func Message(c *gin.Context) {
	message := c.PostForm("message")

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// HelloResponse object
type HelloResponse struct {
	Name string `json:"hello"`
}

// MessageResponse object
type MessageResponse struct {
	Message string `json:"message"`
}

// UserResponse object
type UserResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
