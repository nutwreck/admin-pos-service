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

type handleMenuDetailFunction struct {
	menuDetailFunction entities.EntityMenuDetailFunction
}

func NewHandlerMenuDetailFunction(menuDetailFunction entities.EntityMenuDetailFunction) *handleMenuDetailFunction {
	return &handleMenuDetailFunction{menuDetailFunction: menuDetailFunction}
}

/**
* ========================================================
* Handler Ping Status Master Menu Detail Function Teritory
*=========================================================
 */

func (h *handleMenuDetailFunction) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master Menu Detail Function", http.StatusOK, nil)
}

/**
* =======================================================
* Handler Create New Master Menu Detail Function Teritory
*========================================================
 */
// CreateMasterMenuDetailFunction godoc
// @Summary		Create Master Menu Detail Function
// @Description	Create Master Menu Detail Function
// @Tags		Master Menu Detail Function
// @Accept		json
// @Produce		json
// @Param		menu body schemes.MenuDetailFunctionRequest true "Create Master Menu Detail Function"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu-detail-function/create [post]
func (h *handleMenuDetailFunction) HandlerCreate(ctx *gin.Context) {
	var body schemes.MenuDetailFunction
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMenuDetailFunction(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menuDetailFunction.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master Menu Detail Function name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master Menu Detail Function failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master Menu Detail Function successfully", http.StatusCreated, nil)
}

/**
* ========================================================
* Handler Results All Master Menu Detail Function Teritory
*=========================================================
 */
// GetListMasterMenuDetailFunction godoc
// @Summary		Get List Master Menu Detail Function
// @Description	Get List Master Menu Detail Function
// @Tags		Master Menu Detail Function
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : menudetailfunction.id, menudetailfunction.name, menu.id, menu.name, menudetail.id, menudetail.name, menudetailfunction.active, default is menu.name ASC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
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
// @Router /api/v1/master/menu-detail-function/results [get]
func (h *handleMenuDetailFunction) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.MenuDetailFunction
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
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		body.ID = idParam
	}

	res, totalData, error := h.menuDetailFunction.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master Menu Detail Function data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master Menu Detail Function data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* =========================================================
* Handler Delete Master Menu Detail Function By ID Teritory
*==========================================================
 */
// GetDeleteMasterMenuDetailFunction godoc
// @Summary		Get Delete Master Menu Detail Function
// @Description	Get Delete Master Menu Detail Function
// @Tags		Master Menu Detail Function
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Master Menu Detail Function"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu-detail-function/delete/{id} [delete]
func (h *handleMenuDetailFunction) HandlerDelete(ctx *gin.Context) {
	var body schemes.MenuDetailFunction
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorMenuDetailFunction(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.menuDetailFunction.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Menu Detail Function data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Menu Detail Function data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Menu Detail Function data for this id %s success", id), http.StatusOK, res)
}

/**
* =========================================================
* Handler Update Master Menu Detail Function By ID Teritory
*==========================================================
 */
// GetUpdateMasterMenuDetailFunction godoc
// @Summary		Get Update Master Menu Detail Function
// @Description	Get Update Master Menu Detail Function
// @Tags		Master Menu Detail Function
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Master Menu Detail Function"
// @Param		menu body schemes.MenuDetailFunctionRequest true "Update Master Menu Detail Function"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu-detail-function/update/{id} [put]
func (h *handleMenuDetailFunction) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.MenuDetailFunction
		activeGet = false
	)
	id := ctx.Param("id")
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.MenuID = ctx.PostForm("menu_id")
	body.MenuDetailID = ctx.PostForm("menu_detail_id")
	body.Link = ctx.PostForm("link")
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

	errors, code := ValidatorMenuDetailFunction(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menuDetailFunction.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Menu Detail Function data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master Menu Detail Function data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master Menu Detail Function data success for this id %s", id), http.StatusOK, nil)
}

/**
* =========================================================
*  All Validator User Input For Master Menu Detail Function
*==========================================================
 */

func ValidatorMenuDetailFunction(ctx *gin.Context, input schemes.MenuDetailFunction, Type string) (interface{}, int) {
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
					Field:   "MenuID",
					Message: "Menu ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MenuID",
					Message: "Menu ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "Link",
					Message: "Link is required on body",
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
					Field:   "MenuID",
					Message: "Menu ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MenuID",
					Message: "Menu ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MenuDetailID",
					Message: "Menu Detail ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "Link",
					Message: "Link is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
