package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/webapp-auth-practice/config"
	"github.com/mochi-yu/webapp-auth-practice/controller"
	"github.com/mochi-yu/webapp-auth-practice/ent"
	"github.com/mochi-yu/webapp-auth-practice/middleware"
	"github.com/mochi-yu/webapp-auth-practice/repository"
	"github.com/mochi-yu/webapp-auth-practice/usecase"
)

func NewRouter(db *ent.Client, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors(cfg))
	router.Use(middleware.Options())
	router.Use(middleware.Transaction(db))

	mr := repository.NewMessageRepository(db)
	mu := usecase.NewMessageUsecase(mr)
	mc := controller.NewMessageController(mu)
	miscC := controller.NewMiscController(cfg)

	router.POST("/messages", mc.PostMessage)
	router.GET("/messages", mc.GetMessages)

	router.GET("/signin/github", miscC.SignIn)

	return router
}
