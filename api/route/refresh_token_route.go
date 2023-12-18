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

func NewRefreshTokenRouter(env config.EnvConf, timeout time.Duration, db *ent.Client, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
