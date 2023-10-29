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

type repositoryProductCategory struct {
	db *gorm.DB
}

func NewRepositoryProductCategory(db *gorm.DB) *repositoryProductCategory {
	return &repositoryProductCategory{db: db}
}

/**
* ======================================================
* Repository Create New Master Product Category Teritory
*=======================================================
 */

func (r *repositoryProductCategory) EntityCreate(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError) {
	var productCategory models.ProductCategory
	productCategory.MerchantID = input.MerchantID
	productCategory.OutletID = input.OutletID
	productCategory.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productCategory)

	checkData := db.Debug().Where("merchant_id = ? AND outlet_id = ? AND name = ?", productCategory.MerchantID, productCategory.OutletID, productCategory.Name).First(&productCategory)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &productCategory, <-err
	}

	addData := db.Debug().Create(&productCategory).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &productCategory, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &productCategory, <-err
}

/**
* =======================================================
* Repository Results All Master Product Category Teritory
*========================================================
 */

func (r *repositoryProductCategory) EntityResults(input *schemes.ProductCategory) (*[]schemes.GetProductCategory, int64, schemes.SchemeDatabaseError) {
	var (
		productcategory []models.ProductCategory
		result          []schemes.GetProductCategory
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "productcategory.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productcategory)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(productcategory.*) AS count_data
		FROM master.product_categorys AS productcategory
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			productcategory.id,
			productcategory.name,
			productcategory.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			outlet.id AS outlet_id,
			outlet.name AS outlet_name,
			productcategory.created_at
		FROM master.product_categorys AS productcategory
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON productcategory.merchant_id = merchant.id AND merchant.active = true
		JOIN master.outlets AS outlet ON productcategory.outlet_id = outlet.id AND outlet.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategory.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.OutletID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategory.outlet_id = ?`
		args = append(args, input.OutletID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategory.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND productcategory.id = ?`
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
* Repository Delete Master Product Category By ID Teritory
*===============================================================
 */

func (r *repositoryProductCategory) EntityDelete(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError) {
	var productcategory models.ProductCategory
	productcategory.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productcategory)

	checkData := db.Debug().First(&productcategory)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &productcategory, <-err
	}

	deleteData := db.Debug().Delete(&productcategory)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &productcategory, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &productcategory, <-err
}

/**
* =============================================================
* Repository Update Master Product CategoryBy ID Teritory
*==============================================================
 */

func (r *repositoryProductCategory) EntityUpdate(input *schemes.ProductCategory) (*models.ProductCategory, schemes.SchemeDatabaseError) {
	var productcategory models.ProductCategory
	productcategory.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&productcategory)

	checkData := db.Debug().First(&productcategory)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &productcategory, <-err
	}

	productcategory.MerchantID = input.MerchantID
	productcategory.OutletID = input.OutletID
	productcategory.Name = input.Name
	productcategory.Active = input.Active

	updateData := db.Debug().Updates(&productcategory)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &productcategory, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &productcategory, <-err
}
