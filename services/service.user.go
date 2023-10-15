package services

import (
	"github.com/nutwreck/admin-pos-service/entities"
	"github.com/nutwreck/admin-pos-service/models"
	"github.com/nutwreck/admin-pos-service/schemes"
)

type serviceUser struct {
	user entities.EntityUser
}

func NewServiceUser(user entities.EntityUser) *serviceUser {
	return &serviceUser{user: user}
}

func (s *serviceUser) EntityAddUser(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var schema schemes.User
	schema.Name = input.Name
	schema.Email = input.Email
	schema.Password = input.Password
	schema.MerchantID = input.MerchantID
	schema.RoleID = input.RoleID

	res, err := s.user.EntityAddUser(&schema)
	return res, err
}

func (s *serviceUser) EntityLogin(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var schema schemes.User
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.user.EntityLogin(&schema)
	return res, err
}

func (s *serviceUser) EntityGetUser(input *schemes.User) (*models.User, schemes.SchemeDatabaseError) {
	var schema schemes.User
	schema.ID = input.ID
	schema.Email = input.Email

	res, err := s.user.EntityGetUser(&schema)
	return res, err
}

func (s *serviceUser) EntityGetRole(input *schemes.Role) (*models.Role, schemes.SchemeDatabaseError) {
	var schema schemes.Role
	schema.ID = input.ID

	res, err := s.user.EntityGetRole(&schema)
	return res, err
}

func (s *serviceUser) EntityGetMerchant(input *schemes.Merchant) (*models.Merchant, schemes.SchemeDatabaseError) {
	var schema schemes.Merchant
	schema.ID = input.ID

	res, err := s.user.EntityGetMerchant(&schema)
	return res, err
}

func (s *serviceUser) EntityGetUserOutlet(input *schemes.UserOutlet) (*[]schemes.GetUserOutlet, schemes.SchemeDatabaseError) {
	var schema schemes.UserOutlet
	schema.UserID = input.UserID

	res, err := s.user.EntityGetUserOutlet(&schema)
	return res, err
}

func (s *serviceUser) EntityUpdate(input *schemes.UpdateUser) (*models.User, schemes.SchemeDatabaseError) {
	var schema schemes.UpdateUser
	schema.Active = input.Active
	schema.Name = input.Name
	schema.OldPassword = input.OldPassword
	schema.NewPassword = input.NewPassword
	schema.DataPassword = input.DataPassword
	schema.RoleID = input.RoleID
	schema.MerchantID = input.MerchantID
	schema.ID = input.ID

	res, err := s.user.EntityUpdate(&schema)
	return res, err
}
