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

type repositoryUnitOfMeasurement struct {
	db *gorm.DB
}

func NewRepositoryUnitOfMeasurement(db *gorm.DB) *repositoryUnitOfMeasurement {
	return &repositoryUnitOfMeasurement{db: db}
}

/**
* ===============================================
* Repository Create New Master UOM Teritory
*================================================
 */

func (r *repositoryUnitOfMeasurement) EntityCreate(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError) {
	var uom models.UnitOfMeasurement
	uom.MerchantID = input.MerchantID
	uom.UOMTypeID = input.UOMTypeID
	uom.Symbol = input.Symbol
	uom.ConversionFactor = input.ConversionFactor
	uom.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uom)

	checkData := db.Debug().Where("merchant_id = ? AND uom_type_id = ? AND name = ?", uom.MerchantID, uom.UOMTypeID, uom.Name).First(&uom)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &uom, <-err
	}

	addData := db.Debug().Create(&uom).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &uom, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &uom, <-err
}

/**
* ================================================
* Repository Results All Master UOM Teritory
*=================================================
 */

func (r *repositoryUnitOfMeasurement) EntityResults(input *schemes.UnitOfMeasurement) (*[]schemes.GetUnitOfMeasurement, int64, schemes.SchemeDatabaseError) {
	var (
		uom             []models.UnitOfMeasurement
		result          []schemes.GetUnitOfMeasurement
		countData       schemes.CountData
		args            []interface{}
		totalData       int64
		sortData        string = "uom.created_at DESC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uom)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(uom.*) AS count_data
		FROM master.unit_of_measurements AS uom
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			uom.id,
			uom.name,
			uom.symbol,
			uom.conversion_factor,
			uom.active,
			merchant.id AS merchant_id,
			merchant.name AS merchant_name,
			uomType.id AS uom_type_id,
			uomType.name AS uom_type_name,
			uom.created_at
		FROM master.unit_of_measurements AS uom
	`

	queryAdditional = `
		JOIN master.merchants AS merchant ON uom.merchant_id = merchant.id AND merchant.active = true
		JOIN master.unit_of_measurement_types AS uomType ON uom.uom_type_id = uomType.id AND uomType.active = true
	`

	queryAdditional += ` WHERE TRUE`

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND uom.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.UOMTypeID != constants.EMPTY_VALUE {
		queryAdditional += ` AND uom.uom_type_id = ?`
		args = append(args, input.UOMTypeID)
	}

	if input.MerchantID != constants.EMPTY_VALUE {
		queryAdditional += ` AND uom.merchant_id = ?`
		args = append(args, input.MerchantID)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND uom.name LIKE ?`
		args = append(args, "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		queryAdditional += ` AND uom.id = ?`
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
* Repository Delete Master UOM By ID Teritory
*==================================================
 */

func (r *repositoryUnitOfMeasurement) EntityDelete(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError) {
	var uom models.UnitOfMeasurement
	uom.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uom)

	checkData := db.Debug().First(&uom)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &uom, <-err
	}

	deleteData := db.Debug().Delete(&uom)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &uom, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &uom, <-err
}

/**
* =================================================
* Repository Update Master UOM By ID Teritory
*==================================================
 */

func (r *repositoryUnitOfMeasurement) EntityUpdate(input *schemes.UnitOfMeasurement) (*models.UnitOfMeasurement, schemes.SchemeDatabaseError) {
	var uom models.UnitOfMeasurement
	uom.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&uom)

	checkData := db.Debug().First(&uom)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &uom, <-err
	}

	uom.MerchantID = input.MerchantID
	uom.UOMTypeID = input.UOMTypeID
	uom.Symbol = input.Symbol
	uom.ConversionFactor = input.ConversionFactor
	uom.Name = input.Name
	uom.Active = input.Active

	updateData := db.Debug().Updates(&uom)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &uom, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &uom, <-err
}
