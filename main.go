package main

import (
	"cosme/controllers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"cosme/commands"
	"cosme/configs"
	"cosme/models"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
}

func main() {

	// 解析命令命令行参数
	commands.ParsecCmd()

	// 解析配置文件
	configs.ParseConfigFile()

	// 初始化数据库
	models.NewMysql()

	// gin 部分
	r := gin.Default()
	r.GET("/", controllers.IndexHandle)
	// 获取监听地址:端口
	addr := configs.SySConfig.GetLocalAddr()
	log.Info("listen http://",addr)
	r.Run(addr)
}
