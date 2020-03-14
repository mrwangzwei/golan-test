package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"self-test/config"
)

var Mysql *gorm.DB

func InitMysql() {
	c := config.Conf
	gin.SetMode(gin.ReleaseMode)
	var err error
	connSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=PRC",
		c.MysqlConfig.MysqlUser,
		c.MysqlConfig.MysqlPwd,
		c.MysqlConfig.MysqlHost,
		c.MysqlConfig.MysqlPort,
		c.MysqlConfig.MysqlDbname)
	Mysql, err = gorm.Open("mysql", connSource)

	if err != nil {
		defer Mysql.Close()
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	Mysql.DB().SetMaxOpenConns(c.MysqlConfig.MysqlPoolOpen) //设置最大打开的连接数
	Mysql.DB().SetMaxIdleConns(c.MysqlConfig.MysqlPoolIdle) //设置闲置的连接数
	Mysql.LogMode(false)

}
