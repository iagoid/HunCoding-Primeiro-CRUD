package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/logger"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/validation"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/controller/model/request"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// podemos ver uma utilização do domain e do service. ostrando a impooortancia de dividir o
	// sitea em camadas. Aqui eu criei um domain, passsando os valores da request
	// mas para realizar algo com esses dados (no caso um create) tenho que criar um service,
	// e o service só aceita como paranetro o dominio do usuário
	// assim podemos disponiibilizar os metodos que queeremos, porém não disponiibilizamos os mnetoodos para
	// alterar os valores passado

	// permite aapenas pegar os valores passados
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	// permite apenas acessar os metodos disponiveis (passando um dominio)
	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Error trying to validate user info", zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, domain)
}
