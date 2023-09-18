package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/helpers"
	"github.com/nutwreck/admin-pos-service/pkg"
)

func AuthRole(roles map[string]bool) gin.HandlerFunc {
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

		if !roles[accessToken.Role] {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": "403", "message": "Role Access Not Allowed"})
		}

		ctx.Next()
	})
}
