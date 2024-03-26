package router

import (
	"jin-example/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tagsController *controller.TagsController, userController *controller.UserControllers) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")

	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	userRouter := baseRouter.Group("/user")
	userRouter.POST("", userController.CreateUser)
	userRouter.PATCH("/:userId", userController.UpdateUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
	userRouter.GET("/:userId", userController.FindUserById)
	userRouter.GET("", userController.FindAllUser)
	return router
}
