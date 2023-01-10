// specific implement of input port in ther Use Cases, before saving it in database

package controller

import (
	"api_client/usecase/interactor"
	usecasemodel "api_client/usecase/model"
	"api_client/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetLicense(c echo.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (cc *userController) GetLicense(c echo.Context) error {
	params := new(usecasemodel.LicenseRequest)

	if err := c.Bind(params); !errors.Is(err, nil) {
		return utils.NewBadRequestError(err)
	}
	u, err := cc.userInteractor.GetLicense(params)
	if err != nil {
		return c.XML(http.StatusBadRequest, u)
	}
	return c.XML(http.StatusOK, u)
}
