package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"self-test/config"
)

var Mysql *gorm.DB

func InitMysql(c *config.ServerConf) {
	gin.SetMode(gin.ReleaseMode)
	var err error
	connSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=PRC",
		c.MysqlUser,
		c.MysqlPwd,
		c.MysqlHost,
		c.MysqlPort,
		c.MysqlDbname)
	Mysql, err = gorm.Open("mysql", connSource)

	if err != nil {
		defer Mysql.Close()
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	Mysql.DB().SetMaxOpenConns(c.MysqlPoolOpen)		//设置最大打开的连接数
	Mysql.DB().SetMaxIdleConns(c.MysqlPoolIdle)		//设置闲置的连接数
	Mysql.LogMode(false)

}
