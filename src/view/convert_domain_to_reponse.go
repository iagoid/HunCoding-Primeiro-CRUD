package view

import (
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/controller/model/response"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model"
)

func ConvertDoomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
