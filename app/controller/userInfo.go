package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	request_struct "self-test/app/common-data/request-struct"
	"self-test/app/model"
	"self-test/app/utils"
	"self-test/dao/mysql"
)

var UserInfo userInfo

type userInfo struct{}

func (*userInfo) FindUserInfo(c *gin.Context) {
	var req request_struct.FindUserInfo
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, utils.Respone(utils.WRONG_PARAM, err.Error(), nil))
		return
	}

	var err error
	db := mysql.Mysql
	//查询列表
	var userList []model.UserInfoModel
	err = db.Model(&model.UserInfoModel{}).Where("name LIKE ? AND phone = ?", "%"+req.Name+"%", req.Phone).Find(&userList).Error

	//单个查询
	var singleUser model.UserInfoModel
	err = db.Model(&model.UserInfoModel{}).Where(map[string]interface{}{"name": req.Name, "phone": req.Phone}).First(&singleUser).Error

	//创建新记录
	createUser := model.UserInfoModel{
		Name:  "ssss",
		Phone: "sssssss",
	}
	trans := db.Begin() //开事务
	err = trans.Model(&model.UserInfoModel{}).Create(&createUser).Error
	//修改记录
	err = trans.Model(&model.UserInfoModel{}).Where("name = ?", req.Name).Update(map[string]interface{}{
		"name":  "qqq",
		"phone": "11111111111",
	}).Error
	if err != nil {
		trans.Rollback()
	}
	trans.Commit()

	var finaluserList []model.UserInfoModel
	err = db.Model(&model.UserInfoModel{}).Find(&finaluserList).Error

	c.JSON(http.StatusOK, utils.Respone(1, "查询成功", map[string]interface{}{
		"userlist":      userList,
		"singleUser":    singleUser,
		"finaluserList": finaluserList}))
	return
}
