package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-pos-service/handlers"
	"github.com/nutwreck/admin-pos-service/middlewares"
	"github.com/nutwreck/admin-pos-service/repositories"
	"github.com/nutwreck/admin-pos-service/services"
	"gorm.io/gorm"
)

func NewRouteMaster(db *gorm.DB, router *gin.Engine) {
	route := router.Group("/api/v1/master")
	route.Use(middlewares.AuthToken(db))
	route.Use(middlewares.AuthRole(db))

	//Master Role
	repositoryRole := repositories.NewRepositoryRole(db)
	serviceRole := services.NewServiceRole(repositoryRole)
	handlerRole := handlers.NewHandlerRole(serviceRole)
	router.GET("/api/v1/master/role/ping", handlerRole.HandlerPing)
	router.GET("/api/v1/master/role/results", handlerRole.HandlerResults)
	route.POST("/role/create", handlerRole.HandlerCreate)
	route.DELETE("/role/delete/:id", handlerRole.HandlerDelete)
	route.PUT("/role/update/:id", handlerRole.HandlerUpdate)

	// Master Menu
	repositoryMenu := repositories.NewRepositoryMenu(db)
	serviceMenu := services.NewServiceMenu(repositoryMenu)
	handlerMenu := handlers.NewHandlerMenu(serviceMenu)
	router.GET("/api/v1/master/menu/ping", handlerMenu.HandlerPing)
	route.POST("/menu/create", handlerMenu.HandlerCreate)
	route.GET("/menu/results", handlerMenu.HandlerResults)
	route.DELETE("/menu/delete/:id", handlerMenu.HandlerDelete)
	route.PUT("/menu/update/:id", handlerMenu.HandlerUpdate)

	// Master Menu Detail
	repositoryMenuDetail := repositories.NewRepositoryMenuDetail(db)
	serviceMenuDetail := services.NewServiceMenuDetail(repositoryMenuDetail)
	handlerMenuDetail := handlers.NewHandlerMenuDetail(serviceMenuDetail)
	router.GET("/api/v1/master/menu-detail/ping", handlerMenuDetail.HandlerPing)
	route.POST("/menu-detail/create", handlerMenuDetail.HandlerCreate)
	route.GET("/menu-detail/results", handlerMenuDetail.HandlerResults)
	route.DELETE("/menu-detail/delete/:id", handlerMenuDetail.HandlerDelete)
	route.PUT("/menu-detail/update/:id", handlerMenuDetail.HandlerUpdate)
}
