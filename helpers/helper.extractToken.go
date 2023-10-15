package helpers

import (
	"encoding/json"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/pkg"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type AccessToken struct {
	ID       string `json:"ucode"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Merchant string `json:"merchant"`
}

func ExtractToken(claimsToken *jwt.Token) AccessToken {
	data := claimsToken.Claims.(jwt.MapClaims)
	parseToken := make(map[string]interface{})
	var extractToken AccessToken

	for _, v := range data {
		stringify, _ := json.Marshal(&v)
		json.Unmarshal([]byte(stringify), &parseToken)

	}

	stringify, _ := json.Marshal(&parseToken)
	json.Unmarshal([]byte(stringify), &extractToken)

	return extractToken
}

func GetDataTokenBearer(JWTBearer string) (schemes.User, bool) {
	var result schemes.User

	if JWTBearer == constants.EMPTY_VALUE {
		return result, constants.FALSE_VALUE
	}

	token := strings.Split(JWTBearer, " ")

	if len(token) < 2 {
		return result, constants.FALSE_VALUE
	}

	decodeToken, err := pkg.VerifyToken(strings.TrimSpace(token[1]), pkg.GodotEnv("JWT_SECRET_KEY"))

	if err != nil {
		return result, constants.FALSE_VALUE
	}

	accessToken := ExtractToken(decodeToken)
	result.ID = accessToken.ID
	result.Email = accessToken.Email
	result.MerchantID = accessToken.Merchant
	result.RoleID = accessToken.Role

	return result, constants.TRUE_VALUE
}
