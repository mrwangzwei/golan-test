package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self-test/app/common/data/req"
	mysqlModel "self-test/app/model/mysql"
	"self-test/app/utils"
	"self-test/dao/mysql"
)

var UserInfo userInfo

type userInfo struct{}

func (*userInfo) FindUserInfo(c *gin.Context) {
	param := req.FindUserInfo{}
	var err error
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, utils.FailRespone(utils.WrongParam, err.Error(), nil))
		return
	}
	if err = param.Validator(); err != nil {
		c.JSON(http.StatusOK, utils.FailRespone(utils.WrongParam, err.Error(), nil))
		return
	}
	db := mysql.Mysql
	//查询列表
	var userList []mysqlModel.UserInfoModel
	err = db.Model(&mysqlModel.UserInfoModel{}).Where("name LIKE ? AND phone = ?", "%"+param.Name+"%", param.Phone).Find(&userList).Error

	//单个查询
	var singleUser mysqlModel.UserInfoModel
	err = db.Model(&mysqlModel.UserInfoModel{}).Where(map[string]interface{}{"name": param.Name, "phone": param.Phone}).First(&singleUser).Error

	//创建新记录
	createUser := mysqlModel.UserInfoModel{
		Name:  "ssss",
		Phone: "sssssss",
	}
	trans := db.Begin() //开事务
	err = trans.Model(&mysqlModel.UserInfoModel{}).Create(&createUser).Error
	//修改记录
	err = trans.Model(&mysqlModel.UserInfoModel{}).Where("name = ?", param.Name).Update(map[string]interface{}{
		"name":  "qqq",
		"phone": "11111111111",
	}).Error
	if err != nil {
		trans.Rollback()
	}
	trans.Commit()

	var finaluserList []mysqlModel.UserInfoModel
	err = db.Model(&mysqlModel.UserInfoModel{}).Find(&finaluserList).Error

	c.JSON(http.StatusOK, utils.SuccessRespone(map[string]interface{}{
		"userlist":      userList,
		"singleUser":    singleUser,
		"finaluserList": finaluserList}))
}
