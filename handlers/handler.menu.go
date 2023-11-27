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

type handleMenu struct {
	menu entities.EntityMenu
}

func NewHandlerMenu(menu entities.EntityMenu) *handleMenu {
	return &handleMenu{menu: menu}
}

/**
* =============================================
* Handler Ping Status Master Menu Teritory
*==============================================
 */

func (h *handleMenu) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master Menu", http.StatusOK, nil)
}

/**
* ============================================
* Handler Create New Master Menu Teritory
*=============================================
 */
// CreateMasterMenu godoc
// @Summary		Create Master Menu
// @Description	Create Master Menu
// @Tags		Master Menu
// @Accept		json
// @Produce		json
// @Param		menu body schemes.MenuRequest true "Create Master Menu"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu/create [post]
func (h *handleMenu) HandlerCreate(ctx *gin.Context) {
	var body schemes.Menu
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMenu(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menu.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master Menu name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master Menu failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master Menu successfully", http.StatusCreated, nil)
}

/**
* =============================================
* Handler Results All Master Menu Teritory
*==============================================
 */
// GetListMasterMenu godoc
// @Summary		Get List Master Menu
// @Description	Get List Master Menu
// @Tags		Master Menu
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : menu.id, menu.name, menu.active, merchant.id, merchant.name, menu.created_at, default is menu.created_at DESC | If you don't want to use it, fill it blank"
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
// @Router /api/v1/master/menu/results [get]
func (h *handleMenu) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.Menu
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

	res, totalData, error := h.menu.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master Menu data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master Menu data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* =================================================
* Handler Results All Master Menu Relation Teritory
*==================================================
 */
// GetListMasterMenuRelation godoc
// @Summary		Get List Master Menu Relation
// @Description	Get List Master Menu Relation
// @Tags		Master Menu
// @Accept		json
// @Produce		json
// @Param sort query string false "Available Sorting Menu Name | Use ASC or DESC, default is ASC | If you don't want to use it, fill it blank"
// @Param merchant_id query string true "Search by merchant"
// @Param name query string false "Search by menu name using LIKE pattern"
// @Param id query string false "Search by menu ID"
// @Success 200 {object} schemes.ResponsesPagination
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu/results-relation [get]
func (h *handleMenu) HandlerResultRelations(ctx *gin.Context) {
	var (
		bodyMenu               schemes.Menu
		bodyMenuDetail         schemes.MenuDetail
		bodyMenuDetailFunction schemes.MenuDetailFunction
		result                 []schemes.GetMenuRelation
	)

	sortParam := ctx.DefaultQuery("sort", constants.EMPTY_VALUE)
	if sortParam != constants.EMPTY_VALUE {
		bodyMenu.Sort = sortParam
	}
	merchantParam := ctx.DefaultQuery("merchant_id", constants.EMPTY_VALUE)
	if merchantParam != constants.EMPTY_VALUE {
		bodyMenu.MerchantID = merchantParam
	} else {
		helpers.APIResponse(ctx, "Merchant ID is required on param", http.StatusBadRequest, nil)
		return
	}
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		bodyMenu.Name = nameParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		bodyMenu.ID = idParam
	}

	//res, error := h.menu.EntityResultRelations(&body)
	res, error := h.menu.EntityGetMenu(&bodyMenu)

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Master Menu data not found", error.Code, nil)
		return
	}

	if len(*res) > 0 {
		for _, item := range *res {
			var listDetail []schemes.ListMenuDetail
			var ListDetailFunction []schemes.ListMenuDetailFunction

			//Get Menu Detail
			if sortParam != constants.EMPTY_VALUE {
				bodyMenuDetail.Sort = sortParam
			}
			if item.MerchantID != constants.EMPTY_VALUE {
				bodyMenuDetail.MerchantID = item.MerchantID
			}
			if item.ID != constants.EMPTY_VALUE {
				bodyMenuDetail.MenuID = item.ID
			}
			resDetail, errorDetail := h.menu.EntityGetMenuDetail(&bodyMenuDetail)
			if errorDetail.Type == "error_results_01" {
				listDetail = nil
			} else {
				if len(*resDetail) > 0 {
					for _, itemDetail := range *resDetail {
						// Get Menu Detail Function
						if sortParam != constants.EMPTY_VALUE {
							bodyMenuDetailFunction.Sort = sortParam
						}
						if itemDetail.MerchantID != constants.EMPTY_VALUE {
							bodyMenuDetailFunction.MerchantID = itemDetail.MerchantID
						}
						if itemDetail.MenuID != constants.EMPTY_VALUE {
							bodyMenuDetailFunction.MenuID = itemDetail.MenuID
						}
						if itemDetail.ID != constants.EMPTY_VALUE {
							bodyMenuDetailFunction.MenuDetailID = itemDetail.ID
						}
						resFunction, errorFunction := h.menu.EntityGetMenuDetailFunction(&bodyMenuDetailFunction)
						if errorFunction.Type == "error_results_01" {
							ListDetailFunction = nil
						} else {
							if len(*resFunction) > 0 {
								for _, itemFunction := range *resFunction {
									addlistFunction := schemes.ListMenuDetailFunction{
										ID:     itemFunction.ID,
										Name:   itemFunction.Name,
										Link:   itemFunction.Link,
										Active: itemFunction.Active,
									}
									ListDetailFunction = append(ListDetailFunction, addlistFunction)
								}
							}
						}

						addlistDetail := schemes.ListMenuDetail{
							ID:                 itemDetail.ID,
							Name:               itemDetail.Name,
							Link:               itemDetail.Link,
							Image:              itemDetail.Image,
							Icon:               itemDetail.Icon,
							Active:             itemDetail.Active,
							ListDetailFunction: ListDetailFunction,
						}
						listDetail = append(listDetail, addlistDetail)
					}
				}
			}

			addResult := schemes.GetMenuRelation{
				ID:         item.ID,
				Name:       item.Name,
				Active:     item.Active,
				ListDetail: listDetail,
			}
			result = append(result, addResult)
		}
	}

	helpers.APIResponse(ctx, "Master Menu data already to use", http.StatusOK, result)
}

/**
* ==============================================
* Handler Delete Master Menu By ID Teritory
*===============================================
 */
// GetDeleteMasterMenu godoc
// @Summary		Get Delete Master Menu
// @Description	Get Delete Master Menu
// @Tags		Master Menu
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Master Menu"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu/delete [delete]
func (h *handleMenu) HandlerDelete(ctx *gin.Context) {
	var body schemes.Menu
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorMenu(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.menu.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Menu data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Menu data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Menu data for this id %s success", id), http.StatusOK, res)
}

/**
* ==============================================
* Handler Update Master Menu By ID Teritory
*===============================================
 */
// GetUpdateMasterMenu godoc
// @Summary		Get Update Master Menu
// @Description	Get Update Master Menu
// @Tags		Master Menu
// @Accept		json
// @Produce		json
// @Param		id query string true "Update Master Menu"
// @Param		menu body schemes.MenuRequest true "Update Master Menu"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu/update [put]
func (h *handleMenu) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.Menu
		activeGet = false
	)
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
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

	errors, code := ValidatorMenu(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menu.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Menu data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master Menu data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master Menu data success for this id %s", id), http.StatusOK, nil)
}

/**
* ==============================================
*  All Validator User Input For Master Menu
*===============================================
 */

func ValidatorMenu(ctx *gin.Context, input schemes.Menu, Type string) (interface{}, int) {
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
