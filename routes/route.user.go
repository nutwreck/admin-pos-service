package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/handlers"
	"github.com/nutwreck/admin-pos-service/middlewares"
	"github.com/nutwreck/admin-pos-service/repositories"
	"github.com/nutwreck/admin-pos-service/services"
)

func NewRouteUser(db *gorm.DB, router *gin.Engine) {
	repositoryUser := repositories.NewRepositoryUser(db)
	serviceUser := services.NewServiceUser(repositoryUser)
	handlerUser := handlers.NewHandlerUser(serviceUser)
	routeUser := "/api/v1/auth"

	route := router.Group(routeUser)

	routePrivate := router.Group(routeUser)
	routePrivate.Use(middlewares.AuthToken())
	routePrivate.Use(middlewares.AuthRole(configs.RoleConfig))

	route.GET("/ping", handlerUser.HandlerPing)
	route.POST("/register", handlerUser.HandlerRegister)
	route.POST("/login", handlerUser.HandlerLogin)
	routePrivate.GET("/refresh-token", handlerUser.HandlerRefreshToken)
	routePrivate.PUT("/update", handlerUser.HandlerUpdate)
	routePrivate.GET("/data-user", handlerUser.HandleDataUser)
}
