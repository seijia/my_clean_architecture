// specific implement of output port in Use Cases, before it passing to view
package presenter

import (
	usecasemodel "api_client/usecase/model"
	"api_client/usecase/presenter"
	"time"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (cp *userPresenter) ResponseLicense() (*usecasemodel.AuthResult, error) {

	timestamp := time.Now().Format("2006/01/02 15:4:5")
	timeLimit := time.Now().AddDate(0, 0, 1).Format("2006/01/02 15:4:5")
	u := &usecasemodel.AuthResult{
		Status: usecasemodel.AuthStatus{
			Error: "0",
		},
		LicenseInfo: usecasemodel.AuthLicenseInfo{
			ConvertID: "b8hpcg8hv3amvi9dol0g",
			StartTime: timestamp,
			EndTime:   timeLimit,
		},
	}
	return u, nil
}
