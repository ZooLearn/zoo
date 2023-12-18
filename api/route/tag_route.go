package route

import (
	"time"

	"github.com/ZooLearn/zoo/api/controller"
	"github.com/ZooLearn/zoo/usecase"

	"github.com/ZooLearn/zoo/config"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/repository"
	"github.com/gin-gonic/gin"
)

func NewTagRouter(env config.EnvConf, timeout time.Duration, db *ent.Client, groupPublic *gin.RouterGroup, groupProtect *gin.RouterGroup) {
	tr := repository.NewTagRepository(db)
	lc := controller.TagController{
		TagUsecase: usecase.NewTagUsecase(tr, timeout),
	}
	groupPublic.GET("/tag/:id", lc.GetByID)
	groupPublic.GET("/tag/all", lc.FetchAll)
	groupPublic.GET("/tag", lc.Fetch)
	groupPublic.POST("/tag", lc.Create)
	groupPublic.PUT("/tag", lc.Update)
	groupPublic.DELETE("/tag/:id", lc.Delete)
}
