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

func (s *serviceUser) EntityRegister(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError) {
	var schema schemes.SchemeUser
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Role = input.Role

	res, err := s.user.EntityRegister(&schema)
	return res, err
}

func (s *serviceUser) EntityLogin(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError) {
	var schema schemes.SchemeUser
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.user.EntityLogin(&schema)
	return res, err
}

func (s *serviceUser) EntityGetUser(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError) {
	var schema schemes.SchemeUser
	schema.ID = input.ID
	schema.Email = input.Email

	res, err := s.user.EntityGetUser(&schema)
	return res, err
}

func (s *serviceUser) EntityUpdate(input *schemes.SchemeUpdateUser) (*models.ModelUser, schemes.SchemeDatabaseError) {
	var schema schemes.SchemeUpdateUser
	schema.Active = input.Active
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.OldPassword = input.OldPassword
	schema.NewPassword = input.NewPassword
	schema.DataPassword = input.DataPassword
	schema.Role = input.Role
	schema.ID = input.ID

	res, err := s.user.EntityUpdate(&schema)
	return res, err
}
