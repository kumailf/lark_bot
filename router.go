package main

import (
	"code.byted.org/larkim/oapi_demo/biz"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var post = map[string]gin.HandlerFunc{
	"/webhook/event": biz.ReceiveEvent,
}

func GetPostRouter() map[string]gin.HandlerFunc {
	return post
}

func Ping(c *gin.Context) {
	logrus.Info("a sample app log")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
