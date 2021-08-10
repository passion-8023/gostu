package validates

type Auth struct {
	Account  string `form:"account" json:"account" binding:"required" field:"账号"`
	Password string `form:"password" json:"password" binding:"required" field:"密码"`
}

type SignUpParam struct {
	Age        uint8  `form:"age" json:"age" binding:"gte=1,lte=130"`
	Name       string `form:"name" json:"name" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"re_password" json:"re_password" binding:"required,eqfield=Password"`
}

type ShowUploadFile struct {
	Url string `form:"url" json:"url" binding:"required"`
}
