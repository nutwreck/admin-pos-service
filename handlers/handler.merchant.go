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

type handlerMerchant struct {
	merchant entities.EntityMerchant
}

func NewHandlerMerchant(merchant entities.EntityMerchant) *handlerMerchant {
	return &handlerMerchant{merchant: merchant}
}

/**
* ======================================
* Handler Ping Status Merchant Teritory
*=======================================
 */

func (h *handlerMerchant) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Merchant", http.StatusOK, nil)
}

/**
* =====================================
* Handler Create New Merchant Teritory
*======================================
 */
// CreateMasterMerchant godoc
// @Summary		Create Master Merchant
// @Description	Create Master Merchant
// @Tags		Master Merchant
// @Accept		mpfd
// @Produce		json
// @Param 		name formData string true "Name of the Merchant | input with lowercase"
// @Param 		phone formData string true "Phone of the Merchant | input numeric"
// @Param 		address formData string false "Address of the Merchant"
// @Param 		description formData string false "Description of the Merchant"
// @Param 		logo formData file false "File to be uploaded | Max Size File 1MB"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/merchant/create [post]
func (h *handlerMerchant) HandlerCreate(ctx *gin.Context) {
	var (
		body                   schemes.Merchant
		encryptedImageFileName string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileLogo, _ := ctx.FormFile("logo")
	body.Name = ctx.PostForm("name")
	body.Phone = ctx.PostForm("phone")
	body.Address = ctx.PostForm("address")
	body.Description = ctx.PostForm("description")
	if fileLogo != nil {
		//Validasi data
		validationMIME := helpers.ValidationMIMEFile(fileLogo.Filename, mimeTypeData)
		if !validationMIME {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Tipe file yang diupload bukan image",
					Value:   fileLogo.Filename,
					Param:   "Image",
					Tag:     "file type",
				},
			}
			err := schemes.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      errorsWithoutKeys,
			}
			ctx.AbortWithStatusJSON(err.StatusCode, err)
			return
		}

		if fileLogo.Size > configs.MaxFileSize1MB {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Ukuran file terlalu besar (maksimum 1MB)",
					Value:   fileLogo.Filename,
					Param:   "Image",
					Tag:     "file size",
				},
			}
			err := schemes.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      errorsWithoutKeys,
			}
			ctx.AbortWithStatusJSON(err.StatusCode, err)
			return
		}
	}
	if fileLogo != nil {
		//Body data
		encryptedImageFileName = helpers.EncryptFileName(fileLogo.Filename)
		body.Logo = encryptedImageFileName

		//Upload file
		uploadFile := helpers.UploadFileToStorageClient(fileLogo, encryptedImageFileName, configs.ACLPublicRead)
		if uploadFile != nil {
			fmt.Println("UPLOAD LOGO ERROR ==> " + uploadFile.Error())
			helpers.APIResponse(ctx, "Upload image failed", http.StatusInternalServerError, nil)
			return
		}
	}

	errors, code := ValidatorMerchant(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.merchant.EntityCreate(&body)

	if error.Type == "error_create_02" {
		// Delete file jika proses simpan gagal
		if fileLogo != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE LOGO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Merchant failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Merchant successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Merchant Teritory
*=======================================
 */
// GetListMasterMerchant godoc
// @Summary		Get List Master Merchant
// @Description	Get List Master Merchant
// @Tags		Master Merchant
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : merchant.id, merchant.name, merchant.active, merchant.created_at, default is merchant.name ASC | If you don't want to use it, fill it blank"
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
// @Router /api/v1/master/merchant/results [get]
func (h *handlerMerchant) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.Merchant
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

	res, totalData, error := h.merchant.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Merchant data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Merchant data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ======================================
* Handler Delete Merchant By ID Teritory
*=======================================
 */
// GetDeleteMasterMerchant godoc
// @Summary		Get Delete Master Merchant
// @Description	Get Delete Master Merchant
// @Tags		Master Merchant
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Master Merchant"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/merchant/delete [delete]
func (h *handlerMerchant) HandlerDelete(ctx *gin.Context) {
	var body schemes.Merchant
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorMerchant(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.merchant.EntityResult(&body)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Merchant data not found", error.Code, nil)
		return
	}

	res, error := h.merchant.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Merchant data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Merchant data for this id %v failed", id), error.Code, nil)
		return
	}

	//SAAT DELETE BERHASIL DELETE IMAGE FILE SEBELUMNYA DISTORAGE
	if getDataPrevious.Logo != constants.EMPTY_VALUE {
		deleteFile := helpers.DeleteFileFromStorageClient(getDataPrevious.Logo)
		if deleteFile != nil {
			fmt.Println("DELETE LOGO ERROR ==> " + deleteFile.Error())
			helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
			return
		}
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Merchant data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Merchant By ID Teritory
*=======================================
 */
// GetUpdateMasterMerchant godoc
// @Summary		Get Update Master Merchant
// @Description	Get Update Master Merchant
// @Tags		Master Merchant
// @Accept		mpfd
// @Produce		json
// @Param		id query string true "Update Master Merchant"
// @Param 		name formData string true "Name of the Merchant | input with lowercase"
// @Param 		phone formData string true "Phone of the Merchant | input numeric"
// @Param 		address formData string true "Address of the Merchant"
// @Param 		description formData string true "Description of the Merchant"
// @Param 		logo formData file false "File to be uploaded | Max Size File 1MB"
// @Param 		active formData bool true "Status Data"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/merchant/update [put]
func (h *handlerMerchant) HandlerUpdate(ctx *gin.Context) {
	var (
		body                   schemes.Merchant
		activeGet              = false
		encryptedImageFileName string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileLogo, _ := ctx.FormFile("logo")
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.Description = ctx.PostForm("description")
	body.Address = ctx.PostForm("address")
	body.Phone = ctx.PostForm("phone")
	activeStr := ctx.PostForm("active")
	if activeStr == "true" {
		activeGet = constants.TRUE_VALUE
	}
	body.Active = &activeGet
	if fileLogo != nil {
		//Validasi data
		validationMIME := helpers.ValidationMIMEFile(fileLogo.Filename, mimeTypeData)
		if !validationMIME {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Tipe file yang diupload bukan image",
					Value:   fileLogo.Filename,
					Param:   "Image",
					Tag:     "file type",
				},
			}
			err := schemes.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      errorsWithoutKeys,
			}
			ctx.AbortWithStatusJSON(err.StatusCode, err)
			return
		}

		if fileLogo.Size > configs.MaxFileSize1MB {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Ukuran file terlalu besar (maksimum 1MB)",
					Value:   fileLogo.Filename,
					Param:   "Image",
					Tag:     "file size",
				},
			}
			err := schemes.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      errorsWithoutKeys,
			}
			ctx.AbortWithStatusJSON(err.StatusCode, err)
			return
		}
	}
	if fileLogo != nil {
		//Body data
		encryptedImageFileName = helpers.EncryptFileName(fileLogo.Filename)
		body.Logo = encryptedImageFileName

		//Upload file
		uploadFile := helpers.UploadFileToStorageClient(fileLogo, encryptedImageFileName, configs.ACLPublicRead)
		if uploadFile != nil {
			fmt.Println("UPLOAD LOGO ERROR ==> " + uploadFile.Error())
			helpers.APIResponse(ctx, "Upload image failed", http.StatusInternalServerError, nil)
			return
		}
	}

	errors, code := ValidatorMerchant(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.merchant.EntityResult(&body)
	if error.Type == "error_result_01" {
		if fileLogo != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE LOGO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}

		helpers.APIResponse(ctx, "Merchant data not found", error.Code, nil)
		return
	}

	//Update data
	_, error = h.merchant.EntityUpdate(&body)

	if error.Type == "error_update_01" || error.Type == "error_update_02" {
		// Delete file jika proses simpan gagal
		if fileLogo != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE LOGO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}
	}

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Merchant data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Merchant data failed for this id %s", id), error.Code, nil)
		return
	}

	//SAAT UPDATE BERHASIL DELETE IMAGE FILE SEBELUMNYA DISTORAGE
	if fileLogo != nil && getDataPrevious.Logo != constants.EMPTY_VALUE {
		deleteFile := helpers.DeleteFileFromStorageClient(getDataPrevious.Logo)
		if deleteFile != nil {
			fmt.Println("DELETE LOGO ERROR ==> " + deleteFile.Error())
			helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
			return
		}
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Merchant data success for this id %s", id), http.StatusOK, nil)
}

/**
* ======================================
*  All Validator User Input For Merchant
*=======================================
 */

func ValidatorMerchant(ctx *gin.Context, input schemes.Merchant, Type string) (interface{}, int) {
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
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
