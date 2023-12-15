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
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
}
