package router

import (
	v1 "Blog-BackEnd/router/api/v1"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter(){
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	userGroup:=r.Group("/admin")
	userGroup.GET("/users/:id",v1.GetUser)
	userGroup.POST("/users",v1.AddUser)
	userGroup.DELETE("/users/:id",v1.DeleteUser)
	userGroup.PUT("/users/:id",v1.EditUser)
}