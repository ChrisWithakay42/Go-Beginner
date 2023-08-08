package router

import (
	"github.com/gin-gonic/gin"
	"go_rest/controller"
	"net/http"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	baseRouter := router.Group("/api")
	userRouter := baseRouter.Group("/user")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	return router
}
