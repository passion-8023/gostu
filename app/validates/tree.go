package validates

type Tree struct {
	Num []int `form:"num" json:"num" binding:"required" field:"数据"`
}
