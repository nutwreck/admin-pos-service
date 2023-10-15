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

type handlePaymentCategory struct {
	paymentCategory entities.EntityPaymentCategory
}

func NewHandlerPaymentCategory(paymentCategory entities.EntityPaymentCategory) *handlePaymentCategory {
	return &handlePaymentCategory{paymentCategory: paymentCategory}
}

/**
* =====================================================
* Handler Ping Status Master Payment Category Teritory
*======================================================
 */

func (h *handlePaymentCategory) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master Payment Category", http.StatusOK, nil)
}

/**
* ====================================================
* Handler Create New Master Payment Category Teritory
*=====================================================
 */
// CreateMasterPaymentCategory godoc
// @Summary		Create Master Payment Category
// @Description	Create Master Payment Category
// @Tags		Master Payment Category
// @Accept		json
// @Produce		json
// @Param		paymentcategory body schemes.PaymentCategoryRequest true "Create Master Payment Category"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/payment-category/create [post]
func (h *handlePaymentCategory) HandlerCreate(ctx *gin.Context) {
	var body schemes.PaymentCategory
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorPaymentCategory(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.paymentCategory.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master Payment Category name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master Payment Category failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master Payment Category successfully", http.StatusCreated, nil)
}

/**
* ====================================================
* Handler Results All Master Payment Category Teritory
*=====================================================
 */
// GetListMasterPaymentCategory godoc
// @Summary		Get List Master Payment Category
// @Description	Get List Master Payment Category
// @Tags		Master Payment Category
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : paymentcategory.id, paymentcategory.name, paymentcategory.active, paymentcategory.created_at, default is paymentcategory.created_at DESC | If you don't want to use it, fill it blank"
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
// @Router /api/v1/master/payment-category/results [get]
func (h *handlePaymentCategory) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.PaymentCategory
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

	res, totalData, error := h.paymentCategory.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master Payment Category data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master Payment Category data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* =====================================================
* Handler Delete Master Payment Category By ID Teritory
*======================================================
 */
// GetDeleteMasterPaymentCategory godoc
// @Summary		Get Delete Master Payment Category
// @Description	Get Delete Master Payment Category
// @Tags		Master Payment Category
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Master Payment Category"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/payment-category/delete/{id} [delete]
func (h *handlePaymentCategory) HandlerDelete(ctx *gin.Context) {
	var body schemes.PaymentCategory
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorPaymentCategory(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.paymentCategory.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Payment Category data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Payment Category data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Payment Category data for this id %s success", id), http.StatusOK, res)
}

/**
* =====================================================
* Handler Update Master Payment Category By ID Teritory
*======================================================
 */
// GetUpdateMasterPaymentCategory godoc
// @Summary		Get Update Master Payment Category
// @Description	Get Update Master Payment Category
// @Tags		Master Payment Category
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Master Payment Category"
// @Param		paymentcategory body schemes.PaymentCategoryRequest true "Update Master Payment Category"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/payment-category/update/{id} [put]
func (h *handlePaymentCategory) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.PaymentCategory
		activeGet = false
	)
	id := ctx.Param("id")
	body.ID = id
	body.Name = ctx.PostForm("name")
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

	errors, code := ValidatorPaymentCategory(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.paymentCategory.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Payment Category data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master Payment Category data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master Payment Category data success for this id %s", id), http.StatusOK, nil)
}

/**
* ==============================================
*  All Validator User Input For Master Payment Category
*===============================================
 */

func ValidatorPaymentCategory(ctx *gin.Context, input schemes.PaymentCategory, Type string) (interface{}, int) {
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
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
