package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-pos-service/handlers"
)

func NewRouteConstant(router *gin.Engine) {
	route := router.Group("/api/v1/constant")

	route.GET("/jenis-kelamin", handlers.HandlerJenisKelamin)
	route.GET("/status-pernikahan", handlers.HandlerStatusPernikahan)
	route.GET("/role-user", handlers.HandlerRoleUser)
}
