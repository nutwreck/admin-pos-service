package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/helpers"
	"github.com/nutwreck/admin-pos-service/pkg"
	"github.com/nutwreck/admin-pos-service/schemes"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handlerMappingRoleMenuUser struct {
	mappingRoleMenuUser entities.EntityMappingRoleMenuUser
}

func NewHandlerMappingRoleMenuUser(mappingRoleMenuUser entities.EntityMappingRoleMenuUser) *handlerMappingRoleMenuUser {
	return &handlerMappingRoleMenuUser{mappingRoleMenuUser: mappingRoleMenuUser}
}

/**
* =============================================
* Handler Ping Status Mapping Role Menu User Teritory
*==============================================
 */

func (h *handlerMappingRoleMenuUser) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Mapping Role Menu User", http.StatusOK, nil)
}

/**
* ============================================
* Handler Create New Mapping Role Menu User Teritory
*=============================================
 */
// CreateMappingRoleMenuUser godoc
// @Summary		Create Mapping Role Menu User
// @Description	Create Mapping Role Menu User
// @Tags		Mapping Role Menu User
// @Accept		json
// @Produce		json
// @Param		mappingrolemenuuser body []schemes.MappingRoleMenuUserRequest true "Create Mapping Role Menu User"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu-user/create [post]
func (h *handlerMappingRoleMenuUser) HandlerCreate(ctx *gin.Context) {
	var body []schemes.MappingRoleMenuUser
	var datas []schemes.MappingRoleMenuUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	for _, input := range body {
		errors, code := ValidatorMappingRoleMenuUser(ctx, input, "create")
		if code > 0 {
			helpers.ErrorResponse(ctx, errors)
			return
		}
	}

	for _, req := range body {
		var mappRMU schemes.MappingRoleMenuUser
		mappRMU.UserID = req.UserID
		mappRMU.MappingRoleMenuID = req.MappingRoleMenuID
		mappRMU.MerchantID = req.MerchantID

		datas = append(datas, mappRMU)
	}

	_, error := h.mappingRoleMenuUser.EntityCreate(&datas)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Mapping Role Menu User name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Mapping Role Menu User failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Mapping Role Menu User successfully", http.StatusCreated, nil)
}

/**
* =============================================
* Handler Results All Mapping Role Menu User Teritory
*==============================================
 */
// GetListMappingRoleMenuUser godoc
// @Summary		Get List Mapping Role Menu User
// @Description	Get List Mapping Role Menu User
// @Tags		Mapping Role Menu User
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : mappingrolemenuuser.id, merchant.id, merchant.name, user.id, user.name, mappingrolemenu.id, mappingrolemenu.name, mappingrolemenuuser.created_at, default is mappingrolemenuuser.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param id query string false "Search by ID"
// @Success 200 {object} schemes.ResponsesPagination
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu-user/results [get]
func (h *handlerMappingRoleMenuUser) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.MappingRoleMenuUser
		reqPage       = configs.FirstPage
		reqPerPage    = configs.TotalPerPage
		pages         int
		perPages      int
		totalPagesDiv float64
		totalPages    int
		totalDatas    int
	)

	sortParam := ctx.DefaultQuery("sort", constants.EMPTY_VALUE)
	if sortParam != constants.EMPTY_VALUE {
		body.Sort = sortParam
	}
	pageParam := ctx.DefaultQuery("page", constants.EMPTY_VALUE)
	body.Page = reqPage
	if pageParam != constants.EMPTY_VALUE {
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			helpers.APIResponsePagination(ctx, "Convert Params Failed", http.StatusInternalServerError, nil, pages, perPages, totalPages, totalDatas)
			return
		}
		reqPage = page
		body.Page = page
	}
	perPageParam := ctx.DefaultQuery("perpage", constants.EMPTY_VALUE)
	body.PerPage = reqPerPage
	if perPageParam != constants.EMPTY_VALUE {
		perPage, err := strconv.Atoi(perPageParam)
		if err != nil {
			helpers.APIResponsePagination(ctx, "Convert Params Failed", http.StatusInternalServerError, nil, pages, perPages, totalPages, totalDatas)
			return
		}
		reqPerPage = perPage
		body.PerPage = perPage
	}
	merchantParam := ctx.DefaultQuery("merchant_id", constants.EMPTY_VALUE)
	if merchantParam != constants.EMPTY_VALUE {
		body.MerchantID = merchantParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		body.ID = idParam
	}

	res, totalData, error := h.mappingRoleMenuUser.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Mapping Role Menu User data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Mapping Role Menu User data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ==============================================
* Handler Delete Mapping Role Menu User By ID Teritory
*===============================================
 */
// GetDeleteMappingRoleMenuUser godoc
// @Summary		Get Delete Mapping Role Menu User
// @Description	Get Delete Mapping Role Menu User
// @Tags		Mapping Role Menu User
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Mapping Role Menu User"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu-user/delete [delete]
func (h *handlerMappingRoleMenuUser) HandlerDelete(ctx *gin.Context) {
	var body schemes.MappingRoleMenuUser
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorMappingRoleMenuUser(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.mappingRoleMenuUser.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Mapping Role Menu User data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Mapping Role Menu User data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Mapping Role Menu User data for this id %s success", id), http.StatusOK, res)
}

/**
* ==============================================
* Handler Update Mapping Role Menu User By ID Teritory
*===============================================
 */
// GetUpdateMappingRoleMenuUser godoc
// @Summary		Get Update Mapping Role Menu User
// @Description	Get Update Mapping Role Menu User
// @Tags		Mapping Role Menu User
// @Accept		json
// @Produce		json
// @Param		id query string true "Update Mapping Role Menu User"
// @Param		mappingrolemenuuser body schemes.MappingRoleMenuUserRequest true "Update Mapping Role Menu User"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu-user/update [put]
func (h *handlerMappingRoleMenuUser) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.MappingRoleMenuUser
		activeGet = false
	)
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id
	body.UserID = ctx.PostForm("user_id")
	body.MappingRoleMenuID = ctx.PostForm("mapping_role_menu_id")
	body.MerchantID = ctx.PostForm("merchant_id")
	activeStr := ctx.PostForm("active")
	if activeStr == "true" {
		activeGet = constants.TRUE_VALUE
	}
	body.Active = &activeGet

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMappingRoleMenuUser(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.mappingRoleMenuUser.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Mapping Role Menu User data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Mapping Role Menu User data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Mapping Role Menu User data success for this id %s", id), http.StatusOK, nil)
}

/**
* ==============================================
*  All Validator User Input For Mapping Role Menu User
*===============================================
 */

func ValidatorMappingRoleMenuUser(ctx *gin.Context, input schemes.MappingRoleMenuUser, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "Merchant ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "Merchant ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "UserID",
					Message: "User ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "UserID",
					Message: "User ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MappingRoleMenuID",
					Message: "Mapping Role Menu ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MappingRoleMenuID",
					Message: "Mapping Role Menu ID must be uuid",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "ID",
					Message: "ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "ID",
					Message: "ID must be uuid",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "ID",
					Message: "ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "ID",
					Message: "ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "Merchant ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "Merchant ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "UserID",
					Message: "User ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "UserID",
					Message: "User ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MappingRoleMenuID",
					Message: "Mapping Role Menu ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MappingRoleMenuID",
					Message: "Mapping Role Menu ID must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
