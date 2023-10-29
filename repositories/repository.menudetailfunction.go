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

type repositoryMenuDetailFunction struct {
	db *gorm.DB
}

func NewRepositoryMenuDetailFunction(db *gorm.DB) *repositoryMenuDetailFunction {
	return &repositoryMenuDetailFunction{db: db}
}

/**
* ==========================================================
* Repository Create New Master Menu Detail Function Teritory
*===========================================================
 */

func (r *repositoryMenuDetailFunction) EntityCreate(input *[]schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	err := make(chan schemes.SchemeDatabaseError, 1)

	// Mulai transaksi
	tx := r.db.Begin()

	for _, input := range *input {
		var menuDetailFunction models.MenuDetailFunction
		menuDetailFunction.MerchantID = input.MerchantID
		menuDetailFunction.Name = input.Name
		menuDetailFunction.MenuID = input.MenuID
		menuDetailFunction.Link = input.Link
		menuDetailFunction.MenuDetailID = input.MenuDetailID

		db := tx.Model(&menuDetailFunction)

		checkName := db.Debug().Where("merchant_id = ? AND name = ? AND menu_id = ? AND menu_detail_id = ?", menuDetailFunction.MerchantID, menuDetailFunction.Name, menuDetailFunction.MenuID, menuDetailFunction.MenuDetailID).First(&menuDetailFunction)

		if checkName.RowsAffected > 0 {
			// Rollback transaksi jika ada kesalahan
			tx.Rollback()
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusConflict,
				Type: "error_create_01",
			}
			return nil, <-err
		}

		add := db.Debug().Create(&menuDetailFunction)

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
* ===========================================================
* Repository Results All Master Menu Detail Function Teritory
*============================================================
 */

func (r *repositoryMenuDetailFunction) EntityResults(input *schemes.MenuDetailFunction) (*[]schemes.GetMenuDetailFunction, int64, schemes.SchemeDatabaseError) {
	var (
		menuDetailFunction []models.MenuDetailFunction
		result             []schemes.GetMenuDetailFunction
		countData          schemes.CountData
		args               []interface{}
		totalData          int64
		sortData           string = "menudetailfunction.created_at DESC"
		queryCountData     string = constants.EMPTY_VALUE
		queryData          string = constants.EMPTY_VALUE
		queryAdditional    string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetailFunction)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(menudetailfunction.*) AS count_data
		FROM master.menu_detail_functions AS menudetailfunction
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			menudetailfunction.id,
			menudetailfunction.name,
			menudetailfunction.link,
			menu.id AS menu_id,
			menu.name AS menu_name,
			menudetail.id AS menu_detail_id,
			menudetail.name AS menu_detail_name,
			menudetailfunction.active,
			menudetailfunction.created_at,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name
		FROM master.menu_detail_functions AS menudetailfunction
	`

	queryAdditional = `
		JOIN master.menus AS menu ON menudetailfunction.menu_id = menu.id AND menu.active = true
		JOIN master.menu_details AS menudetail ON menudetailfunction.menu_detail_id = menudetail.id AND menudetail.active = true
		JOIN master.merchants AS merchant ON menudetailfunction.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetailfunction.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetailfunction.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetailfunction.id = ?`
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
* ============================================================
* Repository Delete Master Menu Detail Function By ID Teritory
*=============================================================
 */

func (r *repositoryMenuDetailFunction) EntityDelete(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var menuDetailFunction models.MenuDetailFunction
	menuDetailFunction.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetailFunction)

	checkId := db.Debug().First(&menuDetailFunction)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &menuDetailFunction, <-err
	}

	delete := db.Debug().Delete(&menuDetailFunction)

	if delete.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &menuDetailFunction, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menuDetailFunction, <-err
}

/**
* ============================================================
* Repository Update Master Menu Detail Function By ID Teritory
*=============================================================
 */

func (r *repositoryMenuDetailFunction) EntityUpdate(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var menuDetailFunction models.MenuDetailFunction
	menuDetailFunction.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetailFunction)

	checkMenuId := db.Debug().First(&menuDetailFunction)

	if checkMenuId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &menuDetailFunction, <-err
	}

	menuDetailFunction.MerchantID = input.MerchantID
	menuDetailFunction.Name = input.Name
	menuDetailFunction.Link = input.Link
	menuDetailFunction.MenuID = input.MenuID
	menuDetailFunction.MenuDetailID = input.MenuDetailID
	menuDetailFunction.Active = input.Active

	updateMenu := db.Debug().Updates(&menuDetailFunction)

	if updateMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &menuDetailFunction, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menuDetailFunction, <-err
}

/**
* ============================================================
* Repository Result Master Menu Detail Function By ID Teritory
*=============================================================
 */
func (r *repositoryMenuDetailFunction) EntityResult(input *schemes.MenuDetailFunction) (*models.MenuDetailFunction, schemes.SchemeDatabaseError) {
	var menuDetailFunction models.MenuDetailFunction
	menuDetailFunction.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetailFunction)

	getData := db.Debug().First(&menuDetailFunction, "id = ?", input.ID)

	if getData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &menuDetailFunction, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menuDetailFunction, <-err
}
