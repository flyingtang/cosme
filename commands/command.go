package commands

import (
	"flag"

	"github.com/gin-gonic/gin"
)

// 主要是解析命令
func ParsecCmd() {
	var env string
	flag.StringVar(&env, "env", "dev", "dev | pro | test")
	flag.Parse()

	// 解析env变量
	parseEnv(env)
}

var Env string

// 解析env变量
func parseEnv(env string) {

	switch env {
	case "pro":
		gin.SetMode("ReleaseMode")
	case "test":
		gin.SetMode("TestMode")
	case "dev":
		gin.SetMode("")
	default:
		panic("error env parameter")
	}
	Env = env
}
