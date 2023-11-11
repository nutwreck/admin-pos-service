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

type handleUnitOfMeasurement struct {
	uom entities.EntityUnitOfMeasurement
}

func NewHandlerUnitOfMeasurement(uom entities.EntityUnitOfMeasurement) *handleUnitOfMeasurement {
	return &handleUnitOfMeasurement{uom: uom}
}

/**
* =============================================
* Handler Ping Status Master UOM Teritory
*==============================================
 */

func (h *handleUnitOfMeasurement) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master UOM", http.StatusOK, nil)
}

/**
* ============================================
* Handler Create New Master Menu Teritory
*=============================================
 */
// CreateMasterUOM godoc
// @Summary		Create Master UOM
// @Description	Create Master UOM
// @Tags		Master UOM
// @Accept		json
// @Produce		json
// @Param		uom body schemes.UnitOfMeasurementRequest true "Create Master UOM"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/uom/create [post]
func (h *handleUnitOfMeasurement) HandlerCreate(ctx *gin.Context) {
	var body schemes.UnitOfMeasurement
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUnitOfMeasurement(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.uom.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master UOM name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master UOM failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master UOM successfully", http.StatusCreated, nil)
}

/**
* =============================================
* Handler Results All Master UOM Teritory
*==============================================
 */
// GetListMasterUOM godoc
// @Summary		Get List Master UOM
// @Description	Get List Master UOM
// @Tags		Master UOM
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : uom.id, uom.name, uom.symbol, uom.conversion_factor, uom.active, merchant.id, merchant.name, uomType.id, uomType.name, uom.created_at, default is uom.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param uom_type_id query string false "Search by UOM Type"
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
// @Router /api/v1/master/uom/results [get]
func (h *handleUnitOfMeasurement) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.UnitOfMeasurement
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
	uomTypeParam := ctx.DefaultQuery("uom_type_id", constants.EMPTY_VALUE)
	if uomTypeParam != constants.EMPTY_VALUE {
		body.UOMTypeID = uomTypeParam
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

	res, totalData, error := h.uom.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master UOM data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master UOM data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ==============================================
* Handler Delete Master UOM By ID Teritory
*===============================================
 */
// GetDeleteMasterUOM godoc
// @Summary		Get Delete Master UOM
// @Description	Get Delete Master UOM
// @Tags		Master UOM
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Master UOM"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/uom/delete [delete]
func (h *handleUnitOfMeasurement) HandlerDelete(ctx *gin.Context) {
	var body schemes.UnitOfMeasurement
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorUnitOfMeasurement(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.uom.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master UOM data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master UOM data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master UOM data for this id %s success", id), http.StatusOK, res)
}

/**
* ==============================================
* Handler Update Master UOM By ID Teritory
*===============================================
 */
// GetUpdateMasterUOM godoc
// @Summary		Get Update Master UOM
// @Description	Get Update Master UOM
// @Tags		Master UOM
// @Accept		json
// @Produce		json
// @Param		id query string true "Update Master UOM"
// @Param		uom body schemes.UnitOfMeasurementRequest true "Update Master UOM"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/uom/update [put]
func (h *handleUnitOfMeasurement) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.UnitOfMeasurement
		activeGet = false
	)
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.UOMTypeID = ctx.PostForm("unit_of_measurement_id")
	body.Symbol = ctx.PostForm("symbol")
	conversionFactorStr := ctx.PostForm("conversion_factor")
	conversionFactor, _ := strconv.ParseFloat(conversionFactorStr, 64)
	body.ConversionFactor = conversionFactor
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

	errors, code := ValidatorUnitOfMeasurement(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.uom.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master UOM data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master UOM data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master UOM data success for this id %s", id), http.StatusOK, nil)
}

/**
* ==============================================
*  All Validator User Input For Master UOM
*===============================================
 */

func ValidatorUnitOfMeasurement(ctx *gin.Context, input schemes.UnitOfMeasurement, Type string) (interface{}, int) {
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
					Field:   "UOMTypeID",
					Message: "UOM Type ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "UOMTypeID",
					Message: "UOM Type ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "ConversionFactor",
					Message: "Conversion Factor is required on param",
				},
				{
					Tag:     "numeric",
					Field:   "ConversionFactor",
					Message: "Conversion Factor must be number",
				},
				{
					Tag:     "required",
					Field:   "Symbol",
					Message: "Symbol is required on param",
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
					Field:   "UOMTypeID",
					Message: "UOM Type ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "UOMTypeID",
					Message: "UOM Type ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "ConversionFactor",
					Message: "Conversion Factor is required on param",
				},
				{
					Tag:     "numeric",
					Field:   "ConversionFactor",
					Message: "Conversion Factor must be number",
				},
				{
					Tag:     "required",
					Field:   "Symbol",
					Message: "Symbol is required on param",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
