package routes

import (
	"log"
	"rbarrero/visago/gobex"
	"rbarrero/visago/handlers"
	"rbarrero/visago/middlewares"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// SetupRoutes Set routes
func SetupRoutes() *gin.Engine {
	var gobex gobex.Gobex = gobex.LdapInit()
	router := gin.Default()
	//router.Use(middlewares.GobexMiddleware(gobex))
	authMiddleware := middlewares.JwtMiddleware()
	router.GET("/ping", handlers.PingHandler)
	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	auth := router.Group("/auth")
	auth.Use(middlewares.GobexMiddleware(gobex))
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/ping_with_token", handlers.PingHandler)
		auth.POST("/change_password", handlers.ChangePasswordHandler)
	}

	return router
}
