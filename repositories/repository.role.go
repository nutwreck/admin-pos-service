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

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	checkRoleName := db.Debug().Where("name = ? AND type = ?", role.Name, role.Type).First(&role)

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

func (r *repositoryRole) EntityResults(input *schemes.Role) (*[]models.Role, int64, schemes.SchemeDatabaseError) {
	var (
		role      []models.Role
		totalData int64
		sort      string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&role)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sort = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	if input.Type != constants.EMPTY_VALUE {
		db = db.Where("type = ?", input.Type)
	}

	if input.Name != constants.EMPTY_VALUE {
		db = db.Where("name LIKE ?", "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		db = db.Where("id LIKE ?", "%"+input.ID+"%")
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkRole := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&role)

	if checkRole.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &role, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.Menu{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &role, totalData, <-err
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
