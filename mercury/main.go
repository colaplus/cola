package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"mercury/controller/account"
	"mercury/controller/ask"
	"mercury/controller/category"
	"mercury/dal/db"
	"mercury/filter"
	"mercury/id_gen"
	maccount "mercury/middleware/account"
)

func initTemplate(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.icon", "./static/favicon.icon")
	router.Static("/css/", "./static/css/")
	router.Static("/fonts/", "./static/fonts/")
	router.Static("/img/", "./static/img/")
	router.Static("/js/", "./static/js/")
}

func initDb() (err error) {
	dns := "root:Km_123456@tcp(39.101.194.252:3306)/mercury?parseTime=true"
	err = db.Init(dns)
	if err != nil {
		return
	}
	return
}

func initSession() (err error) {
	err = maccount.InitSession("mem", "")
	return
}

func main() {
	router := gin.Default()

	err := initDb()
	if err != nil {
		panic(err)
	}

	err = initSession()
	if err != nil {
		panic(err)
	}

	err = id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	err = filter.Init("./data/filter.dat.txt")
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	initTemplate(router)
	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	router.GET("/api/category/lists", category.GetCategoryListHandle)
	router.POST("/api/ask/submit", ask.QuestionSubmitHandle)

	router.Run(":9090")
}
