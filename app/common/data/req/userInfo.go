package req

import "errors"

type FindUserInfo struct {
	Name  string `form:"name" json:"name"`
	Phone string `form:"phone" json:"phone"`
}

func (f *FindUserInfo) Validator() error {
	if len(f.Name) < 1 {
		return errors.New("name is required")
	}
	return nil
}
