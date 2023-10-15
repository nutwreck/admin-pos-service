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

type handlerUserOutlet struct {
	userOutlet entities.EntityUserOutlet
}

func NewHandlerUserOutlet(userOutlet entities.EntityUserOutlet) *handlerUserOutlet {
	return &handlerUserOutlet{userOutlet: userOutlet}
}

/**
* ===============================================
* Handler Ping Status Master User Outlet Teritory
*================================================
 */

func (h *handlerUserOutlet) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master User Outlet", http.StatusOK, nil)
}

/**
* ==============================================
* Handler Create New Master User Outlet Teritory
*===============================================
 */
// CreateMasterUserOutlet godoc
// @Summary		Create User Outlet
// @Description	Create User Outlet
// @Tags		User Outlet
// @Accept		json
// @Produce		json
// @Param		useroutlet body schemes.UserOutletRequest true "Create User Outlet"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/user-outlet/create [post]
func (h *handlerUserOutlet) HandlerCreate(ctx *gin.Context) {
	var body schemes.UserOutlet
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUserOutlet(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.userOutlet.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "User Outlet name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new User Outlet failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new User Outlet successfully", http.StatusCreated, nil)
}

/**
* ===============================================
* Handler Results All Master User Outlet Teritory
*================================================
 */
// GetListMasterUserOutlet godoc
// @Summary		Get List User Outlet
// @Description	Get List User Outlet
// @Tags		User Outlet
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : useroutlet.id, user.id, user.name, outlet.id, outlet.name, useroutlet.active, merchant.id, merchant.name, useroutlet.created_at, default is useroutlet.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param outlet_id query string false "Search by outlet"
// @Param user_id query string false "Search by user"
// @Param id query string false "Search by ID"
// @Success 200 {object} schemes.ResponsesPagination
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/user-outlet/results [get]
func (h *handlerUserOutlet) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.UserOutlet
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
	outletParam := ctx.DefaultQuery("outlet_id", constants.EMPTY_VALUE)
	if outletParam != constants.EMPTY_VALUE {
		body.OutletID = outletParam
	}
	userParam := ctx.DefaultQuery("user_id", constants.EMPTY_VALUE)
	if userParam != constants.EMPTY_VALUE {
		body.UserID = userParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		body.ID = idParam
	}

	res, totalData, error := h.userOutlet.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "User Outlet data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "User Outlet data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ================================================
* Handler Delete Master User Outlet By ID Teritory
*=================================================
 */
// GetDeleteMasterUserOutlet godoc
// @Summary		Get Delete User Outlet
// @Description	Get Delete User Outlet
// @Tags		User Outlet
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete User Outlet"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/user-outlet/delete/{id} [delete]
func (h *handlerUserOutlet) HandlerDelete(ctx *gin.Context) {
	var body schemes.UserOutlet
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorUserOutlet(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.userOutlet.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("User Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete User Outlet data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete User Outlet data for this id %s success", id), http.StatusOK, res)
}

/**
* ================================================
* Handler Update Master User Outlet By ID Teritory
*=================================================
 */
// GetUpdateMasterUserOutlet godoc
// @Summary		Get Update User Outlet
// @Description	Get Update User Outlet
// @Tags		User Outlet
// @Accept		json
// @Produce		json
// @Param		id path string true "Update User Outlet"
// @Param		useroutlet body schemes.UserOutletRequest true "Update User Outlet"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/user-outlet/update/{id} [put]
func (h *handlerUserOutlet) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.UserOutlet
		activeGet = false
	)
	id := ctx.Param("id")
	body.ID = id
	body.UserID = ctx.PostForm("user_id")
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

	errors, code := ValidatorUserOutlet(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.userOutlet.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("User Outlet data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update User Outlet data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update User Outlet data success for this id %s", id), http.StatusOK, nil)
}

/**
* ================================================
*  All Validator User Input For Master User Outlet
*=================================================
 */

func ValidatorUserOutlet(ctx *gin.Context, input schemes.UserOutlet, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
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
					Field:   "OutletID",
					Message: "Outlet ID is required on param",
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
					Field:   "OutletID",
					Message: "Outlet ID is required on param",
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
