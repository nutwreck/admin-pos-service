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

type repositorySales struct {
	db *gorm.DB
}

func NewRepositorySales(db *gorm.DB) *repositorySales {
	return &repositorySales{db: db}
}

/**
* ==========================================
* Repository Create New Sales Teritory
*===========================================
 */

func (r *repositorySales) EntityCreate(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales models.Sales
	sales.Name = input.Name
	sales.Phone = input.Phone
	sales.Address = input.Address
	sales.Description = input.Description
	sales.MerchantID = input.MerchantID
	sales.OutletID = input.OutletID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&sales)

	checkPhone := db.Debug().First(&sales, "phone = ?", sales.Phone)

	if checkPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &sales, <-err
	}

	addData := db.Debug().Create(&sales).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &sales, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &sales, <-err
}

/**
* ==========================================
* Repository Results All Sales Teritory
*===========================================
 */

func (r *repositorySales) EntityResults(input *schemes.Sales) (*[]schemes.GetSales, int64, schemes.SchemeDatabaseError) {
	var (
		sales           []models.Sales
		result          []schemes.GetSales
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "sales.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&sales)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(sales.*) AS count_data
		FROM master.sales AS sales
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			sales.id,
			sales.name,
			sales.phone,
			COALESCE(sales.address,'') AS address,
			COALESCE(sales.description,'') AS description,
			sales.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			outlet.id AS outlet_id,
			outlet.name AS outlet_name,
			sales.created_at
		FROM master.sales AS sales
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON sales.merchant_id = merchant.id AND merchant.active = true
		JOIN master.outlets AS outlet ON sales.outlet_id = outlet.id AND outlet.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND sales.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.OutletID != constants.EMPTY_VALUE {
		queryAdditional += ` AND sales.outlet_id = ?`
		args = append(args, input.OutletID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND sales.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND sales.id = ?`
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
* Repository Result Sales By ID Teritory
*===========================================
 */

func (r *repositorySales) EntityResult(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales models.Sales
	sales.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&sales)

	checkID := db.Debug().First(&sales)

	if checkID.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &sales, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &sales, <-err
}

/**
* ==========================================
* Repository Delete Sales By ID Teritory
*===========================================
 */

func (r *repositorySales) EntityDelete(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales models.Sales
	sales.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&sales)

	checkID := db.Debug().First(&sales)

	if checkID.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
	}

	// Menggunakan Preload untuk mengisi data relasi Merchant
	db.Preload("Merchant").First(&sales)

	// Menggunakan Preload untuk mengisi data relasi Outlet
	db.Preload("Outlet").First(&sales)

	deleteData := db.Debug().Delete(&sales)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &sales, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &sales, <-err
}

/**
* ==========================================
* Repository Update Sales By ID Teritory
*===========================================
 */

func (r *repositorySales) EntityUpdate(input *schemes.Sales) (*models.Sales, schemes.SchemeDatabaseError) {
	var sales models.Sales
	sales.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&sales)

	checkID := db.Debug().First(&sales)

	if checkID.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &sales, <-err
	}

	sales.Name = input.Name
	sales.Phone = input.Phone
	sales.Address = input.Address
	sales.Description = input.Description
	sales.MerchantID = input.MerchantID
	sales.OutletID = input.OutletID
	sales.Active = input.Active

	updateData := db.Debug().Updates(&sales)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &sales, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &sales, <-err
}
