package gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gostu/app/models"
	"gostu/app/validates"
	"gostu/app/validates/rules"
	"gostu/pkg/gorm"
	"gostu/pkg/logger"
	"gostu/pkg/response"
)

func AddInfo(ctx *gin.Context)  {
	//user := models.User{
	//	Username: "小王爷",
	//	Sex:      1,
	//	Mobile:   "123455",
	//	LoginAt: time.Now(),
	//}

	//result := gorm.Eloquent.Create(&user)
	//result := gorm.Eloquent.Select("Username", "Mobile", "LoginAt", "CreatedAt").Create(&user)
	//result := gorm.Eloquent.Omit("Username", "Sex", "Mobile").Create(&user)

	//users := []models.User{
	//	{
	//		Username: "马云",
	//		Sex: 1,
	//		Mobile: "1380000000",
	//		LoginAt: time.Now(),
	//	},
	//	{
	//		Username: "吴亦凡",
	//		Sex: 1,
	//		Mobile: "110110110",
	//		LoginAt: time.Now(),
	//	},
	//	{
	//		Username: "郑爽",
	//		Sex: 2,
	//		Mobile: "1002004380",
	//		LoginAt: time.Now(),
	//	},
	//}



	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}

	//response.SuccessResponse(ctx, "", result)
}

func QueryInfo(ctx *gin.Context)  {
	//users := models.GetInstance().GetTotalUsersInfo("username, age", map[string]interface{}{"age":34})
	users := []models.User{}

	gorm.Eloquent.Select("username, age").Find(&users)


	response.SuccessResponse(ctx, "获取成功", users)



	//user := models.User{
	//	ID: 8,
	//}
	//users := []models.User{}
	//user := models.User{}
	//if err := gorm.Eloquent.First(&user).Error; err != nil {
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//	return
	//}

	//if err := gorm.Eloquent.Take(&user).Error; err != nil {
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//	return
	//}

	//if err := gorm.Eloquent.Last(&user).Error; err != nil {
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//	return
	//}

	//if err := gorm.Eloquent.Find(&users).Error; err != nil {
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//}

	//if err := gorm.Eloquent.First(&user, 7).Error; err != nil {
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//	return
	//}

	//if err := gorm.Eloquent.Where("username=?", "吴亦凡").First(&user).Error; err != nil {
	//	response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
	//	return
	//}

	//gorm.Eloquent.Where("username <> ?", "吴亦凡").Find(&users)
	//gorm.Eloquent.Where("username like ?", "%凡%").Find(&users)
	//gorm.Eloquent.Where("login_at Between ? and ?", "2021-08-04 14:00:00", "2021-08-04 14:50:00").Find(&users)

	//gorm.Eloquent.Where(&models.User{Sex: 1, Mobile: "110"}).First(&user)

	//gorm.Eloquent.Where(map[string]interface{}{"sex":1, "mobile":9999}).Find(&users)

	//gorm.Eloquent.Where([]int{6, 8, 9}).Find(&users)

	//gorm.Eloquent.Where(&models.User{Age: 0, Sex: 1}).Find(&users)

	//users := []models.User{}
	//user := models.User{}
	////gorm.Eloquent.First(&user, 6)
	////gorm.Eloquent.First(&user, "sex=?", 1)
	////gorm.Eloquent.Find(&users, models.User{Age: 29})
	//
	//init := gorm.Eloquent.Attrs(models.User{LoginAt: time.Now()}).FirstOrInit(&user, models.User{Username: "马云"})
	//init.Create(&user)
	//
	//gorm.Eloquent.Where("amount > ?", gorm.Eloquent.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	//// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	//
	//response.SuccessResponse(ctx, "", user)
}

func GetUserInfo(ctx *gin.Context)  {
	var data validates.Userinfo
	if err := ctx.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			logger.Logger.Debugln("参数错误", err)
			response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
			return
		}
		logger.Logger.Errorln("系统错误", rules.Translate(errs))
		response.ErrorResponse(ctx, response.ValidateCheckError, rules.Translate(errs))
		return
	}

	var user models.Userinfo
	if err := gorm.Eloquent.Where("username = ?", data.Username).Select("username").First(&user).Error; err != nil {
		logger.Logger.Infoln("获取错误", err)
		response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
		return
	}

	logger.Logger.Infoln("获取成功", user)
	response.SuccessResponse(ctx, "", user)
}
