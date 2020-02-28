package model

import "time"

type Userinfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
}

func (Userinfo) TableName() string {
	return "userinfo"
}
