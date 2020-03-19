package mysql

import "time"

type UserInfoModel struct {
	ID        uint64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserInfoModel) TableName() string {
	return "userinfo"
}
