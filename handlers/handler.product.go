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

type handleProduct struct {
	product entities.EntityProduct
}

func NewHandlerProduct(product entities.EntityProduct) *handleProduct {
	return &handleProduct{product: product}
}

/**
* ===============================================
* Handler Ping Status Master Product Teritory
*================================================
 */

func (h *handleProduct) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Master Product", http.StatusOK, nil)
}

/**
* ==============================================
* Handler Create New Master Product Teritory
*===============================================
 */
// CreateMasterProduct godoc
// @Summary		Create Master Product
// @Description	Create Master Product
// @Tags		Master Product
// @Accept		json
// @Produce		json
// @Param		product body []schemes.ProductRequest true "Create Master Product"
// @Success 200 {object} schemes.Responses
// @Success 201 {object} schemes.Responses201Example
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/product/create [post]
func (h *handleProduct) HandlerCreate(ctx *gin.Context) {
	var (
		body         []schemes.Product
		datas        []schemes.Product
		mimeTypeData = configs.AllowedImageMimeTypes
	)

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	for _, input := range body {
		errors, code := ValidatorProduct(ctx, input, "create")
		if code > 0 {
			helpers.ErrorResponse(ctx, errors)
			return
		}
	}

	//Check File Upload
	for _, files := range body {
		fileImage, _, err := helpers.Base64ToFile(files.Image)
		if err != nil {
			errorsWithoutKeys := []schemes.ResultMsgErrorValidator{
				{
					Message: err.Error(),
					Value:   fileImage.Filename,
					Param:   "Image",
					Tag:     "file validation",
				},
			}
			err := schemes.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      errorsWithoutKeys,
			}
			ctx.AbortWithStatusJSON(err.StatusCode, err)
			return
		}
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
	}

	for _, req := range body {
		var (
			products               schemes.Product
			encryptedImageFileName string
		)

		fileImage, decodedFileImage, _ := helpers.Base64ToFile(req.Image)
		products.MerchantID = req.MerchantID
		products.MerchantID = req.MerchantID
		products.OutletID = req.OutletID
		products.ProductCategoryID = req.ProductCategoryID
		products.ProductCategorySubID = req.ProductCategorySubID
		products.Code = req.Code
		products.Name = req.Name
		products.Barcode = req.Barcode
		products.CapitalPrice = req.CapitalPrice
		products.SellingPrice = req.SellingPrice
		products.SupplierID = req.SupplierID
		products.UnitOfMeasurementID = req.UnitOfMeasurementID

		if fileImage != nil {
			//Body data
			encryptedImageFileName = helpers.EncryptFileName(fileImage.Filename)
			products.Image = encryptedImageFileName

			//Upload file
			uploadFile := helpers.UploadFileBase64ToStorageClient(decodedFileImage, encryptedImageFileName, configs.ACLPublicRead)
			if uploadFile != nil {
				fmt.Println("UPLOAD IMAGE ERROR ==> " + uploadFile.Error())
				helpers.APIResponse(ctx, "Upload image failed", http.StatusInternalServerError, nil)
				return
			}
		}

		datas = append(datas, products)
	}

	_, error := h.product.EntityCreate(&datas)

	if error.Type == "error_create_01" || error.Type == "error_create_02" {
		// Delete file jika proses simpan gagal
		for _, del := range datas {
			if del.Image != constants.EMPTY_VALUE {
				deleteFile := helpers.DeleteFileFromStorageClient(del.Image)
				if deleteFile != nil {
					fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
					helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
					return
				}
			}
		}
	}

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Master Product name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Master Product failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Master Product successfully", http.StatusCreated, nil)
}

/**
* ===============================================
* Handler Results All Master Product Teritory
*================================================
 */
// GetListMasterProduct godoc
// @Summary		Get List Master Product
// @Description	Get List Master Product
// @Tags		Master Product
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : product.id, product.merchant_id, merchant.name, product.outlet_id, outlet.name AS outlet_name, product.product_category_id, product_category.name, product.product_category_sub_id, product_category_sub.name, product.code, product.name, product.barcode, product.capital_price, product.selling_price, product.supplier_id, supplier.name, product.unit_of_measurement_id, unit_of_measurement.name, product.active, default is product.created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param merchant_id query string false "Search by merchant"
// @Param outlet_id query string false "Search by outlet"
// @Param code query string false "Search by code"
// @Param product_category_id query string false "Search by Product Category"
// @Param product_category_sub_id query string false "Search by Product Category Sub"
// @Param unit_of_measurement_id query string false "Search by Unit Of Measurement"
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
// @Router /api/v1/master/product/results [get]
func (h *handleProduct) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.Product
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

	res, totalData, error := h.product.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Master Product data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Master Product data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ================================================
* Handler Delete Master Product By ID Teritory
*=================================================
 */
// GetDeleteMasterProduct godoc
// @Summary		Get Delete Master Product
// @Description	Get Delete Master Product
// @Tags		Master Product
// @Accept		json
// @Produce		json
// @Param		id query string true "Delete Master Product"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/product/delete [delete]
func (h *handleProduct) HandlerDelete(ctx *gin.Context) {
	var body schemes.Product
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id

	errors, code := ValidatorProduct(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.product.EntityResult(&body)
	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, "Master Product data not found", error.Code, nil)
		return
	}

	res, error := h.product.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Product data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Product data for this id %v failed", id), error.Code, nil)
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

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Master Product data for this id %s success", id), http.StatusOK, res)
}

/**
* ================================================
* Handler Update Master Product By ID Teritory
*=================================================
 */
// GetUpdateMasterProduct godoc
// @Summary		Get Update Master Product
// @Description	Get Update Master Product
// @Tags		Master Product
// @Accept		mpfd
// @Produce		json
// @Param		id query string true "Update Master Product"
// @Param 		merchant_id formData string true "Merchant ID (UUID)"
// @Param 		outlet_id formData string true "Outlet ID (UUID)"
// @Param 		product_category_id formData string true "Product Category ID (UUID)"
// @Param 		product_category_sub_id formData string true "Product Category Sub ID (UUID)"
// @Param 		code formData string true "Code of the Product"
// @Param 		name formData string true "Name of the Product"
// @Param 		barcode formData string false "Barcode of the Product"
// @Param 		capital_price formData string true "Capital Price of the Product"
// @Param 		selling_price formData string true "Selling Price of the Product"
// @Param 		supplier_id formData string false "Supplier ID (UUID)"
// @Param 		unit_of_measurement_id formData string true "UOM ID (UUID)"
// @Param 		image formData file false "File to be uploaded | Max Size File 1MB"
// @Param 		active formData bool true "Status Data"
// @Success 200 {object} schemes.Responses
// @Failure 400 {object} schemes.Responses400Example
// @Failure 401 {object} schemes.Responses401Example
// @Failure 403 {object} schemes.Responses403Example
// @Failure 404 {object} schemes.Responses404Example
// @Failure 409 {object} schemes.Responses409Example
// @Failure 500 {object} schemes.Responses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/master/product/update [put]
func (h *handleProduct) HandlerUpdate(ctx *gin.Context) {
	var (
		body                   schemes.Product
		activeGet              = false
		encryptedImageFileName string
		mimeTypeData           = configs.AllowedImageMimeTypes
	)

	fileImage, _ := ctx.FormFile("image")
	id := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	body.ID = id
	body.Name = ctx.PostForm("name")
	body.MerchantID = ctx.PostForm("merchant_id")
	body.OutletID = ctx.PostForm("outlet_id")
	body.ProductCategoryID = ctx.PostForm("product_category_id")
	body.ProductCategorySubID = ctx.PostForm("product_category_sub_id")
	body.Code = ctx.PostForm("code")
	body.Name = ctx.PostForm("name")
	body.Barcode = ctx.PostForm("barcode")
	capitalPriceStr := ctx.PostForm("capital_price")
	capitalPrice, _ := strconv.ParseFloat(capitalPriceStr, 64)
	body.CapitalPrice = capitalPrice
	sellingPriceStr := ctx.PostForm("selling_price")
	sellingPrice, _ := strconv.ParseFloat(sellingPriceStr, 64)
	body.SellingPrice = sellingPrice
	body.SupplierID = ctx.PostForm("supplier_id")
	body.UnitOfMeasurementID = ctx.PostForm("unit_of_measurement_id")
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

	errors, code := ValidatorProduct(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	//Cek data sebelumnya
	getDataPrevious, error := h.product.EntityResult(&body)
	if error.Type == "error_result_01" {
		if fileImage != nil {
			deleteFile := helpers.DeleteFileFromStorageClient(encryptedImageFileName)
			if deleteFile != nil {
				fmt.Println("DELETE IMAGE ERROR ==> " + deleteFile.Error())
				helpers.APIResponse(ctx, "Delete image failed", http.StatusInternalServerError, nil)
				return
			}
		}

		helpers.APIResponse(ctx, "Master Product data not found", error.Code, nil)
		return
	}

	//Update data
	_, error = h.product.EntityUpdate(&body)

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
	}

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Master Product data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Master Product data failed for this id %s", id), error.Code, nil)
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

	helpers.APIResponse(ctx, fmt.Sprintf("Update Master Product data success for this id %s", id), http.StatusOK, nil)
}

/**
* ================================================
*  All Validator User Input For Master Product
*=================================================
 */

func ValidatorProduct(ctx *gin.Context, input schemes.Product, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig
	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
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
					Field:   "OutletID",
					Message: "Outlet ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "OutletID",
					Message: "Outlet ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "ProductCategoryID",
					Message: "Product Category ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "ProductCategoryID",
					Message: "Product Category ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "ProductCategorySubID",
					Message: "Product Category Sub ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "ProductCategorySubID",
					Message: "Product Category Sub ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "Code",
					Message: "Code is required on body",
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
					Field:   "CapitalPrice",
					Message: "Capital Price is required on body",
				},
				{
					Tag:     "numeric",
					Field:   "CapitalPrice",
					Message: "Capital Price must be number",
				},
				{
					Tag:     "required",
					Field:   "SellingPrice",
					Message: "Selling Price is required on body",
				},
				{
					Tag:     "numeric",
					Field:   "SellingPrice",
					Message: "Selling Price must be number",
				},
				{
					Tag:     "required",
					Field:   "UnitOfMeasurementID",
					Message: "Unit Of Measurement ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "UnitOfMeasurementID",
					Message: "Unit Of Measurement ID must be uuid",
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
					Field:   "OutletID",
					Message: "Outlet ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "OutletID",
					Message: "Outlet ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "ProductCategoryID",
					Message: "Product Category ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "ProductCategoryID",
					Message: "Product Category ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "ProductCategorySubID",
					Message: "Product Category Sub ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "ProductCategorySubID",
					Message: "Product Category Sub ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "Code",
					Message: "Code is required on body",
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
					Field:   "CapitalPrice",
					Message: "Capital Price is required on body",
				},
				{
					Tag:     "numeric",
					Field:   "CapitalPrice",
					Message: "Capital Price must be number",
				},
				{
					Tag:     "required",
					Field:   "SellingPrice",
					Message: "Selling Price is required on body",
				},
				{
					Tag:     "numeric",
					Field:   "SellingPrice",
					Message: "Selling Price must be number",
				},
				{
					Tag:     "required",
					Field:   "UnitOfMeasurementID",
					Message: "Unit Of Measurement ID is required on body",
				},
				{
					Tag:     "uuid",
					Field:   "UnitOfMeasurementID",
					Message: "Unit Of Measurement ID must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
