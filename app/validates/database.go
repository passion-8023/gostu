package validates

type Userinfo struct {
	Username string `gorm:"column:username" db:"username" json:"username" form:"username" binding:"required"`
	Department string `gorm:"column:department" db:"department" json:"department" form:"department" binding:"required"`
}

type UpdateUserInfo struct {
	Uid int `json:"uid" form:"uid" binding:"required,numeric"`
	Username string `json:"username" form:"username" binding:"required"`
	Department string `json:"department" form:"department" binding:"required"`
}



