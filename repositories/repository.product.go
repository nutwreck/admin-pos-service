package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db: db}
}

/**
* =================================================
* Repository Create New Master Product Teritory
*==================================================
 */

func (r *repositoryProduct) EntityCreate(input *[]schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	err := make(chan schemes.SchemeDatabaseError, 1)

	// Mulai transaksi
	tx := r.db.Begin()

	for _, input := range *input {
		var product models.Product
		product.MerchantID = input.MerchantID
		product.OutletID = input.OutletID
		product.ProductCategoryID = input.ProductCategoryID
		product.ProductCategorySubID = input.ProductCategorySubID
		product.Code = input.Code
		product.Name = input.Name
		product.Barcode = input.Barcode
		product.CapitalPrice = input.CapitalPrice
		product.SellingPrice = input.SellingPrice
		product.SupplierID = input.SupplierID
		product.UnitOfMeasurementID = input.UnitOfMeasurementID
		product.Image = input.Image
		product.Active = input.Active

		db := tx.Model(&product)

		checkData := db.Debug().Where("merchant_id = ? AND name = ? AND outlet_id = ? AND product_category_id = ? AND product_category_sub_id = ?", product.MerchantID, product.Name, product.OutletID, product.ProductCategoryID, product.ProductCategorySubID).First(&product)

		if checkData.RowsAffected > 0 {
			// Rollback transaksi jika ada kesalahan
			tx.Rollback()
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusConflict,
				Type: "error_create_01",
			}
			return nil, <-err
		}

		add := db.Debug().Create(&product)

		if add.RowsAffected < 1 {
			// Rollback transaksi jika ada kesalahan
			tx.Rollback()
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusForbidden,
				Type: "error_create_02",
			}
			return nil, <-err
		}
	}

	// Commit transaksi jika semuanya berhasil
	tx.Commit()

	err <- schemes.SchemeDatabaseError{}
	return nil, <-err
}

/**
* ==================================================
* Repository Results All Master Product Teritory
*===================================================
 */

func (r *repositoryProduct) EntityResults(input *schemes.Product) (*[]schemes.GetProduct, int64, schemes.SchemeDatabaseError) {
	var (
		product         []models.Product
		result          []schemes.GetProduct
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "product.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(product.*) AS count_data
		FROM master.products AS product
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT 
			product.id,
			product.merchant_id,
			merchant.name AS merchant_name,
			product.outlet_id,
			outlet.name AS outlet_name,
			product.product_category_id,
			product_category.name AS product_category_name,
			product.product_category_sub_id,
			product_category_sub.name AS product_category_sub_name,
			product.code,
			product."name",
			product.barcode,
			product.capital_price,
			product.selling_price,
			coalesce(product.supplier_id,'') AS supplier_id,
			coalesce(supplier.name,'') AS supplier_name,
			product.unit_of_measurement_id,
			unit_of_measurement.name AS unit_of_measurement_name,
			CASE
				WHEN product.image != '' THEN CONCAT('` + configs.AccessFile + `',product.image)
				ELSE ''
			END AS image,
			product.active
		FROM master.products product
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON product.merchant_id = merchant.id AND merchant.active = true
		JOIN master.outlets AS outlet ON product.outlet_id = outlet.id AND outlet.active = true
		JOIN master.product_categorys AS product_category ON product.product_category_id  = product_category.id
		JOIN master.product_category_subs AS product_category_sub ON product.product_category_sub_id = product_category_sub.id
		LEFT JOIN master.suppliers as supplier on product.supplier_id = supplier.id
		JOIN master.unit_of_measurements as unit_of_measurement on product.unit_of_measurement_id = unit_of_measurement.id
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.OutletID != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.outlet_id = ?`
		args = append(args, input.OutletID)
	}

	if input.ProductCategoryID != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.product_category_id = ?`
		args = append(args, input.ProductCategoryID)
	}

	if input.ProductCategorySubID != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.product_category_sub_id = ?`
		args = append(args, input.ProductCategorySubID)
	}

	if input.UnitOfMeasurementID != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.unit_of_measurement_id = ?`
		args = append(args, input.UnitOfMeasurementID)
	}

	if input.Code != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.code = ?`
		args = append(args, input.Code)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND product.id = ?`
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
* ===================================================
* Repository Delete Master Product By ID Teritory
*====================================================
 */

func (r *repositoryProduct) EntityDelete(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var product models.Product
	product.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkData := db.Debug().First(&product)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &product, <-err
	}

	delete := db.Debug().Delete(&product)

	if delete.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ===================================================
* Repository Update Master Product By ID Teritory
*====================================================
 */

func (r *repositoryProduct) EntityUpdate(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var product models.Product
	product.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	checkMenuId := db.Debug().First(&product)

	if checkMenuId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &product, <-err
	}

	product.MerchantID = input.MerchantID
	product.OutletID = input.OutletID
	product.ProductCategoryID = input.ProductCategoryID
	product.ProductCategorySubID = input.ProductCategorySubID
	product.Code = input.Code
	product.Name = input.Name
	product.Barcode = input.Barcode
	product.CapitalPrice = input.CapitalPrice
	product.SellingPrice = input.SellingPrice
	product.SupplierID = input.SupplierID
	product.UnitOfMeasurementID = input.UnitOfMeasurementID
	product.Image = input.Image
	product.Active = input.Active

	updateMenu := db.Debug().Updates(&product)

	if updateMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}

/**
* ==================================================
* Repository Result Master Product By ID Teritory
*===================================================
 */
func (r *repositoryProduct) EntityResult(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var product models.Product
	product.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&product)

	getData := db.Debug().First(&product, "id = ?", input.ID)

	if getData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &product, <-err
}
