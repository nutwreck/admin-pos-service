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

type repositoryMenuDetail struct {
	db *gorm.DB
}

func NewRepositoryMenuDetail(db *gorm.DB) *repositoryMenuDetail {
	return &repositoryMenuDetail{db: db}
}

/**
* =================================================
* Repository Create New Master Menu Detail Teritory
*==================================================
 */

func (r *repositoryMenuDetail) EntityCreate(input *[]schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	err := make(chan schemes.SchemeDatabaseError, 1)

	// Mulai transaksi
	tx := r.db.Begin()

	for _, input := range *input {
		var menuDetail models.MenuDetail
		menuDetail.MerchantID = input.MerchantID
		menuDetail.Name = input.Name
		menuDetail.MenuID = input.MenuID
		menuDetail.Link = input.Link
		menuDetail.Image = input.Image
		menuDetail.Icon = input.Icon

		db := tx.Model(&menuDetail)

		checkMenuName := db.Debug().Where("merchant_id = ? AND name = ? AND menu_id = ?", menuDetail.MerchantID, menuDetail.Name, menuDetail.MenuID).First(&menuDetail)

		if checkMenuName.RowsAffected > 0 {
			// Rollback transaksi jika ada kesalahan
			tx.Rollback()
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusConflict,
				Type: "error_create_01",
			}
			return nil, <-err
		}

		addmenu := db.Debug().Create(&menuDetail)

		if addmenu.RowsAffected < 1 {
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
* Repository Results All Master Menu Detail Teritory
*===================================================
 */

func (r *repositoryMenuDetail) EntityResults(input *schemes.MenuDetail) (*[]schemes.GetMenuDetail, int64, schemes.SchemeDatabaseError) {
	var (
		menuDetail      []models.MenuDetail
		result          []schemes.GetMenuDetail
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "menudetail.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetail)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(menudetail.*) AS count_data
		FROM master.menu_details AS menudetail
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			menudetail.id,
			menudetail.menu_id,
			COALESCE(menu.name,'MENU NOT FOUND') AS menu_name,
			menudetail.name,
			menudetail.link,
			CASE
				WHEN menudetail.image != '' THEN CONCAT('` + configs.AccessFile + `',menudetail.image)
				ELSE ''
			END AS image,
			CASE
				WHEN menudetail.icon != '' THEN CONCAT('` + configs.AccessFile + `',menudetail.icon)
				ELSE ''
			END AS icon,
			menudetail.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			menudetail.created_at
		FROM master.menu_details AS menudetail
	`

	queryAdditional = `
		JOIN master.menus AS menu ON menudetail.menu_id = menu.id AND menu.active = true
		JOIN master.merchants AS merchant ON menudetail.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetail.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetail.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND menudetail.id = ?`
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
* Repository Delete Master Menu Detail By ID Teritory
*====================================================
 */

func (r *repositoryMenuDetail) EntityDelete(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail models.MenuDetail
	menuDetail.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetail)

	checkMenuId := db.Debug().First(&menuDetail)

	if checkMenuId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &menuDetail, <-err
	}

	deleteMenu := db.Debug().Delete(&menuDetail)

	if deleteMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &menuDetail, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menuDetail, <-err
}

/**
* ===================================================
* Repository Update Master Menu Detail By ID Teritory
*====================================================
 */

func (r *repositoryMenuDetail) EntityUpdate(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail models.MenuDetail
	menuDetail.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetail)

	checkMenuId := db.Debug().First(&menuDetail)

	if checkMenuId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &menuDetail, <-err
	}

	menuDetail.MerchantID = input.MerchantID
	menuDetail.Name = input.Name
	menuDetail.Link = input.Link
	menuDetail.Image = input.Image
	menuDetail.Icon = input.Icon
	menuDetail.Active = input.Active

	updateMenu := db.Debug().Updates(&menuDetail)

	if updateMenu.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &menuDetail, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menuDetail, <-err
}

/**
* ==================================================
* Repository Result Master Menu Detail By ID Teritory
*===================================================
 */
func (r *repositoryMenuDetail) EntityResult(input *schemes.MenuDetail) (*models.MenuDetail, schemes.SchemeDatabaseError) {
	var menuDetail models.MenuDetail
	menuDetail.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&menuDetail)

	getData := db.Debug().First(&menuDetail, "id = ?", input.ID)

	if getData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &menuDetail, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &menuDetail, <-err
}
