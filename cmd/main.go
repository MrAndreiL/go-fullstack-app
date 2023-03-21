package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/MrAndreiL/go-fullstack-app/apis/authenticator"
	"github.com/MrAndreiL/go-fullstack-app/database"
)

func main() {
	// Set up database connection.
	database.Connect()

	defer database.CloseConnection()

	// Set up sessions.
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60}) // 1 hour
	router.Use(sessions.Sessions("mysession", store))

	// Set up Gin routing mechanism.
	router.POST("/apis/register", authenticator.PostAuth)
	router.POST("/apis/login", authenticator.PostLogin)

	// Start up service.
	router.Run("localhost:8081")
}
