package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/logger"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/validation"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/controller/model/request"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// podemos ver uma utilização do domain e do service. Mostrando a importancia de dividir o
	// sitema em camadas. Aqui eu criei um domain, passsando os valores da request
	// mas para realizar algo com esses dados (no caso um create) tenho que criar um service,
	// e o service só aceita como parametro o dominio do usuário
	// assim podemos disponibilizar os metodos que queremos, porém não disponiibilizamos os metodos para
	// alterar os valores passado

	// permite apenas pegar os valores passados
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	// permite apenas acessar os metodos disponiveis (passando um dominio)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertDoomainToResponse(domain))
}
