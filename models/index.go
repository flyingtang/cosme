package models

import (
	"cosme/commands"
	"cosme/configs"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// 创建数据库连接
var MySQL *sql.DB

// 定义结构体

// 初始化数据库

// 数据库如果不存在就创建数据库
func createDatabase() {
	MysqlUrl := configs.SySConfig.GetMySqlURL()
	log.Info("MySQL address : ", MysqlUrl)
	db, err := sql.Open("mysql", MysqlUrl)
	if err != nil {
		log.Fatal("sql.Open", err.Error())
	}
	defer db.Close()

	createDatabaseSql := fmt.Sprintf("create database if not exists `cosme_%s` character set UTF8", commands.Env)
	if _, err := db.Exec(createDatabaseSql); err != nil {
		log.Fatal("db.Exec(mkdirDatabase)", err.Error())
	} else {
		log.Info("create database successful...")
	}
}

// 初始化数据库
func NewMysql() {
	// 初始化数据库
	createDatabase()
	MysqlUrl := configs.SySConfig.GetDatabaseURL()
	log.Info("connect MySQL addr is: ", MysqlUrl)

	db, err := sql.Open("mysql", MysqlUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	MySQL = db
	InitialTable()
}

// 初始化表
func InitialTable() {
	for _, sql := range tables {
		if _, err := MySQL.Exec(sql); err != nil {
			log.Fatal("db.Exec(sql)", err.Error())
			continue
		}
		log.Info(sql[:45], ".... successful")
	}
}
