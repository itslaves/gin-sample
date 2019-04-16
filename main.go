package main

import (
	"gin-sample/article"
	"gin-sample/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SessionKey = "sessionKey"
	SessionSecret = "sessionSecret"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte(SessionSecret))
	r.Use(sessions.Sessions(SessionKey, store))

	r.LoadHTMLGlob("templates/*")

	r.GET("/login", auth.LoginView)
	r.GET("/auth/callback", auth.Authenticate)

	r.GET("logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
	})

	r.GET("/index", func(c *gin.Context) {
		user := auth.CurrentUser(c)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"name": user.Name,
			"email": user.Email,
		})
	})

	r.GET("/articles", article.Retrieve)
	r.POST("/articles", article.Create)
	r.PUT("/articles/:id", article.Update)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}