package pkg

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/schemes"
)

func Sign(configs *schemes.JWtMetaRequest) (string, time.Time, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * (24 * 60) * configs.Options.ExpiredAt)
	claims := jwt.MapClaims{}
	claims["jwt"] = configs.Data
	claims["exp"] = expiredAt.Unix()
	claims["audience"] = configs.Options.Audience
	claims["authorization"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(configs.SecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, expiredAt.Local(), err
	}

	return accessToken, expiredAt.Local(), nil
}

func VerifyToken(accessToken, SecretPublicKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretPublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func GenerateRefreshTokenFromClaims(claims jwt.MapClaims, jwtSecretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return constants.EMPTY_VALUE, err
	}

	return tokenString, nil
}

func ConvertToken(tokenString string) (*schemes.SchemeJWTConvert, error) {
	var (
		jwtSecretKey = []byte(GodotEnv("JWT_SECRET_KEY"))
		result       schemes.SchemeJWTConvert
	)

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		audience := claims["audience"].(string)
		authorization := claims["authorization"].(bool)
		//expiration := claims["exp"].(float64)
		jwtData := claims["jwt"].(map[string]interface{})
		email := jwtData["email"].(string)
		id := jwtData["id"].(string)
		role := jwtData["role"].(string)

		// Format Unix timestamp to time.Time
		// expirationTime := time.Unix(int64(expiration), 0)

		if audience != GodotEnv("JWT_AUD") {
			log.Println("Error:", "Invalid token claims - AUD")
			return nil, errors.New("invalid token claims - AUD")
		}

		if !authorization {
			log.Println("Error:", "Invalid token claims - Authorization False")
			return nil, errors.New("invalid token claims - Authorization False")
		}

		result.ID = id
		result.Email = email
		result.Role = role

		return &result, nil
	} else {
		log.Println("Error:", "Invalid token claims")
		return nil, errors.New("invalid token claims")
	}
}
