package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

// UserDomainInterface interface utilizada como parametro no service
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword()
}

func NewUserDomain(email string, password string, name string, age int8) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

// existe a necessidade de criar getters pois como estou utilizando a interface
// para passar as funções (dividindo o sistema em camadas)
// eu preciso de funções que permitam pegar os valores através da interface

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
