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

	// Master Menu Detail Function
	repositoryMenuDetailFunction := repositories.NewRepositoryMenuDetailFunction(db)
	serviceMenuDetailFunction := services.NewServiceMenuDetailFunction(repositoryMenuDetailFunction)
	handlerMenuDetailFunction := handlers.NewHandlerMenuDetailFunction(serviceMenuDetailFunction)
	router.GET("/api/v1/master/menu-detail-function/ping", handlerMenuDetailFunction.HandlerPing)
	route.POST("/menu-detail-function/create", handlerMenuDetailFunction.HandlerCreate)
	route.GET("/menu-detail-function/results", handlerMenuDetailFunction.HandlerResults)
	route.DELETE("/menu-detail-function/delete/:id", handlerMenuDetailFunction.HandlerDelete)
	route.PUT("/menu-detail-function/update/:id", handlerMenuDetailFunction.HandlerUpdate)

	// Master Merchant
	repositoryMerchant := repositories.NewRepositoryMerchant(db)
	serviceMerchant := services.NewServiceMerchant(repositoryMerchant)
	handlerMerchant := handlers.NewHandlerMerchant(serviceMerchant)
	router.GET("/api/v1/master/merchant/ping", handlerMerchant.HandlerPing)
	route.POST("/merchant/create", handlerMerchant.HandlerCreate)
	route.GET("/merchant/results", handlerMerchant.HandlerResults)
	route.DELETE("/merchant/delete/:id", handlerMerchant.HandlerDelete)
	route.PUT("/merchant/update/:id", handlerMerchant.HandlerUpdate)

	// Master Outlet
	repositoryOutlet := repositories.NewRepositoryOutlet(db)
	serviceOutlet := services.NewServiceOutlet(repositoryOutlet)
	handlerOutlet := handlers.NewHandlerOutlet(serviceOutlet)
	router.GET("/api/v1/master/outlet/ping", handlerOutlet.HandlerPing)
	route.POST("/outlet/create", handlerOutlet.HandlerCreate)
	route.GET("/outlet/results", handlerOutlet.HandlerResults)
	route.DELETE("/outlet/delete/:id", handlerOutlet.HandlerDelete)
	route.PUT("/outlet/update/:id", handlerOutlet.HandlerUpdate)

	// Master Supplier
	repositorySupplier := repositories.NewRepositorySupplier(db)
	serviceSupplier := services.NewServiceSupplier(repositorySupplier)
	handlerSupplier := handlers.NewHandlerSupplier(serviceSupplier)
	router.GET("/api/v1/master/supplier/ping", handlerSupplier.HandlerPing)
	route.POST("/supplier/create", handlerSupplier.HandlerCreate)
	route.GET("/supplier/results", handlerSupplier.HandlerResults)
	route.DELETE("/supplier/delete/:id", handlerSupplier.HandlerDelete)
	route.PUT("/supplier/update/:id", handlerSupplier.HandlerUpdate)
}
