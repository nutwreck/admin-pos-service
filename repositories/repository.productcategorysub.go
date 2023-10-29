package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
	"gorm.io/gorm"
)

type repositoryProductCategorySub struct {
	db *gorm.DB
}

func NewRepositoryProductCategorySub(db *gorm.DB) *repositoryProductCategorySub {
	return &repositoryProductCategorySub{db: db}
}

/**
* ==========================================================
* Repository Create New Master Product Category Sub Teritory
*===========================================================
 */

func (r *repositoryProductCategorySub) EntityCreate(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError) {
	var productCategorySub models.ProductCategorySub
	productCategorySub.MerchantID = input.MerchantID
	productCategorySub.OutletID = input.OutletID
	productCategorySub.ProductCategoryID = input.ProductCategoryID
	productCategorySub.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productCategorySub)

	checkData := db.Debug().Where("merchant_id = ? AND outlet_id = ? AND product_category_id = ? AND name = ?", productCategorySub.MerchantID, productCategorySub.OutletID, productCategorySub.ProductCategoryID, productCategorySub.Name).First(&productCategorySub)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &productCategorySub, <-err
	}

	addData := db.Debug().Create(&productCategorySub).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &productCategorySub, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &productCategorySub, <-err
}

/**
* ============================================================
* Repository Results All Master Product Category Sub Teritory
*=============================================================
 */

func (r *repositoryProductCategorySub) EntityResults(input *schemes.ProductCategorySub) (*[]schemes.GetProductCategorySub, int64, schemes.SchemeDatabaseError) {
	var (
		productcategorysub []models.ProductCategorySub
		result             []schemes.GetProductCategorySub
		countData          schemes.CountData
		args               []interface{}
		totalData          int64
		sortData           string = "productcategorysub.created_at DESC"
		queryCountData     string = constants.EMPTY_VALUE
		queryData          string = constants.EMPTY_VALUE
		queryAdditional    string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productcategorysub)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(productcategorysub.*) AS count_data
		FROM master.product_category_subs AS productcategorysub
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			productcategorysub.id,
			productcategorysub.name,
			productcategorysub.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			outlet.id AS outlet_id,
			outlet.name AS outlet_name,
			productcategory.id AS product_category_id,
			productcategory.name AS product_category_name,
			productcategorysub.created_at
		FROM master.product_category_subs AS productcategorysub
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON productcategorysub.merchant_id = merchant.id AND merchant.active = true
		JOIN master.outlets AS outlet ON productcategorysub.outlet_id = outlet.id AND outlet.active = true
		JOIN master.product_categorys AS productcategory ON productcategorysub.product_category_id = productcategory.id AND productcategory.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategorysub.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.OutletID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategorysub.outlet_id = ?`
		args = append(args, input.OutletID)
	}

	if input.ProductCategoryID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategory.id = ?`
		args = append(args, input.ProductCategoryID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategorysub.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategorysub.id = ?`
		args = append(args, input.ID)
	}

	//Eksekusi query ambil jumlah data tanpa limit
	db.Raw(queryCountData+queryAdditional, args...).Scan(&countData)

	queryAdditional += ` ORDER BY ` + sortData

	if input.Page != constants.EMPTY_NUMBER || input.PerPage != constants.EMPTY_NUMBER {
		queryAdditional += ` LIMIT ?`
		args = append(args, int(input.PerPage))

		queryAdditional += ` OFFSET ?`
		args = append(args, offset)
	}

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&result)

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, totalData, <-err
	}

	// Menghitung total data yang diambil
	totalData = countData.CountData

	err <- schemes.SchemeDatabaseError{}
	return &result, totalData, <-err
}

/**
* ==============================================================
* Repository Delete Master Product Category Sub By ID Teritory
*===============================================================
 */

func (r *repositoryProductCategorySub) EntityDelete(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError) {
	var productcategorysub models.ProductCategorySub
	productcategorysub.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productcategorysub)

	checkData := db.Debug().First(&productcategorysub)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &productcategorysub, <-err
	}

	deleteData := db.Debug().Delete(&productcategorysub)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &productcategorysub, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &productcategorysub, <-err
}

/**
* =============================================================
* Repository Update Master Product CategoryBy ID Teritory
*==============================================================
 */

func (r *repositoryProductCategorySub) EntityUpdate(input *schemes.ProductCategorySub) (*models.ProductCategorySub, schemes.SchemeDatabaseError) {
	var productcategorysub models.ProductCategorySub
	productcategorysub.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productcategorysub)

	checkData := db.Debug().First(&productcategorysub)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &productcategorysub, <-err
	}

	productcategorysub.MerchantID = input.MerchantID
	productcategorysub.OutletID = input.OutletID
	productcategorysub.ProductCategoryID = input.ProductCategoryID
	productcategorysub.Name = input.Name
	productcategorysub.Active = input.Active

	updateData := db.Debug().Updates(&productcategorysub)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &productcategorysub, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &productcategorysub, <-err
}
