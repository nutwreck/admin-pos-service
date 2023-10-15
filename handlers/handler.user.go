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
	"github.com/nutwreck/admin-pos-service/constants"
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
* Handler Add New Account
*======================================-
 */

// AddNewUser godoc
// @Summary		Add New User
// @Description	add by json user
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		user body schemes.UserRequest true "Add New User"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/auth/add-user [post]
func (h *handlerUser) HandlerAddUser(ctx *gin.Context) {
	var (
		body     schemes.User
		bodyRole schemes.Role
	)

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "add_user")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	bearer := ctx.GetHeader("Authorization")
	checkMerchant, errToken := helpers.GetDataTokenBearer(bearer)
	if !errToken {
		helpers.APIResponse(ctx, "Get data token is not valid", http.StatusBadRequest, nil)
		return
	}

	//Cek role tipe sys atau user - jika sys (dari admin) => merchant id bebas dari dashboard, jika user (dari user) => harus sesuai merchant id user yang daftarin user lainnya
	bodyRole.ID = body.RoleID
	checkRole, errRole := h.user.EntityGetRole(&bodyRole)
	if errRole.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Role account is not found", errRole.Code, nil)
		return
	}

	if checkRole.Type == constants.ROLE_USER {
		if body.MerchantID != checkMerchant.MerchantID {
			helpers.APIResponse(ctx, "Merchant is not valid", http.StatusBadRequest, nil)
			return
		}
	}

	_, error := h.user.EntityAddUser(&body)

	if error.Type == "error_add_user_01" {
		helpers.APIResponse(ctx, "Email already taken", error.Code, nil)
		return
	}

	if error.Type == "error_add_user_02" {
		helpers.APIResponse(ctx, "Add new user account failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Add new user account success", http.StatusOK, nil)
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
// @Param		user body schemes.LoginUser true "Login User"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Router /api/v1/auth/login [post]
func (h *handlerUser) HandlerLogin(ctx *gin.Context) {
	var (
		body             schemes.User
		bodyRole         schemes.Role
		bodyMerchant     schemes.Merchant
		bodyUserOutlet   schemes.UserOutlet
		resGetOutletUser []schemes.GetOutletUser
		result           schemes.GetUser
	)
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

	if !*res.Active {
		helpers.APIResponse(ctx, "User account is disabled", http.StatusForbidden, nil)
		return
	}

	if error.Type == "error_login_02" {
		helpers.APIResponse(ctx, "Email or Password is wrong", error.Code, nil)
		return
	}

	accessToken, expiredAt, errorJwt := pkg.Sign(&schemes.JWtMetaRequest{
		Data:      gin.H{"ucode": res.ID, "email": res.Email, "role": res.RoleID, "merchant": res.MerchantID},
		SecretKey: pkg.GodotEnv("JWT_SECRET_KEY"),
		Options:   schemes.JwtMetaOptions{Audience: pkg.GodotEnv("JWT_AUD"), ExpiredAt: configs.DayExpiredJWT},
	})

	if errorJwt != nil {
		helpers.APIResponse(ctx, "Generate access token failed", http.StatusBadRequest, nil)
		return
	}

	// Get Role
	bodyRole.ID = res.RoleID
	resGetRole, error := h.user.EntityGetRole(&bodyRole)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Role account in not found", error.Code, nil)
		return
	}
	if !*resGetRole.Active {
		helpers.APIResponse(ctx, "Role account is disabled", http.StatusForbidden, nil)
		return
	}

	// Get Merchant
	bodyMerchant.ID = res.MerchantID
	resGetMerchant, error := h.user.EntityGetMerchant(&bodyMerchant)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Merchant in not found", error.Code, nil)
		return
	}
	if !*resGetMerchant.Active {
		helpers.APIResponse(ctx, "Merchant is disabled", http.StatusForbidden, nil)
		return
	}

	// Get User Outlet
	bodyUserOutlet.UserID = res.ID
	resGetUserOutlet, _ := h.user.EntityGetUserOutlet(&bodyUserOutlet)

	result.Active = *res.Active
	result.Email = res.Email
	result.Name = res.Name
	result.ID = res.ID
	result.Role = schemes.GetRole{
		ID:   resGetRole.ID,
		Name: resGetRole.Name,
		Type: resGetRole.Type,
	}
	result.Merchant = schemes.GetMerchant{
		ID:          resGetMerchant.ID,
		Name:        resGetMerchant.Name,
		Phone:       resGetMerchant.Phone,
		Address:     resGetMerchant.Address,
		Logo:        configs.AccessFile + resGetMerchant.Logo,
		Description: resGetMerchant.Description,
		CreatedAt:   resGetMerchant.CreatedAt,
		Active:      resGetMerchant.Active,
	}
	if len(*resGetUserOutlet) > 0 {
		for _, userOutlet := range *resGetUserOutlet {
			userOutletData := schemes.GetOutletUser{
				ID:          userOutlet.OutletID,
				Name:        userOutlet.OutletName,
				Phone:       userOutlet.OutletPhone,
				Address:     userOutlet.OutletAddress,
				Description: userOutlet.OutletDescription,
				CreatedAt:   userOutlet.OutletCreatedAt,
				Active:      userOutlet.OutletActive,
			}

			resGetOutletUser = append(resGetOutletUser, userOutletData)
		}

		result.Outlet = resGetOutletUser
	}

	helpers.APIResponse(ctx, "Login successfully", http.StatusOK, gin.H{"accessToken": accessToken, "expiredAt": expiredAt, "user": result})
}

// RefreshTokenUser godoc
// @Summary		Refresh Token User
// @Description	Refresh Token user by Header Bearer Token
// @Tags		User
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
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
// @Param		User body schemes.UpdateUserExample true "Update User"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/auth/update [put]
func (h *handlerUser) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.UpdateUser
		bodyUser  schemes.User
		result    schemes.User
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
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}
	id = resGetUser.ID

	//Get Body
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.OldPassword = ctx.PostForm("old_password")
	body.NewPassword = ctx.PostForm("new_password")
	body.DataPassword = resGetUser.Password
	body.RoleID = ctx.PostForm("role_id")
	body.MerchantID = ctx.PostForm("merchant_id")
	activeStr := ctx.PostForm("active")
	if activeStr == "true" {
		activeGet = constants.TRUE_VALUE
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
	result.Name = resultUpdate.Name
	result.ID = resGetUser.ID
	result.RoleID = resultUpdate.RoleID

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
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/auth/data-user [get]
func (h *handlerUser) HandleDataUser(ctx *gin.Context) {
	var (
		bodyUser         schemes.User
		bodyRole         schemes.Role
		bodyMerchant     schemes.Merchant
		bodyUserOutlet   schemes.UserOutlet
		resGetOutletUser []schemes.GetOutletUser
		result           schemes.GetUser
	)

	bearer := ctx.GetHeader("Authorization")
	token := strings.Split(bearer, " ")
	existingToken := strings.TrimSpace(token[1])
	convertToken, err := pkg.ConvertToken(existingToken)
	if err != nil {
		helpers.APIResponse(ctx, "Failed Convert Token!", http.StatusInternalServerError, nil)
		return
	}

	//Get User
	bodyUser.ID = convertToken.ID
	resGetUser, error := h.user.EntityGetUser(&bodyUser)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}

	// Get Role
	bodyRole.ID = resGetUser.RoleID
	resGetRole, error := h.user.EntityGetRole(&bodyRole)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Role account in not found", error.Code, nil)
		return
	}
	if !*resGetRole.Active {
		helpers.APIResponse(ctx, "Role account is disabled", http.StatusForbidden, nil)
		return
	}

	// Get Merchant
	bodyMerchant.ID = resGetUser.MerchantID
	resGetMerchant, error := h.user.EntityGetMerchant(&bodyMerchant)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Merchant in not found", error.Code, nil)
		return
	}
	if !*resGetMerchant.Active {
		helpers.APIResponse(ctx, "Merchant is disabled", http.StatusForbidden, nil)
		return
	}

	// Get User Outlet
	bodyUserOutlet.UserID = resGetUser.ID
	resGetUserOutlet, _ := h.user.EntityGetUserOutlet(&bodyUserOutlet)

	result.Active = *resGetUser.Active
	result.Email = resGetUser.Email
	result.Name = resGetUser.Name
	result.ID = resGetUser.ID
	result.Role = schemes.GetRole{
		ID:   resGetRole.ID,
		Name: resGetRole.Name,
		Type: resGetRole.Type,
	}
	result.Merchant = schemes.GetMerchant{
		ID:          resGetMerchant.ID,
		Name:        resGetMerchant.Name,
		Phone:       resGetMerchant.Phone,
		Address:     resGetMerchant.Address,
		Logo:        configs.AccessFile + resGetMerchant.Logo,
		Description: resGetMerchant.Description,
		CreatedAt:   resGetMerchant.CreatedAt,
		Active:      resGetMerchant.Active,
	}
	if len(*resGetUserOutlet) > 0 {
		for _, userOutlet := range *resGetUserOutlet {
			userOutletData := schemes.GetOutletUser{
				ID:          userOutlet.OutletID,
				Name:        userOutlet.OutletName,
				Phone:       userOutlet.OutletPhone,
				Address:     userOutlet.OutletAddress,
				Description: userOutlet.OutletDescription,
				CreatedAt:   userOutlet.OutletCreatedAt,
				Active:      userOutlet.OutletActive,
			}

			resGetOutletUser = append(resGetOutletUser, userOutletData)
		}

		result.Outlet = resGetOutletUser
	}

	helpers.APIResponse(ctx, "User data already to use", http.StatusOK, result)
}

/**
* ======================================
*  All Validator User Input For User
*=======================================
 */

func ValidatorUser(ctx *gin.Context, input schemes.User, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "add_user" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Name",
					Message: "Name must be lowercase",
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
					Field:   "RoleID",
					Message: "Role ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "RoleID",
					Message: "Role ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "Merchant ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "Merchant ID must be uuid",
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

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}

func ValidatorUpdateUser(ctx *gin.Context, input schemes.UpdateUser) (interface{}, int) {
	schema := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Name",
				Message: "Name is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "Name",
				Message: "Name must be lowercase",
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
				Field:   "RoleID",
				Message: "Role ID is required on body",
			},
			{
				Tag:     "uuid",
				Field:   "RoleID",
				Message: "Role ID must be uuid",
			},
			{
				Tag:     "required",
				Field:   "MerchantID",
				Message: "Merchant ID is required on body",
			},
			{
				Tag:     "uuid",
				Field:   "MerchantID",
				Message: "Merchant ID must be uuid",
			},
		},
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
