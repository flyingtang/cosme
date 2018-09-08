package configs

import (
	"cosme/commands"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/sirupsen/logrus"
)

// 常量

// 类型
type MySQL struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type Config struct {
	MySQL MySQL
	HttpAddr string // 监听地址
	HttpPort string // 监听端口
}

var SySConfig *Config

// 方法
func ParseConfigFile() {
	configPath := fmt.Sprintf("configs/%s.config.json", commands.Env)
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal("ioutil.ReadFile read config file failed...")
	}

	if SySConfig == nil {
		SySConfig = new(Config)
	}
	err = json.Unmarshal(data, SySConfig)
	if err != nil {
		log.Fatal("json.Unmarshal parse config file failed...", err.Error())
	}
	logrus.Info(SySConfig)
	return
}

func (config *Config) GetDatabaseURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database)
}

// 获取MySQL的URL，主要用于创建数据库
func (config *Config) GetMySqlURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port)
}

func (c *Config) GetLocalAddr() string {
	return fmt.Sprintf("%s:%s", SySConfig.HttpAddr, SySConfig.HttpPort)
}
