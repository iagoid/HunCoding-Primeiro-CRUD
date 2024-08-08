package service

import (
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/rest_err"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

// struct vazia que é utilizada para implementar a interface UserDomainService
type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
