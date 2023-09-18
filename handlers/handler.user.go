package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/helpers"
	"github.com/nutwreck/admin-pos-service/pkg"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type handlerUser struct {
	user entities.EntityUser
}

func NewHandlerUser(user entities.EntityUser) *handlerUser {
	return &handlerUser{user: user}
}

/**
* ==================================
* Handler Ping User Status
*==================================
 */

func (h *handlerUser) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping User", http.StatusOK, nil)
}

/**
* ======================================
* Handler Register New Account
*======================================-
 */

// RegisterUser godoc
// @Summary		Register User
// @Description	add by json user
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		user body schemes.SchemeUserRequest true "Add User"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/auth/register [post]
func (h *handlerUser) HandlerRegister(ctx *gin.Context) {
	var body schemes.SchemeUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "register")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.user.EntityRegister(&body)

	if error.Type == "error_register_01" {
		helpers.APIResponse(ctx, "Email already taken", error.Code, nil)
		return
	}

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Register new user account failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Register new user account success", http.StatusOK, nil)
}

/**
* =================================
* Handler Login Auth Account
*==================================
 */

// LoginUser godoc
// @Summary		Login User
// @Description	login user
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		user body schemes.SchemeLoginUser true "Login User"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/auth/login [post]
func (h *handlerUser) HandlerLogin(ctx *gin.Context) {
	var body schemes.SchemeUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "login")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.user.EntityLogin(&body)

	if error.Type == "error_login_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}

	if error.Type == "error_login_02" {
		helpers.APIResponse(ctx, "Email or Password is wrong", error.Code, nil)
		return
	}

	accessToken, expiredAt, errorJwt := pkg.Sign(&schemes.JWtMetaRequest{
		Data:      gin.H{"id": res.ID, "email": res.Email, "role": res.Role},
		SecretKey: pkg.GodotEnv("JWT_SECRET_KEY"),
		Options:   schemes.JwtMetaOptions{Audience: pkg.GodotEnv("JWT_AUD"), ExpiredAt: configs.DayExpiredJWT},
	})

	if errorJwt != nil {
		helpers.APIResponse(ctx, "Generate access token failed", http.StatusBadRequest, nil)
		return
	}

	helpers.APIResponse(ctx, "Login successfully", http.StatusOK, gin.H{"accessToken": accessToken, "expiredAt": expiredAt})
}

// RefreshTokenUser godoc
// @Summary		Refresh Token User
// @Description	Refresh Token user by Header Bearer Token
// @Tags		User
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/auth/refresh-token [get]
func (h *handlerUser) HandlerRefreshToken(ctx *gin.Context) {
	var jwtSecretKey = []byte(pkg.GodotEnv("JWT_SECRET_KEY"))
	bearer := ctx.GetHeader("Authorization")
	token := strings.Split(bearer, " ")
	existingToken := strings.TrimSpace(token[1])
	refreshToken := existingToken

	// Parse the refresh token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		helpers.APIResponse(ctx, "Invalid refreshing token", http.StatusBadRequest, err.Error())
		return
	}

	expiredAt := time.Now().Add(time.Duration(time.Minute) * (24 * 60) * configs.DayExpiredJWT)

	// Update the expired time of the refresh token
	claims["exp"] = expiredAt.Unix()

	// Generate a new refresh token
	newRefreshToken, err := pkg.GenerateRefreshTokenFromClaims(claims, jwtSecretKey)
	if err != nil {
		helpers.APIResponse(ctx, "Failed to generate new refresh token", http.StatusInternalServerError, nil)
		return
	}

	helpers.APIResponse(ctx, "Refresh Token successfully", http.StatusOK, gin.H{"accessToken": newRefreshToken, "expiredAt": expiredAt.Local()})
}

/**
* ===================
* Handler Update User
*====================
 */
// GetUpdateUser godoc
// @Summary		Update User
// @Description	Update User. If the user wants to change the password then it is mandatory to input the old and new passwords, If not then leave it blank.
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		User body schemes.SchemeUpdateUserExample true "Update User"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/auth/update [put]
func (h *handlerUser) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemeUpdateUser
		bodyUser  schemes.SchemeUser
		result    schemes.SchemeUser
		activeGet = false
		id        string
	)

	bearer := ctx.GetHeader("Authorization")
	token := strings.Split(bearer, " ")
	existingToken := strings.TrimSpace(token[1])
	convertToken, err := pkg.ConvertToken(existingToken)
	if err != nil {
		helpers.APIResponse(ctx, "Failed Convert Token!", http.StatusInternalServerError, nil)
		return
	}

	//Validasi User
	bodyUser.ID = convertToken.ID
	resGetUser, error := h.user.EntityGetUser(&bodyUser)
	if error.Type == "error_login_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}
	id = resGetUser.ID

	//Get Body
	body.ID = id
	body.FirstName = ctx.PostForm("first_name")
	body.LastName = ctx.PostForm("last_name")
	body.OldPassword = ctx.PostForm("old_password")
	body.NewPassword = ctx.PostForm("new_password")
	body.DataPassword = resGetUser.Password
	body.Role = ctx.PostForm("role")
	activeStr := ctx.PostForm("active")
	if activeStr == "true" {
		activeGet = true
	}
	body.Active = &activeGet

	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUpdateUser(ctx, body)
	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	resultUpdate, error := h.user.EntityUpdate(&body)
	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update User data failed for this id %s", id), error.Code, nil)
		return
	}
	if error.Type == "error_update_03" {
		helpers.APIResponse(ctx, "If the old password is entered, the new password is required!", error.Code, nil)
		return
	}
	if error.Type == "error_update_04" {
		helpers.APIResponse(ctx, "If the new password is entered, the old password is required!", error.Code, nil)
		return
	}
	if error.Type == "error_update_05" {
		helpers.APIResponse(ctx, "Old Password is wrong", error.Code, nil)
		return
	}

	result.Active = activeGet
	result.Email = resGetUser.Email
	result.FirstName = resultUpdate.FirstName
	result.LastName = resultUpdate.LastName
	result.ID = resGetUser.ID
	result.Role = resultUpdate.Role

	helpers.APIResponse(ctx, fmt.Sprintf("Update User data success for this id %s", id), http.StatusCreated, result)
}

/**
* =================
* Handler Data User
*==================
 */
// GetDataUser godoc
// @Summary		Get Data User
// @Description	Get Data User
// @Tags		User
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/auth/data-user [get]
func (h *handlerUser) HandleDataUser(ctx *gin.Context) {
	var (
		bodyUser schemes.SchemeUser
		result   schemes.SchemeUser
	)

	bearer := ctx.GetHeader("Authorization")
	token := strings.Split(bearer, " ")
	existingToken := strings.TrimSpace(token[1])
	convertToken, err := pkg.ConvertToken(existingToken)
	if err != nil {
		helpers.APIResponse(ctx, "Failed Convert Token!", http.StatusInternalServerError, nil)
		return
	}

	//Validasi User
	bodyUser.ID = convertToken.ID
	resGetUser, error := h.user.EntityGetUser(&bodyUser)
	if error.Type == "error_login_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}

	result.Active = *resGetUser.Active
	result.Email = resGetUser.Email
	result.FirstName = resGetUser.FirstName
	result.LastName = resGetUser.LastName
	result.ID = resGetUser.ID
	result.Role = resGetUser.Role

	helpers.APIResponse(ctx, "User data already to use", http.StatusOK, result)
}

/**
* ======================================
*  All Validator User Input For User
*=======================================
 */

func ValidatorUser(ctx *gin.Context, input schemes.SchemeUser, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "register" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "FirstName",
					Message: "FirstName is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "FirstName",
					Message: "FirstName must be lowercase",
				},
				{
					Tag:     "required",
					Field:   "LastName",
					Message: "LastName is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "LastName",
					Message: "LastName must be lowercase",
				},
				{
					Tag:     "required",
					Field:   "Email",
					Message: "Email is required on body",
				},
				{
					Tag:     "email",
					Field:   "Email",
					Message: "Email format is not valid",
				},
				{
					Tag:     "password",
					Field:   "Password",
					Message: "Password is required on body",
				},
				{
					Tag:     "gte",
					Field:   "Password",
					Message: "Password must be greater than equal 8 character",
				},
				{
					Tag:     "required",
					Field:   "Role",
					Message: "Role is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Role",
					Message: "Role must be lowercase",
				},
			},
		}
	}

	if Type == "login" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "Email",
					Message: "Email is required on body",
				},
				{
					Tag:     "email",
					Field:   "Email",
					Message: "Email format is not valid",
				},
				{
					Tag:     "required",
					Field:   "Password",
					Message: "Password is required on body",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "FirstName",
					Message: "FirstName is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "FirstName",
					Message: "FirstName must be lowercase",
				},
				{
					Tag:     "required",
					Field:   "LastName",
					Message: "LastName is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "LastName",
					Message: "LastName must be lowercase",
				},
				{
					Tag:     "gte",
					Field:   "OldPassword",
					Message: "Old Password must be greater than equal 8 character",
				},
				{
					Tag:     "gte",
					Field:   "NewPassword",
					Message: "New Password must be greater than equal 8 character",
				},
				{
					Tag:     "required",
					Field:   "Role",
					Message: "Role is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Role",
					Message: "Role must be lowercase",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
func ValidatorUpdateUser(ctx *gin.Context, input schemes.SchemeUpdateUser) (interface{}, int) {
	schema := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "FirstName",
				Message: "FirstName is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "FirstName",
				Message: "FirstName must be lowercase",
			},
			{
				Tag:     "required",
				Field:   "LastName",
				Message: "LastName is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "LastName",
				Message: "LastName must be lowercase",
			},
			{
				Tag:     "gte",
				Field:   "OldPassword",
				Message: "Old Password must be greater than equal 8 character",
			},
			{
				Tag:     "gte",
				Field:   "NewPassword",
				Message: "New Password must be greater than equal 8 character",
			},
			{
				Tag:     "required",
				Field:   "Role",
				Message: "Role is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "Role",
				Message: "Role must be lowercase",
			},
		},
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
