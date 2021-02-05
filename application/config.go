/*
# @Time : 2019-07-22 16:14
# @Author : smallForest
# @SoftWare : GoLand
*/
package application

import "dingshi/conf"

// 数据库信息
var Task = conf.Run().Section("task").Key("task_json").String()