package request_struct

type FindUserInfo struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Phone string `form:"phone" json:"phone" binding:"-"`
}
