package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shouyindemo/conf"
	"shouyindemo/middleware"
	"shouyindemo/routers"
	"shouyindemo/utils"
)

var db = make(map[string]string)

func main() {
	conf, err := conf.ParseConfig("./conf/app.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	utils.InitMySqlUtil(conf.Database)
	fmt.Printf("conf:%#v\n", conf)
	r := gin.Default()
	r.Use(middleware.Test)
	routers.RegisterRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9999")
}
