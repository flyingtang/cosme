package controllers

import (

	"github.com/gin-gonic/gin"
	"fmt"
)

func IndexHandle(c *gin.Context){
	fmt.Println("1233")
	c.JSON(200, gin.H{

		"Message":"启动成功",
	})
}
