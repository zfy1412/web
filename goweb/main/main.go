package main

import (

	"github.com/gin-gonic/gin"
	"goweb/send"
	"net/http"
)

func main ()  {
	r := gin.Default()
	r.LoadHTMLFiles(
		"./login.html",
		"./index.html",
		"./register.html",
		"./failed.html",
	)
	r.GET("/login", func(c *gin.Context){
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/index", send.DataAuthority(),func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/register", func(c *gin.Context){
		c.HTML(http.StatusOK, "register.html", nil)

	})
	r.POST("/login",func(c *gin.Context){
		id :=c.PostForm("id")
		username :=c.PostForm("username")
		password :=c.PostForm("password")
		user := &send.User{
			Id:       id,
			Name:     username,
			Password: password,
		}
		flag :=send.Ask(id,username,password)
		if flag==1 {
			key, _ := send.GenerateToken(*user)
			c.JSON(http.StatusOK, gin.H{"pass": true, "token": key})
			c.HTML(http.StatusOK, "index.html", gin.H{
				"id":       id,
				"Name":     username,
				"Password": password,
			},
			)
		} else {
			c.JSON(http.StatusOK, gin.H{"pass": false, "token": nil})
			c.HTML(http.StatusOK, "failed.html", nil)
		}

	})
	r.POST("/register",func(c *gin.Context){
		id :=c.PostForm("id")
		username :=c.PostForm("username")
		password :=c.PostForm("password")
		mark := send.Insert(id,username,password)
		if mark ==1 {
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/login")
			//c.Request.URL.Path = "/login"
			//r.HandleContext(c)
		}
	})
	r.Run(":8080")
}
