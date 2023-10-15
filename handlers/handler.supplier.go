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

type handleSupplier struct {
	supplier entities.EntitySupplier
}

func NewHandlerSupplier(supplier entities.EntitySupplier) *handleSupplier {
	return &handleSupplier{supplier: supplier}
}

/**
* ======================================
* Handler Ping Status Supplier Teritory
*=======================================
 */

func (h *handleSupplier) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Supplier", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Supplier Teritory
*======================================
 */
// CreateMasterSupplier godoc
// @Summary		Create Master Supplier
// @Description	Create Master Supplier
// @Tags		Master Supplier
// @Accept		json
// @Produce		json
// @Param		supplier body schemes.SupplierRequest true "Create Master Supplier"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/supplier/create [post]
func (h *handleSupplier) HandlerCreate(ctx *gin.Context) {
	var body schemes.Supplier
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSupplier(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.supplier.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Supplier phone number already taken", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Supplier failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Supplier successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Supplier Teritory
*=======================================
 */
// GetListMasterSupplier godoc
// @Summary		Get List Master Supplier
// @Description	Get List Master Supplier
// @Tags		Master Supplier
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : supplier.id, supplier.name, supplier.phone, supplier.active, merchant.id, merchant.name, outlet.id, outlet.name, supplier.created_at, default is supplier.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param outlet_id query string false "Search by outlet"
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
// @Router /api/v1/master/supplier/results [get]
func (h *handleSupplier) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.Supplier
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
	outletParam := ctx.DefaultQuery("outlet_id", constants.EMPTY_VALUE)
	if outletParam != constants.EMPTY_VALUE {
		body.OutletID = outletParam
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

	res, totalData, error := h.supplier.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Supplier data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Supplier data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ======================================
* Handler Result Supplier By ID Teritory
*=======================================
 */

func (h *handleSupplier) HandlerResult(ctx *gin.Context) {
	var body schemes.Supplier
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorSupplier(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.supplier.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Supplier data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Supplier data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Supplier By ID Teritory
*=======================================
 */
// GetDeleteMasterSupplier godoc
// @Summary		Get Delete Master Supplier
// @Description	Get Delete Master Supplier
// @Tags		Master Supplier
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Master Supplier"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/supplier/delete/{id} [delete]
func (h *handleSupplier) HandlerDelete(ctx *gin.Context) {
	var body schemes.Supplier
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorSupplier(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.supplier.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Supplier data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Supplier data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Supplier data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Supplier By ID Teritory
*=======================================
 */
// GetUpdateMasterSupplier godoc
// @Summary		Get Update Master Supplier
// @Description	Get Update Master Supplier
// @Tags		Master Supplier
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Master Supplier"
// @Param		supplier body schemes.SupplierRequest true "Update Master Supplier"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/supplier/update/{id} [put]
func (h *handleSupplier) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.Supplier
		activeGet = false
	)
	id := ctx.Param("id")
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.Description = ctx.PostForm("description")
	body.Address = ctx.PostForm("address")
	body.Phone = ctx.PostForm("phone")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.OutletID = ctx.PostForm("outlet_id")
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

	errors, code := ValidatorSupplier(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.supplier.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Supplier data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Supplier data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Supplier data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Supplier
*=======================================
 */

func ValidatorSupplier(ctx *gin.Context, input schemes.Supplier, Type string) (interface{}, int) {
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
					Tag:     "gte",
					Field:   "Phone",
					Message: "Phone number must be 12 character",
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
					Tag:     "max",
					Field:   "Description",
					Message: "Description maximal 1000 character",
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
				{
					Tag:     "required",
					Field:   "OutletID",
					Message: "Outlet ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "OutletID",
					Message: "Outlet ID must be uuid",
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
					Tag:     "gte",
					Field:   "Phone",
					Message: "Phone number must be 12 character",
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
					Tag:     "max",
					Field:   "Description",
					Message: "Description maximal 1000 character",
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
				{
					Tag:     "required",
					Field:   "OutletID",
					Message: "Outlet ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "OutletID",
					Message: "Outlet ID must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
