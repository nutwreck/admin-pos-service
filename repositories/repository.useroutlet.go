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

type repositoryUserOutlet struct {
	db *gorm.DB
}

func NewRepositoryUserOutlet(db *gorm.DB) *repositoryUserOutlet {
	return &repositoryUserOutlet{db: db}
}

/**
* ===============================================
* Repository Create New User Outlet Teritory
*================================================
 */

func (r *repositoryUserOutlet) EntityCreate(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError) {
	var userOutlet models.UserOutlet
	userOutlet.UserID = input.UserID
	userOutlet.OutletID = input.OutletID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&userOutlet)

	checkData := db.Debug().Where("outlet_id = ? AND user_id = ?", userOutlet.OutletID, userOutlet.UserID).First(&userOutlet)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &userOutlet, <-err
	}

	addData := db.Debug().Create(&userOutlet).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &userOutlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &userOutlet, <-err
}

/**
* ================================================
* Repository Results All User Outlet Teritory
*=================================================
 */

func (r *repositoryUserOutlet) EntityResults(input *schemes.UserOutlet) (*[]schemes.GetAllUserOutlet, int64, schemes.SchemeDatabaseError) {
	var (
		userOutlet      []models.UserOutlet
		result          []schemes.GetAllUserOutlet
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "useroutlet.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&userOutlet)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(useroutlet.*) AS count_data
		FROM master.user_outlets AS useroutlet
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			useroutlet.id,
			users.id AS user_id,
			users.name AS user_name,
			outlet.id AS outlet_id,
			outlet.name AS outlet_name,
			useroutlet.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			useroutlet.created_at
		FROM master.user_outlets AS useroutlet
	`

	queryAdditional = `
		JOIN master.outlets AS outlet ON useroutlet.outlet_id = outlet.id AND outlet.active = true
		JOIN master.merchants AS merchant ON outlet.merchant_id = merchant.id AND merchant.active = true
		JOIN master.users ON useroutlet.user_id = users.id AND users.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND merchant.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.UserID != constants.EMPTY_VALUE {
		queryAdditional += ` AND useroutlet.user_id = ?`
		args = append(args, input.UserID)
	}

	if input.OutletID != constants.EMPTY_VALUE {
		queryAdditional += ` AND useroutlet.outlet_id = ?`
		args = append(args, input.OutletID)
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND role.id = ?`
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
* =================================================
* Repository Delete User Outlet By ID Teritory
*==================================================
 */

func (r *repositoryUserOutlet) EntityDelete(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError) {
	var userOutlet models.UserOutlet
	userOutlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&userOutlet)

	checkId := db.Debug().First(&userOutlet)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &userOutlet, <-err
	}

	deleteData := db.Debug().Delete(&userOutlet)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &userOutlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &userOutlet, <-err
}

/**
* =================================================
* Repository Update User Outlet By ID Teritory
*==================================================
 */

func (r *repositoryUserOutlet) EntityUpdate(input *schemes.UserOutlet) (*models.UserOutlet, schemes.SchemeDatabaseError) {
	var userOutlet models.UserOutlet
	userOutlet.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&userOutlet)

	checkId := db.Debug().First(&userOutlet)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &userOutlet, <-err
	}

	userOutlet.UserID = input.UserID
	userOutlet.OutletID = input.OutletID
	userOutlet.Active = input.Active

	updateData := db.Debug().Updates(&userOutlet)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &userOutlet, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &userOutlet, <-err
}
