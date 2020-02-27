package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	request_struct "self-test/app/common-data/request-struct"
	"self-test/app/model"
	"self-test/app/utils"
	"self-test/dao"
)

func FindUserInfo(c *gin.Context) {
	var req request_struct.FindUserInfo
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, utils.Respone(utils.WRONG_PARAM, err.Error(), nil))
		return
	}
	var user_info model.User
	fmt.Println(req.Name)
	db := dao.Mysql
	db.Table("userinfo").Where("name LIKE ?", "%"+req.Name+"%").First(&user_info)

	fmt.Println(user_info)

	c.JSON(http.StatusOK, utils.Respone(1, "查询成功", nil))
	return
}
