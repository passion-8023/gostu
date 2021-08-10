package routers

import (
	"github.com/gin-gonic/gin"
	"gostu/app/controllers"
	"gostu/app/controllers/database"
	"gostu/app/controllers/gorm"
)

func Routers(app *gin.Engine) {
	group := app.Group("api")
	AuthRouter(group)
	DataBaseRouter(group)
	GormRouter(group)
}

func AuthRouter(group *gin.RouterGroup) {
	authRouter := group.Group("")
	{
		authRouter.GET("captcha", controllers.Captcha)
		authRouter.POST("login", controllers.PasswordLogin)
		authRouter.POST("upload", controllers.FileUpload)
		authRouter.GET("show", controllers.ShowFile)
		authRouter.GET("file/test", controllers.FileTest)
		authRouter.GET("file/test/bufio", controllers.FileTestBufio)
		authRouter.GET("file/test/ioutil", controllers.FileTestIoutil)
		authRouter.GET("file/test/write", controllers.FileTestWrite)

		authRouter.GET("get/sin/png", controllers.GetSinPng)
		authRouter.GET("get/cos/png", controllers.GetCosPng)
	}
}

func DataBaseRouter(group *gin.RouterGroup) {
	databaseRouter := group.Group("database")
	{
		databaseRouter.GET("data/add", database.DatabaseInsert)
		databaseRouter.GET("data/update", database.DatabaseUpdate)
		databaseRouter.GET("data/query", database.DatabaseQuery)
	}
}

func GormRouter(group *gin.RouterGroup) {
	gormRouter := group.Group("gorm")
	{
		gormRouter.GET("add/info", gorm.AddInfo)
		gormRouter.GET("select/info", gorm.QueryInfo)
		gormRouter.GET("get/user/info", gorm.GetUserInfo)
	}
}
