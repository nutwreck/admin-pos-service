package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceCustomer struct {
	customer entities.EntityCustomer
}

func NewServiceCustomer(customer entities.EntityCustomer) *serviceCustomer {
	return &serviceCustomer{customer: customer}
}

/**
* ==========================================
* Service Create New Customer Teritory
*===========================================
 */

func (s *serviceCustomer) EntityCreate(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer schemes.Customer
	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address
	customer.Description = input.Description
	customer.MerchantID = input.MerchantID
	customer.OutletID = input.OutletID

	res, err := s.customer.EntityCreate(&customer)
	return res, err
}

/**
* ==========================================
* Service Results All Customer Teritory
*===========================================
 */

func (s *serviceCustomer) EntityResults(input *schemes.Customer) (*[]schemes.GetCustomer, int64, schemes.SchemeDatabaseError) {
	var customer schemes.Customer
	customer.Sort = input.Sort
	customer.Page = input.Page
	customer.PerPage = input.PerPage
	customer.MerchantID = input.MerchantID
	customer.OutletID = input.OutletID
	customer.Name = input.Name
	customer.ID = input.ID

	res, totalData, err := s.customer.EntityResults(&customer)
	return res, totalData, err
}

/**
* ==========================================
* Service Result Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityResult(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer schemes.Customer
	customer.ID = input.ID

	res, err := s.customer.EntityResult(&customer)
	return res, err
}

/**
* ==========================================
* Service Delete Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityDelete(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer schemes.Customer
	customer.ID = input.ID

	res, err := s.customer.EntityDelete(&customer)
	return res, err
}

/**
* ==========================================
* Service Update Customer By ID Teritory
*===========================================
 */

func (s *serviceCustomer) EntityUpdate(input *schemes.Customer) (*models.Customer, schemes.SchemeDatabaseError) {
	var customer schemes.Customer
	customer.ID = input.ID
	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address
	customer.Description = input.Description
	customer.MerchantID = input.MerchantID
	customer.OutletID = input.OutletID
	customer.Active = input.Active

	res, err := s.customer.EntityUpdate(&customer)
	return res, err
}
