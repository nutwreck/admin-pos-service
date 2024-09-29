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

type repositoryMappingRoleMenuUser struct {
	db *gorm.DB
}

func NewRepositoryMappingRoleMenuUser(db *gorm.DB) *repositoryMappingRoleMenuUser {
	return &repositoryMappingRoleMenuUser{db: db}
}

/**
* ===============================================
* Repository Create New Mapping Role Menu User Teritory
*================================================
 */

func (r *repositoryMappingRoleMenuUser) EntityCreate(input *[]schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError) {
	err := make(chan schemes.SchemeDatabaseError, 1)

	// Mulai transaksi
	tx := r.db.Begin()

	for _, input := range *input {
		var mappingRoleMenuUser models.MappingRoleMenuUser
		mappingRoleMenuUser.UserID = input.UserID
		mappingRoleMenuUser.MappingRoleMenuID = input.MappingRoleMenuID
		mappingRoleMenuUser.MerchantID = input.MerchantID

		db := tx.Model(&mappingRoleMenuUser)

		checkRoleName := db.Debug().Where("merchant_id = ? AND user_id = ? AND mapping_role_menu_id = ?", mappingRoleMenuUser.MerchantID, mappingRoleMenuUser.UserID, mappingRoleMenuUser.MappingRoleMenuID).First(&mappingRoleMenuUser)

		if checkRoleName.RowsAffected > 0 {
			tx.Rollback()
			err <- schemes.SchemeDatabaseError{
				Code: http.StatusConflict,
				Type: "error_create_01",
			}
			return nil, <-err
		}

		addData := db.Debug().Create(&mappingRoleMenuUser)

		if addData.RowsAffected < 1 {
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
* ================================================
* Repository Results All Mapping Role Menu User Teritory
*=================================================
 */

func (r *repositoryMappingRoleMenuUser) EntityResults(input *schemes.MappingRoleMenuUser) (*[]schemes.GetMappingRoleMenuUser, int64, schemes.SchemeDatabaseError) {
	var (
		role            []models.MappingRoleMenuUser
		result          []schemes.GetMappingRoleMenuUser
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "mappingrolemenuuser.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(mappingrolemenuuser.*) AS count_data
		FROM master.mapping_role_menu_users AS mappingrolemenuuser
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			mappingrolemenuuser.id,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			user.id AS user_id,
			user.name AS user_name,
			mappingrolemenu.id AS mapping_role_menu_id,
			mappingrolemenu.name AS mapping_role_menu_name,
			mappingrolemenuuser.created_at
		FROM master.mapping_role_menu_users AS mappingrolemenuuser
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON mappingrolemenuuser.merchant_id = merchant.id AND merchant.active = true
		JOIN master.users AS user ON mappingrolemenuuser.user_id = user.id AND user.active = true
		JOIN master.mapping_role_menus AS mappingrolemenu ON mappingrolemenuuser.mapping_role_menu_id = mappingrolemenu.id
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND mappingrolemenuuser.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND mappingrolemenuuser.id = ?`
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
* Repository Delete Mapping Role Menu User By ID Teritory
*==================================================
 */

func (r *repositoryMappingRoleMenuUser) EntityDelete(input *schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError) {
	var mappingRoleMenuUser models.MappingRoleMenuUser
	mappingRoleMenuUser.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&mappingRoleMenuUser)

	checkId := db.Debug().First(&mappingRoleMenuUser)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &mappingRoleMenuUser, <-err
	}

	deleteRole := db.Debug().Delete(&mappingRoleMenuUser)

	if deleteRole.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &mappingRoleMenuUser, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &mappingRoleMenuUser, <-err
}

/**
* =================================================
* Repository Update Mapping Role Menu User By ID Teritory
*==================================================
 */

func (r *repositoryMappingRoleMenuUser) EntityUpdate(input *schemes.MappingRoleMenuUser) (*models.MappingRoleMenuUser, schemes.SchemeDatabaseError) {
	var mappingRoleMenuUser models.MappingRoleMenuUser
	mappingRoleMenuUser.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&mappingRoleMenuUser)

	checkId := db.Debug().First(&mappingRoleMenuUser)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &mappingRoleMenuUser, <-err
	}

	mappingRoleMenuUser.UserID = input.UserID
	mappingRoleMenuUser.MappingRoleMenuID = input.MappingRoleMenuID
	mappingRoleMenuUser.MerchantID = input.MerchantID
	mappingRoleMenuUser.Active = input.Active

	updateRole := db.Debug().Updates(&mappingRoleMenuUser)

	if updateRole.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &mappingRoleMenuUser, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &mappingRoleMenuUser, <-err
}
