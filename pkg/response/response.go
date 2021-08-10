package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//错误返回
func ErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":      400,
		"status_code": code,
		"message":     msg,
	})
}

//成功返回
func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": msg,
		"data":    data,
	})
}
