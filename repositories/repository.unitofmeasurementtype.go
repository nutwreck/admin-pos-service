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

type repositoryUnitOfMeasurementType struct {
	db *gorm.DB
}

func NewRepositoryUnitOfMeasurementType(db *gorm.DB) *repositoryUnitOfMeasurementType {
	return &repositoryUnitOfMeasurementType{db: db}
}

/**
* ===============================================
* Repository Create New Master UOM Type Teritory
*================================================
 */

func (r *repositoryUnitOfMeasurementType) EntityCreate(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError) {
	var uomType models.UnitOfMeasurementType
	uomType.MerchantID = input.MerchantID
	uomType.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uomType)

	checkData := db.Debug().Where("merchant_id = ? AND name = ?", uomType.MerchantID, uomType.Name).First(&uomType)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &uomType, <-err
	}

	addData := db.Debug().Create(&uomType).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &uomType, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &uomType, <-err
}

/**
* ================================================
* Repository Results All Master UOM Type Teritory
*=================================================
 */

func (r *repositoryUnitOfMeasurementType) EntityResults(input *schemes.UnitOfMeasurementType) (*[]schemes.GetUnitOfMeasurementType, int64, schemes.SchemeDatabaseError) {
	var (
		uomType         []models.UnitOfMeasurementType
		result          []schemes.GetUnitOfMeasurementType
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "uomtype.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uomType)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(uomtype.*) AS count_data
		FROM master.unit_of_measurement_types AS uomtype
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			uomtype.id,
			uomtype.name,
			uomtype.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			uomtype.created_at
		FROM master.unit_of_measurement_types AS uomtype
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON uomtype.merchant_id = merchant.id AND merchant.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND uomtype.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND uomtype.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND uomtype.id = ?`
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
* Repository Delete Master UOM Type By ID Teritory
*==================================================
 */

func (r *repositoryUnitOfMeasurementType) EntityDelete(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError) {
	var uomType models.UnitOfMeasurementType
	uomType.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uomType)

	checkData := db.Debug().First(&uomType)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &uomType, <-err
	}

	deleteData := db.Debug().Delete(&uomType)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &uomType, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &uomType, <-err
}

/**
* =================================================
* Repository Update Master UOM Type By ID Teritory
*==================================================
 */

func (r *repositoryUnitOfMeasurementType) EntityUpdate(input *schemes.UnitOfMeasurementType) (*models.UnitOfMeasurementType, schemes.SchemeDatabaseError) {
	var uomType models.UnitOfMeasurementType
	uomType.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uomType)

	checkData := db.Debug().First(&uomType)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &uomType, <-err
	}

	uomType.MerchantID = input.MerchantID
	uomType.Name = input.Name
	uomType.Active = input.Active

	updateData := db.Debug().Updates(&uomType)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &uomType, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &uomType, <-err
}
