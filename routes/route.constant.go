package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-pos-service/handlers"
)

func NewRouteConstant(router *gin.Engine) {
	route := router.Group("/api/v1/constant")

	route.GET("/type-role", handlers.HandlerTypeRole)
}
