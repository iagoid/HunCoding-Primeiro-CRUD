package service

import (
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/logger"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/rest_err"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}
