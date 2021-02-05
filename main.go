/*
# @Time : 2021/2/5 13:24
# @Author : smallForest
# @SoftWare : GoLand
*/
package main

import (
	"dingshi/application"
	"encoding/json"
	"fmt"
	cron "github.com/robfig/cron/v3"
	"github.com/codeskyblue/go-sh"
)

type Task struct {
	Timer   string
	Command string
}

func main() {

	// 获取配置项目的task_json
	task_json := application.Task

	// 将json字符串格式化成数组
	tasks := make([]Task, 1)
	err := json.Unmarshal([]byte(task_json), &tasks)
	if err != nil {
		fmt.Println("格式化参数错误", err)
		return
	}
	fmt.Println("格式化参数成功：")
	fmt.Printf("%+v", tasks)
	fmt.Println()

	crontab := cron.New()

	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	fmt.Printf("一共%d个任务", len(tasks))
	fmt.Println()

	for _, v := range tasks {
		id, err := crontab.AddFunc(v.Timer, func() {
			fmt.Println(v.Command)
			sh.NewSession().SetDir("C:\\phpstudy_pro\\WWW\\jinlianlian-bank-system\\public\\").Command("php","index.php" ,v.Command).Run()

		})
		if err != nil {
			fmt.Println("添加定时任务失败")
			fmt.Println(v.Timer)
			fmt.Println(v.Command)
			return
		}
		fmt.Println("创建任务成功，id：", id)
	}

	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {}
}
