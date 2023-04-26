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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
