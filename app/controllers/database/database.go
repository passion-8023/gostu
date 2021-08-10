package database

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gostu/app/models"
	"gostu/app/validates"
	"gostu/app/validates/rules"
	"gostu/pkg/database"
	"gostu/pkg/response"
	"time"
)

func DatabaseInsert(c *gin.Context) {
	var data validates.Userinfo
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ErrorResponse(c, response.ValidateCheckError, err.Error())
			return
		}
		response.ErrorResponse(c, response.ValidateCheckError, rules.Translate(errs))
		return
	}

	//插入数据
	stmt, err := database.Eloquent.Prepare("insert into userinfo set username =?, department=?, created=?")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			response.ErrorResponse(c, response.ValidateCheckError, err.Error())
			return
		}
	}()

	result, err := stmt.Exec(data.Username, data.Department, time.Now().Format("2006-01-02"))
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	response.SuccessResponse(c, "插入成功", id)
}

func DatabaseUpdate(c *gin.Context) {
	var data validates.UpdateUserInfo
	if err := c.ShouldBind(&data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ErrorResponse(c,3, err.Error())
			return
		}
		response.ErrorResponse(c, response.ValidateCheckError, rules.Translate(errs))
		return
	}
	stmt, err := database.Eloquent.Prepare("update userinfo set username=?, department=? where uid=?")
	if err != nil {
		response.ErrorResponse(c, 1, err.Error())
		return
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			response.ErrorResponse(c, 2, err.Error())
			return
		}
	}()
	result, _ := stmt.Exec(data.Username, data.Department, data.Uid)
	affected, _ := result.RowsAffected()
	response.SuccessResponse(c, "修改成功", affected)
}

func DatabaseQuery(c *gin.Context) {
	//var data models.Userinfo
	//row := database.Eloquent.QueryRow("select * from userinfo where uid=?", c.Query("uid"))
	//err := row.Scan(&data.Uid, &data.Username, &data.Department, &data.Created)
	//if err != nil {
	//	response.ErrorResponse(c, 2, err.Error())
	//	return
	//}
	//response.SuccessResponse(c, "获取成功", data)

	rows, err := database.Eloquent.Query("select * from userinfo where uid > ?", c.Query("uid"))
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	defer rows.Close()
	var result []models.Userinfo
	for rows.Next() {
		var data models.Userinfo
		rows.Scan(&data.Uid, &data.Username, &data.Department, &data.Created)
		result = append(result, data)
	}
	response.SuccessResponse(c, "获取成功", result)
}

