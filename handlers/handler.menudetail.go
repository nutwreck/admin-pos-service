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

type handleMenuDetail struct {
	menuDetail entities.EntityMenuDetail
}

func NewHandlerMenuDetail(menuDetail entities.EntityMenuDetail) *handleMenuDetail {
	return &handleMenuDetail{menuDetail: menuDetail}
}

/**
* ===============================================
* Handler Ping Status Master Menu Detail Teritory
*================================================
 */

func (h *handleMenuDetail) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master Menu Detail", http.StatusOK, nil)
}

/**
* ==============================================
* Handler Create New Master Menu Detail Teritory
*===============================================
 */
// CreateMasterMenuDetail godoc
// @Summary		Create Master Menu Detail
// @Description	Create Master Menu Detail
// @Tags		Master Menu Detail
// @Accept		mpfd
// @Produce		json
// @Param 		merchant_id formData string true "Merchant ID (UUID)"
// @Param 		menu_id formData string true "Menu ID (UUID)"
// @Param 		name formData string true "Name of the Menu Detail"
// @Param 		link formData string true "Link of the Menu Detail"
// @Param 		image formData file false "File to be uploaded | Max Size File 1MB"
// @Param 		icon formData file false "File to be uploaded | Max Size File 1MB"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu-detail/create [post]
func (h *handleMenuDetail) HandlerCreate(ctx *gin.Context) {
	var (
		body                   schemes.MenuDetail
		encryptedImageFileName string
		encryptedIconFileName  string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileImage, _ := ctx.FormFile("image")
	fileIcon, _ := ctx.FormFile("icon")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.Name = ctx.PostForm("name")
	body.MenuID = ctx.PostForm("menu_id")
	body.Link = ctx.PostForm("link")
	if fileImage != nil {
		//Validasi data
		validationMIME := helpers.ValidationMIMEFile(fileImage.Filename, mimeTypeData)
		if !validationMIME {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Tipe file yang diupload bukan image",
					Value:   fileImage.Filename,
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

		if fileImage.Size > configs.MaxFileSize1MB {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Ukuran file terlalu besar (maksimum 1MB)",
					Value:   fileImage.Filename,
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
	if fileIcon != nil {
		//Validasi data
		validationMIME := helpers.ValidationMIMEFile(fileIcon.Filename, mimeTypeData)
		if !validationMIME {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Tipe file yang diupload bukan image",
					Value:   fileIcon.Filename,
					Param:   "Icon",
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

		if fileIcon.Size > configs.MaxFileSize1MB {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Ukuran file terlalu besar (maksimum 1MB)",
					Value:   fileIcon.Filename,
					Param:   "Icon",
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
	if fileImage != nil {
		//Body data
		encryptedImageFileName = helpers.EncryptFileName(fileImage.Filename)
		body.Image = encryptedImageFileName

		//Upload file
		uploadFile := helpers.UploadFileToStorageClient(fileImage, encryptedImageFileName, configs.ACLPublicRead)
		if uploadFile != nil {
			fmt.Println("UPLOAD IMAGE ERROR ==> " + uploadFile.Error())
			helpers.APIResponse(ctx, "Upload image failed", http.StatusInternalServerError, nil)
			return
		}
	}
	if fileIcon != nil {
		//Body data
		encryptedIconFileName = helpers.EncryptFileName(fileIcon.Filename)
		body.Icon = encryptedIconFileName

		//Upload file
		uploadFile := helpers.UploadFileToStorageClient(fileIcon, encryptedIconFileName, configs.ACLPublicRead)
		if uploadFile != nil {
			fmt.Println("UPLOAD ICO ERROR ==> " + uploadFile.Error())
			helpers.APIResponse(ctx, "Upload icon failed", http.StatusInternalServerError, nil)
			return
		}
	}

	errors, code := ValidatorMenuDetail(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menuDetail.EntityCreate(&body)

	if error.Type == "error_create_01" || error.Type == "error_create_02" {
		// Delete file jika proses simpan gagal
		if fileImage != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}
		if fileIcon != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedIconFileName)
			if deleteFile != nil {
				fmt.Println("DELETE ICO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete icon failed", http.StatusInternalServerError, nil)
				return
			}
		}
	}

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master Menu Detail name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master Menu Detail failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master Menu Detail successfully", http.StatusCreated, nil)
}

/**
* ===============================================
* Handler Results All Master Menu Detail Teritory
*================================================
 */
// GetListMasterMenuDetail godoc
// @Summary		Get List Master Menu Detail
// @Description	Get List Master Menu Detail
// @Tags		Master Menu Detail
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : menudetail.id, menudetail.menu_id, menu.name, menudetail.name, menudetail.link, menudetail.active, merchant.id, merchant.name, default is menudetail.created_at DESC | If you don't want to use it, fill it blank"
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
// @Router /api/v1/master/menu-detail/results [get]
func (h *handleMenuDetail) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.MenuDetail
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

	res, totalData, error := h.menuDetail.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master Menu Detail data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master Menu Detail data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ================================================
* Handler Delete Master Menu Detail By ID Teritory
*=================================================
 */
// GetDeleteMasterMenuDetail godoc
// @Summary		Get Delete Master Menu Detail
// @Description	Get Delete Master Menu Detail
// @Tags		Master Menu Detail
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Master Menu Detail"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu-detail/delete/{id} [delete]
func (h *handleMenuDetail) HandlerDelete(ctx *gin.Context) {
	var body schemes.MenuDetail
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorMenuDetail(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.menuDetail.EntityResult(&body)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Master Menu Detail data not found", error.Code, nil)
		return
	}

	res, error := h.menuDetail.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Menu Detail data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Menu Detail data for this id %v failed", id), error.Code, nil)
		return
	}

	//SAAT DELETE BERHASIL DELETE IMAGE FILE SEBELUMNYA DISTORAGE
	if getDataPrevious.Image != constants.EMPTY_VALUE {
		deleteFile := helpers.DeleteFileFromStorageClient(getDataPrevious.Image)
		if deleteFile != nil {
			fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
			helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
			return
		}
	}
	if getDataPrevious.Icon != constants.EMPTY_VALUE {
		deleteFile := helpers.DeleteFileFromStorageClient(getDataPrevious.Icon)
		if deleteFile != nil {
			fmt.Println("DELETE ICO ERROR ==> " + deleteFile.Error())
			helpers.APIResponse(ctx, "Delete icon failed", http.StatusInternalServerError, nil)
			return
		}
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Menu Detail data for this id %s success", id), http.StatusOK, res)
}

/**
* ================================================
* Handler Update Master Menu Detail By ID Teritory
*=================================================
 */
// GetUpdateMasterMenuDetail godoc
// @Summary		Get Update Master Menu Detail
// @Description	Get Update Master Menu Detail
// @Tags		Master Menu Detail
// @Accept		mpfd
// @Produce		json
// @Param		id path string true "Update Master Menu Detail"
// @Param 		merchant_id formData string true "Merchant ID (UUID)"
// @Param 		menu_id formData string true "Menu ID (UUID)"
// @Param 		name formData string true "Name of the Menu Detail"
// @Param 		link formData string true "Link of the Menu Detail"
// @Param 		image formData file false "File to be uploaded | Max Size File 1MB"
// @Param 		icon formData file false "File to be uploaded | Max Size File 1MB"
// @Param 		active formData bool true "Status Data"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/menu-detail/update/{id} [put]
func (h *handleMenuDetail) HandlerUpdate(ctx *gin.Context) {
	var (
		body                   schemes.MenuDetail
		activeGet              = false
		encryptedImageFileName string
		encryptedIconFileName  string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileImage, _ := ctx.FormFile("image")
	fileIcon, _ := ctx.FormFile("icon")
	id := ctx.Param("id")
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.MenuID = ctx.PostForm("menu_id")
	body.Link = ctx.PostForm("link")
	activeStr := ctx.PostForm("active")
	if activeStr == "true" {
		activeGet = constants.TRUE_VALUE
	}
	body.Active = &activeGet
	if fileImage != nil {
		//Validasi data
		validationMIME := helpers.ValidationMIMEFile(fileImage.Filename, mimeTypeData)
		if !validationMIME {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Tipe file yang diupload bukan image",
					Value:   fileImage.Filename,
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

		if fileImage.Size > configs.MaxFileSize1MB {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Ukuran file terlalu besar (maksimum 1MB)",
					Value:   fileImage.Filename,
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
	if fileIcon != nil {
		//Validasi data
		validationMIME := helpers.ValidationMIMEFile(fileIcon.Filename, mimeTypeData)
		if !validationMIME {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Tipe file yang diupload bukan image",
					Value:   fileIcon.Filename,
					Param:   "Icon",
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

		if fileIcon.Size > configs.MaxFileSize1MB {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: "Ukuran file terlalu besar (maksimum 1MB)",
					Value:   fileIcon.Filename,
					Param:   "Icon",
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
	if fileImage != nil {
		//Body data
		encryptedImageFileName = helpers.EncryptFileName(fileImage.Filename)
		body.Image = encryptedImageFileName

		//Upload file
		uploadFile := helpers.UploadFileToStorageClient(fileImage, encryptedImageFileName, configs.ACLPublicRead)
		if uploadFile != nil {
			fmt.Println("UPLOAD IMAGE ERROR ==> " + uploadFile.Error())
			helpers.APIResponse(ctx, "Upload image failed", http.StatusInternalServerError, nil)
			return
		}
	}
	if fileIcon != nil {
		//Body data
		encryptedIconFileName = helpers.EncryptFileName(fileIcon.Filename)
		body.Icon = encryptedIconFileName

		//Upload file
		uploadFile := helpers.UploadFileToStorageClient(fileIcon, encryptedIconFileName, configs.ACLPublicRead)
		if uploadFile != nil {
			fmt.Println("UPLOAD ICO ERROR ==> " + uploadFile.Error())
			helpers.APIResponse(ctx, "Upload icon failed", http.StatusInternalServerError, nil)
			return
		}
	}

	errors, code := ValidatorMenuDetail(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.menuDetail.EntityResult(&body)
	if error.Type == "error_result_01" {
		if fileImage != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}
		if fileIcon != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedIconFileName)
			if deleteFile != nil {
				fmt.Println("DELETE ICO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete icon failed", http.StatusInternalServerError, nil)
				return
			}
		}

		helpers.APIResponse(ctx, "Master Menu Detail data not found", error.Code, nil)
		return
	}

	//Update data
	_, error = h.menuDetail.EntityUpdate(&body)

	if error.Type == "error_update_01" || error.Type == "error_update_02" {
		// Delete file jika proses simpan gagal
		if fileImage != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}
		if fileIcon != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedIconFileName)
			if deleteFile != nil {
				fmt.Println("DELETE ICO ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete icon failed", http.StatusInternalServerError, nil)
				return
			}
		}
	}

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Menu Detail data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master Menu Detail data failed for this id %s", id), error.Code, nil)
		return
	}

	//SAAT UPDATE BERHASIL DELETE IMAGE FILE SEBELUMNYA DISTORAGE
	if fileImage != nil && getDataPrevious.Image != constants.EMPTY_VALUE {
		deleteFile := helpers.DeleteFileFromStorageClient(getDataPrevious.Image)
		if deleteFile != nil {
			fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
			helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
			return
		}
	}
	if fileIcon != nil && getDataPrevious.Icon != constants.EMPTY_VALUE {
		deleteFile := helpers.DeleteFileFromStorageClient(getDataPrevious.Icon)
		if deleteFile != nil {
			fmt.Println("DELETE ICO ERROR ==> " + deleteFile.Error())
			helpers.APIResponse(ctx, "Delete icon failed", http.StatusInternalServerError, nil)
			return
		}
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master Menu Detail data success for this id %s", id), http.StatusOK, nil)
}

/**
* ================================================
*  All Validator User Input For Master Menu Detail
*=================================================
 */

func ValidatorMenuDetail(ctx *gin.Context, input schemes.MenuDetail, Type string) (interface{}, int) {
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
					Field:   "Link",
					Message: "Link is required on body",
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
					Field:   "Link",
					Message: "Link is required on body",
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
