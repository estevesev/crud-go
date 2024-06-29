package view

import (
	"github.com/esteveseverson/crud-go/src/model"
	"github.com/esteveseverson/crud-go/src/model/response"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
