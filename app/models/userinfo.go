package models

import "time"

type Userinfo struct {
	Uid int `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Department string `gorm:"column:department" db:"department" json:"department" form:"department"`
	Created time.Time `gorm:"column:created" db:"created" json:"created" form:"created"`
}
