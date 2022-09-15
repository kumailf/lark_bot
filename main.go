// Start a web server
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	post := GetPostRouter() 
	for k, v := range post {
		r.POST(k, v)
	}
	r.GET("/ping", Ping)
	if err := r.Run(":8089"); err != nil {
		logrus.WithError(err).Errorf("init fail")
	}
}
