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

type repositoryRole struct {
	db *gorm.DB
}

func NewRepositoryRole(db *gorm.DB) *repositoryRole {
	return &repositoryRole{db: db}
}

/**
* ===============================================
* Repository Create New Master Role Teritory
*================================================
 */

func (r *repositoryRole) EntityCreate(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role models.Role
	role.Name = input.Name
	role.Type = input.Type
	role.MerchantID = input.MerchantID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleName := db.Debug().Where("merchant_id = ? AND name = ? AND type = ?", role.MerchantID, role.Name, role.Type).First(&role)

	if checkRoleName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &role, <-err
	}

	addRole := db.Debug().Create(&role).Commit()

	if addRole.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &role, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &role, <-err
}

/**
* ================================================
* Repository Results All Master Role Teritory
*=================================================
 */

func (r *repositoryRole) EntityResults(input *schemes.Role) (*[]schemes.GetAllRole, int64, schemes.SchemeDatabaseError) {
	var (
		role            []models.Role
		result          []schemes.GetAllRole
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "role.created_at DESC"
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
			COUNT(role.*) AS count_data
		FROM master.roles AS role
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			role.id,
			role.name,
			role.type,
			role.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			role.created_at
		FROM master.roles AS role
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON role.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND role.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND role.type LIKE ?`
		args = append(args, input.Name)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND role.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
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
* Repository Delete Master Role By ID Teritory
*==================================================
 */

func (r *repositoryRole) EntityDelete(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role models.Role
	role.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleId := db.Debug().First(&role)

	if checkRoleId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &role, <-err
	}

	deleteRole := db.Debug().Delete(&role)

	if deleteRole.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &role, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &role, <-err
}

/**
* =================================================
* Repository Update Master Role By ID Teritory
*==================================================
 */

func (r *repositoryRole) EntityUpdate(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var role models.Role
	role.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleId := db.Debug().First(&role)

	if checkRoleId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &role, <-err
	}

	role.Name = input.Name
	role.Type = input.Type
	role.MerchantID = input.MerchantID
	role.Active = input.Active

	updateRole := db.Debug().Updates(&role)

	if updateRole.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &role, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &role, <-err
}
