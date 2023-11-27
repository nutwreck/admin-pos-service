package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceProduct struct {
	product entities.EntityProduct
}

func NewServiceProduct(product entities.EntityProduct) *serviceProduct {
	return &serviceProduct{product: product}
}

/**
* ==============================================
* Service Create New Product Teritory
*===============================================
 */

func (s *serviceProduct) EntityCreate(inputs *[]schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var createdProducts []schemes.Product

	// Loop through each input in the batch
	for _, input := range *inputs {
		var product schemes.Product

		product.MerchantID = input.MerchantID
		product.OutletID = input.OutletID
		product.ProductCategoryID = input.ProductCategoryID
		product.ProductCategorySubID = input.ProductCategorySubID
		product.Code = input.Code
		product.Name = input.Name
		product.Barcode = input.Barcode
		product.CapitalPrice = input.CapitalPrice
		product.SellingPrice = input.SellingPrice
		product.SupplierID = input.SupplierID
		product.UnitOfMeasurementID = input.UnitOfMeasurementID
		product.Image = input.Image
		product.Active = input.Active

		// Collect the created Product objects
		createdProducts = append(createdProducts, product)
	}

	res, err := s.product.EntityCreate(&createdProducts)
	return res, err
}

/**
* ===============================================
* Service Results All Product Teritory
*================================================
 */

func (s *serviceProduct) EntityResults(input *schemes.Product) (*[]schemes.GetProduct, int64, schemes.SchemeDatabaseError) {
	var product schemes.Product
	product.Sort = input.Sort
	product.Page = input.Page
	product.PerPage = input.PerPage
	product.MerchantID = input.MerchantID
	product.OutletID = input.OutletID
	product.ProductCategoryID = input.ProductCategoryID
	product.ProductCategorySubID = input.ProductCategorySubID
	product.Code = input.Code
	product.Name = input.Name
	product.Barcode = input.Barcode
	product.CapitalPrice = input.CapitalPrice
	product.SellingPrice = input.SellingPrice
	product.SupplierID = input.SupplierID
	product.UnitOfMeasurementID = input.UnitOfMeasurementID
	product.Image = input.Image
	product.Active = input.Active
	product.ID = input.ID

	res, totalData, err := s.product.EntityResults(&product)
	return res, totalData, err
}

/**
* ================================================
* Service Delete Product By ID Teritory
*=================================================
 */

func (s *serviceProduct) EntityDelete(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var product schemes.Product
	product.ID = input.ID

	res, err := s.product.EntityDelete(&product)
	return res, err
}

/**
* ================================================
* Service Update Product By ID Teritory
*=================================================
 */

func (s *serviceProduct) EntityUpdate(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var product schemes.Product
	product.ID = input.ID
	product.MerchantID = input.MerchantID
	product.OutletID = input.OutletID
	product.ProductCategoryID = input.ProductCategoryID
	product.ProductCategorySubID = input.ProductCategorySubID
	product.Code = input.Code
	product.Name = input.Name
	product.Barcode = input.Barcode
	product.CapitalPrice = input.CapitalPrice
	product.SellingPrice = input.SellingPrice
	product.SupplierID = input.SupplierID
	product.UnitOfMeasurementID = input.UnitOfMeasurementID
	product.Image = input.Image
	product.Active = input.Active

	res, err := s.product.EntityUpdate(&product)
	return res, err
}

/**
* ==========================================
* Service Result Product Teritory
*===========================================
 */

func (s *serviceProduct) EntityResult(input *schemes.Product) (*models.Product, schemes.SchemeDatabaseError) {
	var product schemes.Product
	product.ID = input.ID

	res, err := s.product.EntityResult(&product)
	return res, err
}
