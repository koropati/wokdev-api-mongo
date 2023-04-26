package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"wokdev-api/api/controller"
	"wokdev-api/bootstrap"
	"wokdev-api/domain"
	"wokdev-api/mongo"
	"wokdev-api/repository"
	"wokdev-api/usecase"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
