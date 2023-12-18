package route

import (
	"time"

	"github.com/ZooLearn/zoo/api/middleware"
	"github.com/ZooLearn/zoo/config"
	"github.com/ZooLearn/zoo/ent"
	"github.com/gin-gonic/gin"
)

func Setup(env config.EnvConf, timeout time.Duration, db *ent.Client, gin *gin.Engine) {
	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	// Middleware to verify AccessToken
	NewTagRouter(env, timeout, db, publicRouter, protectedRouter)
	NewProfileRouter(env, timeout, db, protectedRouter)
}
