package repositories

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type repositoryOutlite struct {
	db *gorm.DB
}

func NewRepositoryOutlet(db *gorm.DB) *repositoryOutlite {
	return &repositoryOutlite{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityCreate(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet models.Outlet
	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	outlet.Name = input.Name
	outlet.Phone = phone
	outlet.Address = input.Address
	outlet.MerchantID = input.MerchantID
	outlet.Description = input.Description

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	// checkOutletName := db.Debug().First(&outlet, "name = ?", outlet.Name)

	// if checkOutletName.RowsAffected > 0 {
	// 	err <- schemes.SchemeDatabaseError{
	// 		Code: http.StatusConflict,
	// 		Type: "error_create_01",
	// 	}
	// 	return &outlet, <-err
	// }

	checkOutletPhone := db.Debug().First(&outlet, "phone = ?", outlet.Phone)

	if checkOutletPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_02",
		}
		return &outlet, <-err
	}

	addoutlet := db.Debug().Create(&outlet).Commit()

	if addoutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityResults(input *schemes.Outlet) (*[]schemes.GetOutlet, int64, schemes.SchemeDatabaseError) {
	var (
		outlet          []models.Outlet
		result          []schemes.GetOutlet
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "merchant.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(outlet.*) AS count_data
		FROM master.outlets AS outlet
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			outlet.id,
			outlet.name,
			outlet.phone,
			COALESCE(outlet.address,'') AS address,
			COALESCE(outlet.description,'') AS description,
			outlet.created_at,
			outlet.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name
		FROM master.outlets AS outlet
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON outlet.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND outlet.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND outlet.id = ?`
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

func (r *repositoryOutlite) EntityResult(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &outlet, <-err
	}

	// Menggunakan Preload untuk mengisi data relasi Merchant
	db.Preload("Merchant").First(&outlet)

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityDelete(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &outlet, <-err
	}

	// Menggunakan Preload untuk mengisi data relasi Merchant
	db.Preload("Merchant").First(&outlet)

	deleteoutlet := db.Debug().Delete(&outlet)

	if deleteoutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryOutlite) EntityUpdate(input *schemes.Outlet) (*models.Outlet, schemes.SchemeDatabaseError) {
	var outlet models.Outlet
	outlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&outlet)

	checkOutletName := db.Debug().First(&outlet)

	if checkOutletName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &outlet, <-err
	}

	phone, _ := strconv.ParseUint(input.Phone, 10, 64)
	outlet.Name = input.Name
	outlet.Phone = phone
	outlet.Address = input.Address
	outlet.MerchantID = input.MerchantID
	outlet.Description = input.Description
	outlet.Active = input.Active

	updateoutlet := db.Debug().Updates(&outlet)

	if updateoutlet.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &outlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &outlet, <-err
}
