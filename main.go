package main

import (
	"github.com/bin.jiang/deomapi/controller"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main()  {


	//db.DB()
	//路由

	router := gin.Default()

	router.GET("u/", controller.FetchAll)
	router.POST("/user", controller.AddUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.RemoveUser)

	router.Run(":8080")
}
