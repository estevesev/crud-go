package service

import (
	"github.com/esteveseverson/crud-go/src/configuration/rest_err"
	"github.com/esteveseverson/crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
