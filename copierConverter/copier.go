package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type UserDomain struct {
	ID             string        `copier:"Huncoding"`
	Name           string        `copier:"Name"`
	Age            int8          `copier:"Age"`
	LastName       string        `copier:"LastName"`
	Email          string        `copier:"Email"`
	DocumentNumber string        `copier:"DocumentNumber"`
	Address        AddressDomain `copier:"Address"`
}

type AddressDomain struct {
	Street     string `copier:"Street"`
	Number     int    `copier:"Number"`
	CEP        string `copier:"CEP"`
	State      string `copier:"State"`
	City       string `copier:"City"`
	Complement string `copier:"Complement"`
}

type UserRequest struct {
	ID             string         `copier:"Huncoding"`
	Name           string         `copier:"Name"`
	Age            int8           `copier:"Age"`
	LastName       string         `copier:"LastName"`
	Email          string         `copier:"Email"`
	DocumentNumber string         `copier:"DocumentNumber"`
	Address        AddressRequest `copier:"Address"`
}

type AddressRequest struct {
	Street     string `copier:"Street"`
	Number     int    `copier:"Number"`
	CEP        string `copier:"CEP"`
	State      string `copier:"State"`
	City       string `copier:"City"`
	Complement string `copier:"Complement"`
}

func convertRequestToDomain(
	request UserRequest) *UserDomain {
	return &UserDomain{
		ID:             request.ID,
		Name:           request.Name,
		Age:            request.Age,
		LastName:       request.LastName,
		Email:          request.Email,
		DocumentNumber: request.DocumentNumber,
		Address: AddressDomain{
			Street:     request.Address.Street,
			Number:     request.Address.Number,
			CEP:        request.Address.CEP,
			State:      request.Address.State,
			City:       request.Address.City,
			Complement: request.Address.Complement,
		},
	}
}

func main() {
	userRequest := UserRequest{
		ID:             "123",
		Name:           "HunCoding",
		Age:            22,
		LastName:       "Test",
		Email:          "test@gmail.com",
		DocumentNumber: "12345",
		Address: AddressRequest{
			Street:     "test rua",
			Number:     123,
			CEP:        "12345",
			State:      "SP",
			City:       "SP",
			Complement: "Test",
		},
	}

	domain := convertRequestToDomain(userRequest)
	fmt.Printf("Converted object: %#v \n", domain)

	userDomainCopier := UserDomain{}
	err := copier.Copy(&userDomainCopier, &userRequest)
	if err != nil {
		return
	}
	fmt.Printf("Converted object with copier: %#v \n", userDomainCopier)
}
