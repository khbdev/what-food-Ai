package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SUCCESS
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// ERROR
func Fail(c *gin.Context, status int, msg string) {
	c.JSON(status, Response{
		Success: false,
		Error:   msg,
	})
}