package router

import (
	"gateway/controller"
	"gateway/golang_common/lib"
	"gateway/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	adminLoginRouter := router.Group("/admin_login")
	store, err := sessions.NewRedisStore(10, "tcp", lib.GetStringConf("base.session.redis_server"), lib.GetStringConf("base.session.redis_password"), []byte("secret"))
	if err != nil {
		log.Fatalf("sessions.NewRedisStore err:%v", err)
	}
	adminLoginRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.AdminLoginRegister(adminLoginRouter)
	}

	// router.group 每个独立路径下的接口
	adminRouter := router.Group("/admin")
	adminRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.AdminRegister(adminRouter)
	}

	return router
}
