package controller

import (
	"net/http"

	"github.com/esteveseverson/crud-go/src/configuration/logger"
	"github.com/esteveseverson/crud-go/src/configuration/validation"
	"github.com/esteveseverson/crud-go/src/model"
	"github.com/esteveseverson/crud-go/src/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := domain.CreateUser(); err != nil {
		c.JSON(http.StatusOK, err.Code)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))
	c.String(http.StatusOK, "")
}
