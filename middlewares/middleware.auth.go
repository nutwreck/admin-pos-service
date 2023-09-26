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

func AuthToken(db *gorm.DB) gin.HandlerFunc {
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

		//CEK USER STATUS
		newUser := repositories.NewRepositoryUser(db)
		var bodyUser schemes.User
		bodyUser.ID = accessToken.ID
		_, error := newUser.EntityGetUser(&bodyUser)
		if error.Type == "error_result_01" {
			helpers.APIResponse(ctx, "User account is not never registered or active", error.Code, nil)
			return
		}

		ctx.Set("user", accessToken)
		ctx.Next()
	})
}
