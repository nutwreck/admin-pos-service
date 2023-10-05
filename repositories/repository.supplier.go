package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) *repositorySupplier {
	return &repositorySupplier{db: db}
}

/**
* ==========================================
* Repository Create New Supplier Teritory
*===========================================
 */

func (r *repositorySupplier) EntityCreate(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier models.Supplier
	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address
	supplier.Description = input.Description
	supplier.MerchantID = input.MerchantID
	supplier.OutletID = input.OutletID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierPhone := db.Debug().First(&supplier, "phone = ?", supplier.Phone)

	if checkSupplierPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &supplier, <-err
	}

	addSupplier := db.Debug().Create(&supplier).Commit()

	if addSupplier.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &supplier, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Results All Supplier Teritory
*===========================================
 */

func (r *repositorySupplier) EntityResults(input *schemes.Supplier) (*[]schemes.GetSupplier, int64, schemes.SchemeDatabaseError) {
	var (
		supplier        []models.Supplier
		result          []schemes.GetSupplier
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "merchant.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&supplier)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(supplier.*) AS count_data
		FROM master.suppliers AS supplier
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			supplier.id,
			supplier.name,
			supplier.phone,
			COALESCE(supplier.address,'') AS address,
			COALESCE(supplier.description,'') AS description,
			supplier.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			outlet.id AS outlet_id,
			outlet.name AS outlet_name
		FROM master.suppliers AS supplier
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON supplier.merchant_id = merchant.id AND merchant.active = true
		JOIN master.outlets AS outlet ON supplier.outlet_id = outlet.id AND outlet.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND supplier.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ID != uint64(constants.EMPTY_NUMBER) {
		queryAdditional += ` AND supplier.id = ?`
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
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityResult(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier models.Supplier
	supplier.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierId := db.Debug().First(&supplier)

	if checkSupplierId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &supplier, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityDelete(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier models.Supplier
	supplier.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierId := db.Debug().First(&supplier)

	if checkSupplierId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
	}

	// Menggunakan Preload untuk mengisi data relasi Merchant
	db.Preload("Merchant").First(&supplier)

	// Menggunakan Preload untuk mengisi data relasi Outlet
	db.Preload("Outlet").First(&supplier)

	deleteSupplier := db.Debug().Delete(&supplier)

	if deleteSupplier.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &supplier, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityUpdate(input *schemes.Supplier) (*models.Supplier, schemes.SchemeDatabaseError) {
	var supplier models.Supplier
	supplier.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierId := db.Debug().First(&supplier)

	if checkSupplierId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &supplier, <-err
	}

	supplier.Name = input.Name
	supplier.Phone = input.Phone
	supplier.Address = input.Address
	supplier.Description = input.Description
	supplier.MerchantID = input.MerchantID
	supplier.OutletID = input.OutletID
	supplier.Active = input.Active

	updateSupplier := db.Debug().Updates(&supplier)

	if updateSupplier.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &supplier, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &supplier, <-err
}
