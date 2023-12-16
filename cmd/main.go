package main

import (
	"time"

	route "github.com/ZooLearn/zoo/api/route"
	"github.com/ZooLearn/zoo/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, app.Ent, gin)
	if err := gin.Run(env.ServerAddress); err != nil {
		panic(err)
	}
}
