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

type repositoryCustomer struct {
	db *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *repositoryCustomer {
	return &repositoryCustomer{db: db}
}

/**
* ==========================================
* Repository Create New Customer Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityCreate(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer models.Customer
	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address
	customer.Description = input.Description
	customer.MerchantID = input.MerchantID
	customer.OutletID = input.OutletID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkPhone := db.Debug().First(&customer, "phone = ?", customer.Phone)

	if checkPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &customer, <-err
	}

	addData := db.Debug().Create(&customer).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Results All Customer Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityResults(input *schemes.Customer) (*[]schemes.GetCustomer, int64, schemes.SchemeDatabaseError) {
	var (
		customer        []models.Customer
		result          []schemes.GetCustomer
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "customer.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(customer.*) AS count_data
		FROM master.customers AS customer
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			customer.id,
			customer.name,
			customer.phone,
			COALESCE(customer.address,'') AS address,
			COALESCE(customer.description,'') AS description,
			customer.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			outlet.id AS outlet_id,
			outlet.name AS outlet_name,
			customer.created_at
		FROM master.customers AS customer
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON customer.merchant_id = merchant.id AND merchant.active = true
		JOIN master.outlets AS outlet ON customer.outlet_id = outlet.id AND outlet.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND customer.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.OutletID != constants.EMPTY_VALUE {
		queryAdditional += ` AND customer.outlet_id = ?`
		args = append(args, input.OutletID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND customer.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND customer.id = ?`
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
* Repository Result Customer By ID Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityResult(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer models.Customer
	customer.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkID := db.Debug().First(&customer)

	if checkID.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Delete Customer By ID Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityDelete(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer models.Customer
	customer.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkID := db.Debug().First(&customer)

	if checkID.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
	}

	// Menggunakan Preload untuk mengisi data relasi Merchant
	db.Preload("Merchant").First(&customer)

	// Menggunakan Preload untuk mengisi data relasi Outlet
	db.Preload("Outlet").First(&customer)

	deleteData := db.Debug().Delete(&customer)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}

/**
* ==========================================
* Repository Update Customer By ID Teritory
*===========================================
 */

func (r *repositoryCustomer) EntityUpdate(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer models.Customer
	customer.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&customer)

	checkID := db.Debug().First(&customer)

	if checkID.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &customer, <-err
	}

	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address
	customer.Description = input.Description
	customer.MerchantID = input.MerchantID
	customer.OutletID = input.OutletID
	customer.Active = input.Active

	updateData := db.Debug().Updates(&customer)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &customer, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &customer, <-err
}
