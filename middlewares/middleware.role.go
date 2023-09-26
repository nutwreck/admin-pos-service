package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/helpers"
	"github.com/nutwreck/admin-pos-service/pkg"
	"github.com/nutwreck/admin-pos-service/repositories"
	"github.com/nutwreck/admin-pos-service/schemes"
)

func AuthRole(db *gorm.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		bearer := ctx.GetHeader("Authorization")

		if bearer == constants.EMPTY_VALUE {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "message": "Authorization header is required"})
			return
		}

		token := strings.Split(bearer, " ")

		if len(token) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "message": "Access Token is required"})
			return
		}

		decodeToken, err := pkg.VerifyToken(strings.TrimSpace(token[1]), pkg.GodotEnv("JWT_SECRET_KEY"))

		if err != nil {
			defer logrus.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "message": "Access Token Expired"})
		}

		accessToken := helpers.ExtractToken(decodeToken)

		//CEK ROLE
		newUser := repositories.NewRepositoryUser(db)
		var bodyRole schemes.Role
		bodyRole.ID = accessToken.Role
		_, error := newUser.EntityGetRole(&bodyRole)
		if error.Type == "error_result_01" {
			helpers.APIResponse(ctx, "Role account is not found", error.Code, nil)
			return
		}
		// if !roles[accessToken.Role] {
		// 	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": "403", "message": "Role Access Not Allowed"})
		// }

		ctx.Next()
	})
}
