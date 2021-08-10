package gin

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	http := gin.New()
	http.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(gin.DebugMode)
	return http
}
