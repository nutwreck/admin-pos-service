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

type handleMappingRoleMenu struct {
	mappingRoleMenu entities.EntityMappingRoleMenu
}

func NewHandlerMappingRoleMenu(mappingRoleMenu entities.EntityMappingRoleMenu) *handleMappingRoleMenu {
	return &handleMappingRoleMenu{mappingRoleMenu: mappingRoleMenu}
}

/**
* =====================================================
* Handler Ping Status Mapping Role Menu Teritory
*======================================================
 */

func (h *handleMappingRoleMenu) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Mapping Role Menu", http.StatusOK, nil)
}

/**
* ====================================================
* Handler Create New Mapping Role Menu Teritory
*=====================================================
 */
// CreateMappingRoleMenu godoc
// @Summary		Create Mapping Role Menu
// @Description	Create Mapping Role Menu
// @Tags		Mapping Role Menu
// @Accept		json
// @Produce		json
// @Param		mappingrolemenu body schemes.MappingRoleMenuRequest true "Create Mapping Role Menu"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu/create [post]
func (h *handleMappingRoleMenu) HandlerCreate(ctx *gin.Context) {
	var body schemes.MappingRoleMenu
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMappingRoleMenu(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.mappingRoleMenu.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Mapping Role Menu name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Mapping Role Menu failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Mapping Role Menu successfully", http.StatusCreated, nil)
}

/**
* ====================================================
* Handler Results All Mapping Role Menu Teritory
*=====================================================
 */
// GetListMappingRoleMenu godoc
// @Summary		Get List Mapping Role Menu
// @Description	Get List Mapping Role Menu
// @Tags		Mapping Role Menu
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : mappingrolemenu.id, mappingrolemenu.name, mappingrolemenu.active, mappingrolemenu.created_at, merchant.id, merchant.name, role.id, role.name, menu.id, menu.name, menudetail.id, menudetail.name, menudetailfunction.id, menudetailfunction.name, default is mappingrolemenu.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param name query string false "Search by name using LIKE pattern"
// @Param id query string false "Search by ID"
// @Success 200 {object} schemes.ResponsesPagination
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu/results [get]
func (h *handleMappingRoleMenu) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.MappingRoleMenu
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
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		body.ID = idParam
	}

	res, totalData, error := h.mappingRoleMenu.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Mapping Role Menu data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Mapping Role Menu data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* =====================================================
* Handler Delete Mapping Role Menu By ID Teritory
*======================================================
 */
// GetDeleteMappingRoleMenu godoc
// @Summary		Get Delete Mapping Role Menu
// @Description	Get Delete Mapping Role Menu
// @Tags		Mapping Role Menu
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Mapping Role Menu"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu/delete [delete]
func (h *handleMappingRoleMenu) HandlerDelete(ctx *gin.Context) {
	var body schemes.MappingRoleMenu
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorMappingRoleMenu(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.mappingRoleMenu.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Mapping Role Menu data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Mapping Role Menu data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Mapping Role Menu data for this id %s success", id), http.StatusOK, res)
}

/**
* =====================================================
* Handler Update Mapping Role Menu By ID Teritory
*======================================================
 */
// GetUpdateMappingRoleMenu godoc
// @Summary		Get Update Mapping Role Menu
// @Description	Get Update Mapping Role Menu
// @Tags		Mapping Role Menu
// @Accept		json
// @Produce		json
// @Param		id query string true "Update Mapping Role Menu"
// @Param		mappingrolemenu body schemes.MappingRoleMenuRequest true "Update Mapping Role Menu"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/mapping-role-menu/update [put]
func (h *handleMappingRoleMenu) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.MappingRoleMenu
		activeGet = false
	)
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.RoleID = ctx.PostForm("role_id")
	body.MenuID = ctx.PostForm("menu_id")
	body.MenuDetailID = ctx.PostForm("menu_detail_id")
	body.MenuDetailFunctionID = ctx.PostForm("menu_detail_function_id")
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

	errors, code := ValidatorMappingRoleMenu(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.mappingRoleMenu.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Mapping Role Menu data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Mapping Role Menu data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Mapping Role Menu data success for this id %s", id), http.StatusOK, nil)
}

/**
* ==============================================
*  All Validator User Input For Mapping Role Menu
*===============================================
 */

func ValidatorMappingRoleMenu(ctx *gin.Context, input schemes.MappingRoleMenu, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
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
					Tag:     "max",
					Field:   "Name",
					Message: "Name maximal 200 character",
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
					Field:   "RoleID",
					Message: "Role ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "RoleID",
					Message: "Role ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuID",
					Message: "Menu ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MenuID",
					Message: "Menu ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuDetailFunctionID",
					Message: "Menu Detail Function ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MenuDetailFunctionID",
					Message: "Menu Detail Function ID must be uuid",
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
					Field:   "Name",
					Message: "Name is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Name",
					Message: "Name must be lowercase",
				},
				{
					Tag:     "max",
					Field:   "Name",
					Message: "Name maximal 200 character",
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
					Field:   "RoleID",
					Message: "Role ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "RoleID",
					Message: "Role ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuID",
					Message: "Menu ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MenuID",
					Message: "Menu ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuDetailFunctionID",
					Message: "Menu Detail Function ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "MenuDetailFunctionID",
					Message: "Menu Detail Function ID must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
