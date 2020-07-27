package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strings"
)
type User struct {
	Username  string
	Id		  string
	Usertype  string
}
//记录访问量
var totalRequests = 0
func main(){
	//config.Init()
	New()

}
func New(){
	// Echo instance
	e := echo.New()
	// Middleware
	//中间件打印http请求日志
	e.Use(middleware.Logger())
	//拦截panic错误并且在控制台打印错误日志,避免echo程序崩溃
	e.Use(middleware.Recover())
	e.Use(Count)
	e.Static("/static","static")
	// Route => handler
	e.GET("/",func(c echo.Context) error{
		return c.String(http.StatusOK,"Hello, World!\n")
	})
	e.POST("/users",func(c echo.Context) error{
		return nil
	})
	e.GET("/users/:id", getUser)
	//定义get请求，url模式为：/users/:id  （:id是参数，例如: /users/10, 会匹配这个url模式），绑定getUser控制器函数
	e.PUT("/users/:id",func(c echo.Context)error{
		return nil
	})
	e.DELETE("/users/:id",func(c echo.Context)error{
		return nil
	})
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
/* 控制器函数只接收一个echo.Context上下文参数 关联了当前请求和相应 通过参数c我们可以获取请求参数,
	向客户端相应结果
*/
func getUser(c echo.Context) error{
	// 得到url path上的参数
	id := c.Param("id")
	// 获取query参数
	username := c.QueryParam("username")
	usertype := c.QueryParam("usertype")
	var build strings.Builder  //string.Buffer
	build.WriteString(id+"-")
	build.WriteString(username)
	build.WriteString(":")
	build.WriteString(usertype)
	u := User{
		Username: "邓浩",
		Usertype: "学生",
		Id: "201922081008",
	}
	//return c.String(http.StatusOK,build.String())
	return c.JSON(http.StatusOK,u)
}

//自定义中间件
func Count(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error{
		totalRequests++
		c.Response().Header().Add("requests",fmt.Sprintf("%d", totalRequests))

		return next(c)
	}
}