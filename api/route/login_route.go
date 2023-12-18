package route

import (
	"time"

	"github.com/ZooLearn/zoo/api/controller"
	"github.com/ZooLearn/zoo/config"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/repository"
	"github.com/ZooLearn/zoo/usecase"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env config.EnvConf, timeout time.Duration, db *ent.Client, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
