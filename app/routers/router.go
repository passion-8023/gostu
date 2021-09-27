package routers

import (
	"github.com/gin-gonic/gin"
	"gostu/app/controllers"
	"gostu/app/controllers/database"
	"gostu/app/controllers/es"
	"gostu/app/controllers/gorm"
	"gostu/app/controllers/web"
	"gostu/app/controllers/webcontent"
)

func Routers(app *gin.Engine) {
	group := app.Group("api")
	AuthRouter(group)
	DataBaseRouter(group)
	GormRouter(group)
	WebRouter(group)
	EsRouter(group)
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

func WebRouter(group *gin.RouterGroup) {
	webRouter := group.Group("web")
	{
		webRouter.GET("get/list", web.GetWebList)
		webRouter.GET("get/data", web.GetWebContent)
		webRouter.GET("get/web/content", webcontent.GetWebContent)
	}
}

func EsRouter(group *gin.RouterGroup) {
	esRouter := group.Group("es")
	{
		esRouter.GET("create/index", es.CreateIndex)
		esRouter.GET("check/index", es.CheckIndex)
		esRouter.GET("delete/index", es.DeleteIndex)
		esRouter.GET("bulk/data", es.BulkData)
	}
}
