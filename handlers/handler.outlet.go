package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/helpers"
	"github.com/nutwreck/admin-pos-service/pkg"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type handlerOutlet struct {
	outlet entities.EntityOutlet
}

func NewHandlerOutlet(outlet entities.EntityOutlet) *handlerOutlet {
	return &handlerOutlet{outlet: outlet}
}

/**
* ======================================
* Handler Ping Status Outlet Teritory
*=======================================
 */

func (h *handlerOutlet) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Outlet", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Outlet Teritory
*======================================
 */
// CreateMasterOutlet godoc
// @Summary		Create Master Outlet
// @Description	Create Master Outlet
// @Tags		Master Outlet
// @Accept		json
// @Produce		json
// @Param		outlet body schemes.OutletRequest true "Create Master Outlet"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/outlet/create [post]
func (h *handlerOutlet) HandlerCreate(ctx *gin.Context) {
	var body schemes.Outlet
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorOutlet(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.outlet.EntityCreate(&body)

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Outlet phone number already taken", error.Code, nil)
		return
	}

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new Outlet failed", error.Code, nil)
		return
	}

	if error.Type == "error_create_04" {
		helpers.APIResponse(ctx, "The main outlet is predetermined", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Outlet successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Result Outlet By ID Teritory
*=======================================
 */
// GetListMasterOutlet godoc
// @Summary		Get List Master Outlet
// @Description	Get List Master Outlet
// @Tags		Master Outlet
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : outlet.id, outlet.name, outlet.phone, outlet.created_at, outlet.active, outlet.is_primary, merchant.id, merchant.name, default is merchant.name ASC | If you don't want to use it, fill it blank"
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
// @Router /api/v1/master/outlet/results [get]
func (h *handlerOutlet) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.Outlet
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

	res, totalData, error := h.outlet.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Outlet data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Outlet data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ======================================
* Handler Delete Outlet By ID Teritory
*=======================================
 */
// GetDeleteMasterOutlet godoc
// @Summary		Get Delete Master Outlet
// @Description	Get Delete Master Outlet
// @Tags		Master Outlet
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Master Outlet"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/outlet/delete [delete]
func (h *handlerOutlet) HandlerDelete(ctx *gin.Context) {
	var body schemes.Outlet
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorOutlet(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.outlet.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Outlet data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Outlet data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */
// GetUpdateMasterOutlet godoc
// @Summary		Get Update Master Outlet
// @Description	Get Update Master Outlet
// @Tags		Master Outlet
// @Accept		json
// @Produce		json
// @Param		id query string true "Update Master Outlet"
// @Param		outlet body schemes.OutletRequest true "Update Master Outlet"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/outlet/update [put]
func (h *handlerOutlet) HandlerUpdate(ctx *gin.Context) {
	var body schemes.Outlet
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorOutlet(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.outlet.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Outlet data failed for this id %s", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_03" {
		helpers.APIResponse(ctx, "The main outlet is predetermined", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Outlet data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Outlet
*=======================================
 */

func ValidatorOutlet(ctx *gin.Context, input schemes.Outlet, Type string) (interface{}, int) {
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
					Tag:     "required",
					Field:   "Phone",
					Message: "Phone is required on body",
				},
				{
					Tag:     "numeric",
					Field:   "Phone",
					Message: "Phone must be number",
				},
				{
					Tag:     "required",
					Field:   "Address",
					Message: "Address is required on body",
				},
				{
					Tag:     "max",
					Field:   "Address",
					Message: "Address maximal 1000 character",
				},
				{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "MerchantID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "MerchantID value must be uuid",
				},
				{
					Tag:     "max",
					Field:   "Description",
					Message: "Description maximal 1000 character",
				},
				{
					Tag:     "required",
					Field:   "IsPrimary",
					Message: "Primary is required on body",
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
					Tag:     "required",
					Field:   "Phone",
					Message: "Phone is required on body",
				},
				{
					Tag:     "numeric",
					Field:   "Phone",
					Message: "Phone must be number",
				},
				{
					Tag:     "required",
					Field:   "Address",
					Message: "Address is required on body",
				},
				{
					Tag:     "max",
					Field:   "Address",
					Message: "Address maximal 1000 character",
				},
				{
					Tag:     "required",
					Field:   "MerchantID",
					Message: "MerchantID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "MerchantID",
					Message: "MerchantID value must be uuid",
				},
				{
					Tag:     "max",
					Field:   "Description",
					Message: "Description maximal 1000 character",
				},
				{
					Tag:     "required",
					Field:   "IsPrimary",
					Message: "Primary is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
