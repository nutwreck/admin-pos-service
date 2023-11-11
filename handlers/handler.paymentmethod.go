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

type handlerPaymentMethod struct {
	paymentMethod entities.EntityPaymentMethod
}

func NewHandlerPaymentMethod(paymentMethod entities.EntityPaymentMethod) *handlerPaymentMethod {
	return &handlerPaymentMethod{paymentMethod: paymentMethod}
}

/**
* ===================================================
* Handler Ping Status Master Payment Method Teritory
*====================================================
 */

func (h *handlerPaymentMethod) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master Payment Method", http.StatusOK, nil)
}

/**
* ==================================================
* Handler Create New Master Payment Method Teritory
*===================================================
 */
// CreateMasterPaymentMethod godoc
// @Summary		Create Master Payment Method
// @Description	Create Master Payment Method
// @Tags		Master Payment Method
// @Accept		mpfd
// @Produce		json
// @Param 		merchant_id formData string true "Merchant ID (UUID)"
// @Param 		payment_category_id formData string true "Payment Category ID"
// @Param 		name formData string true "Name of the Payment Method | input with lowercase"
// @Param 		account_number formData string false "Account Number of the Payment Method | input numeric"
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
// @Router /api/v1/master/payment-method/create [post]
func (h *handlerPaymentMethod) HandlerCreate(ctx *gin.Context) {
	var (
		body                   schemes.PaymentMethod
		encryptedImageFileName string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileLogo, _ := ctx.FormFile("logo")
	body.Name = ctx.PostForm("name")
	body.AccountNumber = ctx.PostForm("account_number")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.PaymentCategoryID = ctx.PostForm("payment_category_id")
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

	errors, code := ValidatorPaymentMethod(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.paymentMethod.EntityCreate(&body)

	if error.Type == "error_create_02" || error.Type == "error_create_01" {
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

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master Payment Method name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master Payment Method failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master Payment Method successfully", http.StatusCreated, nil)
}

/**
* ====================================================
* Handler Results All Master Payment Method Teritory
*=====================================================
 */
// GetListMasterPaymentMethod godoc
// @Summary		Get List Master Payment Method
// @Description	Get List Master Payment Method
// @Tags		Master Payment Method
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : paymentmethod.id, paymentmethod.name, paymentcategoryid.id, paymentcategoryid.name, paymentmethod.account_number, paymentmethod.active, paymentmethod.created_at, default is paymentmethod.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param payment_category_id query string false "Search by Payment Category ID"
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
// @Router /api/v1/master/payment-method/results [get]
func (h *handlerPaymentMethod) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.PaymentMethod
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
	paymentCategoryParam := ctx.DefaultQuery("payment_category_id", constants.EMPTY_VALUE)
	if paymentCategoryParam != constants.EMPTY_VALUE {
		body.PaymentCategoryID = paymentCategoryParam
	}
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		body.ID = idParam
	}

	res, totalData, error := h.paymentMethod.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master Payment Method data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master Payment Method data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ===================================================
* Handler Delete Master Payment Method By ID Teritory
*====================================================
 */
// GetDeleteMasterPaymentMethod godoc
// @Summary		Get Delete Master Payment Method
// @Description	Get Delete Master Payment Method
// @Tags		Master Payment Method
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Master Payment Method"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/payment-method/delete [delete]
func (h *handlerPaymentMethod) HandlerDelete(ctx *gin.Context) {
	var body schemes.PaymentMethod
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorPaymentMethod(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.paymentMethod.EntityResult(&body)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Master Payment Method data not found", error.Code, nil)
		return
	}

	res, error := h.paymentMethod.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Payment Method data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Payment Method data for this id %v failed", id), error.Code, nil)
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

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Payment Method data for this id %s success", id), http.StatusOK, res)
}

/**
* =====================================================
* Handler Update Master Payment Method By ID Teritory
*======================================================
 */
// GetUpdateMasterPaymentMethod godoc
// @Summary		Get Update Master Payment Method
// @Description	Get Update Master Payment Method
// @Tags		Master Payment Method
// @Accept		mpfd
// @Produce		json
// @Param		id query string true "Update Master Payment Method"
// @Param 		merchant_id formData string true "Merchant ID (UUID)"
// @Param 		payment_category_id formData string true "Payment Method ID"
// @Param 		name formData string true "Name of the Payment Method | input with lowercase"
// @Param 		account_number formData string false "Account Number of the Payment Method | input numeric"
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
// @Router /api/v1/master/payment-method/update [put]
func (h *handlerPaymentMethod) HandlerUpdate(ctx *gin.Context) {
	var (
		body                   schemes.PaymentMethod
		activeGet              = false
		encryptedImageFileName string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileLogo, _ := ctx.FormFile("logo")
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.PaymentCategoryID = ctx.PostForm("payment_category_id")
	body.AccountNumber = ctx.PostForm("account_number")
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

	errors, code := ValidatorPaymentMethod(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.paymentMethod.EntityResult(&body)
	if error.Type == "error_result_01" {
		if fileLogo != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE LOGO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}

		helpers.APIResponse(ctx, "Master Payment Method data not found", error.Code, nil)
		return
	}

	//Update data
	_, error = h.paymentMethod.EntityUpdate(&body)

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
		helpers.APIResponse(ctx, fmt.Sprintf("Master Payment Method data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master Payment Method data failed for this id %s", id), error.Code, nil)
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

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master Payment Method data success for this id %s", id), http.StatusOK, nil)
}

/**
* ===================================================
*  All Validator User Input For Master Payment Method
*====================================================
 */

func ValidatorPaymentMethod(ctx *gin.Context, input schemes.PaymentMethod, Type string) (interface{}, int) {
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
					Tag:     "numeric",
					Field:   "AccountNumber",
					Message: "Account Number must be number",
				},
				{
					Tag:     "required",
					Field:   "PaymentCategoryID",
					Message: "Payment Category ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "PaymentCategoryID",
					Message: "Payment Category ID must be uuid",
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
					Tag:     "numeric",
					Field:   "AccountNumber",
					Message: "Account Number must be number",
				},
				{
					Tag:     "required",
					Field:   "PaymentCategoryID",
					Message: "Payment Category ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "PaymentCategoryID",
					Message: "Payment Category ID must be uuid",
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
