package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// =========================
// SUCCESS
// =========================

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// =========================
// ERROR
// =========================

func Fail(c *gin.Context, status int, err error) {
	c.JSON(status, Response{
		Success: false,
		Error:   cleanError(err),
	})
}

// =========================
// CLEAN ERROR (CORE LOGIC)
// =========================

func cleanError(err error) string {
	msg := err.Error()

	// grpc error: "rpc error: code = AlreadyExists desc = user already exists"
	if i := strings.Index(msg, "desc = "); i != -1 {
		msg = msg[i+7:]
	}

	// ortiqcha technical qismlar
	msg = strings.ReplaceAll(msg, "rpc error: ", "")
	msg = strings.ReplaceAll(msg, "Error: ", "")

	return msg
}