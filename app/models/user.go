package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Username string `gorm:"default:'小王子'"`
	Password string
	Age int
	Sex int
	Mobile string
	LoginAt time.Time `gorm:"default:'0000-00-00 00:00:00'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//var instance *user
//var once sync.Once
//
//func GetInstance() *user {
//	once.Do(func() {
//		instance = &user{}
//	})
//	return instance
//}
//
////获取user表中的所有信息
//func (u *user) GetTotalUsersInfo(fields string, conditions map[string]interface{}) []user {
//	users := []user{}
//	model := gorm.Eloquent
//	if len(fields) != 0 {
//		model = model.Select(fields)
//	}
//	if conditions != nil {
//		model = model.Where(conditions)
//	}
//	if err := model.Find(&users).Error; err != nil {
//		return nil
//	}
//
//	return users
//}
