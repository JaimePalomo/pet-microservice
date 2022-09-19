package http_ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Pong handler del endpoint /ping
func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
