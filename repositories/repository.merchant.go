package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type repositoryMerchant struct {
	db *gorm.DB
}

func NewRepositoryMerchant(db *gorm.DB) *repositoryMerchant {
	return &repositoryMerchant{db: db}
}

/**
* ==========================================
* Repository Create New Merchant Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityCreate(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant models.Merchant
	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.Description = input.Description

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	// checkMerchantName := db.Debug().First(&merchant, "name = ?", input.Name)

	// if checkMerchantName.RowsAffected > 0 {
	// 	err <- schemes.SchemeDatabaseError{
	// 		Code: http.StatusConflict,
	// 		Type: "error_create_01",
	// 	}
	// 	return &merchant, <-err
	// }

	checkMerchantPhone := db.Debug().First(&merchant, "phone = ?", input.Phone)

	if checkMerchantPhone.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_02",
		}
		return &merchant, <-err
	}

	addMerchant := db.Debug().Create(&merchant).Commit()

	if addMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Results All Merchant Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityResults(input *schemes.Merchant) (*[]schemes.GetMerchant, int64, schemes.SchemeDatabaseError) {
	var (
		merchant        []models.Merchant
		result          []schemes.GetMerchant
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "merchant.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(merchant.*) AS count_data
		FROM master.merchants AS merchant
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			merchant.id,
			merchant.name,
			merchant.phone,
			COALESCE(merchant.address,'') AS address,
			COALESCE(merchant.description,'') AS description,
			CASE
				WHEN merchant.logo != '' THEN CONCAT('` + configs.AccessFile + `',merchant.logo)
				ELSE ''
			END AS logo,
			merchant.active,
			merchant.created_at
		FROM master.merchants AS merchant
	`

	queryAdditional = ` WHERE TRUE`

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND merchant.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND merchant.id = ?`
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

func (r *repositoryMerchant) EntityResult(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant models.Merchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchant := db.Debug().First(&merchant)

	if checkMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityDelete(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant models.Merchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchant := db.Debug().First(&merchant)

	if checkMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &merchant, <-err
	}

	deleteMerchant := db.Debug().Delete(&merchant)

	if deleteMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryMerchant) EntityUpdate(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var merchant models.Merchant
	merchant.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&merchant)

	checkMerchantName := db.Debug().First(&merchant)

	if checkMerchantName.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &merchant, <-err
	}

	merchant.Name = input.Name
	merchant.Phone = input.Phone
	merchant.Address = input.Address
	merchant.Logo = input.Logo
	merchant.Description = input.Description
	merchant.Active = input.Active

	updateMerchant := db.Debug().Updates(&merchant)

	if updateMerchant.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &merchant, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &merchant, <-err
}
