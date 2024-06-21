package main

import (
	"fmt"
	"new-bubble/dao"
	"new-bubble/models"
	"new-bubble/settings"
	"os"
)

const defaultConfFile = "./conf/config.ini"

func main() {
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("use default conf file", confFile)
	}
	//加载配置文件
	if err := settings.Init(confFile); err != nil {
		fmt.Println("init config failed", err)
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMysql(settings.Conf.MysqlConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 创建表
	dao.DB.AutoMigrate(&models.Todo{})
	defer dao.Close() // 程序退出关闭数据库连接
}
