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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
