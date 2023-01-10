// interactor handle data from the outer layer defined as abstract
package interactor

import (
	usecasemodel "api_client/usecase/model"
	"api_client/usecase/presenter"
	"api_client/usecase/repository"
)

type userInteractor struct {
	UserPresenter   presenter.UserPresenter
	RedisRepository repository.RedisRepository
}

type UserInteractor interface {
	GetLicense(*usecasemodel.LicenseRequest) (*usecasemodel.AuthResult, error)
}

func NewUserInteractor(p presenter.UserPresenter, r repository.RedisRepository) UserInteractor {
	return &userInteractor{p, r}
}

func (uc *userInteractor) GetLicense(*usecasemodel.LicenseRequest) (*usecasemodel.AuthResult, error) {

	return uc.UserPresenter.ResponseLicense()
}
