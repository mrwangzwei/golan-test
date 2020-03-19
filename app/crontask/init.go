package crontask

import (
	"fmt"
	"github.com/robfig/cron"
)

//秒 分 时 日 月
const spec = "0 0 1 * * ?"

func Run() {
	c := cron.New()
	err := c.AddFunc(spec, func() {
		//todo 执行内容
	})
	if err != nil {
		fmt.Println("定时任务开启失败", err)
	}
	c.Start()
	defer c.Stop()
	select {}
}
