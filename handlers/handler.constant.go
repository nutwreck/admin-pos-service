package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/helpers"
)

/**
* ==========================================
* Handler Results All Type Role Teritory
*===========================================
 */
// GetConstants godoc
// @Summary		Get List Type Role
// @Description	Get List Type Role
// @Tags		Constant
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Router /api/v1/constant/type-role [get]
func HandlerTypeRole(ctx *gin.Context) {
	if len(constants.RoleTypes) == constants.EMPTY_NUMBER {
		helpers.APIResponse(ctx, "Type Role data not found", http.StatusInternalServerError, nil)
		return
	}
	helpers.APIResponse(ctx, "Type Role data already to use", http.StatusOK, constants.RoleTypes)
}
