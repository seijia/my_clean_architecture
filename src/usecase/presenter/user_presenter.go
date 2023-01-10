// presenter handle the data from use cases to the outer layers defined as abstract
package presenter

import (
	usecasemodel "api_client/usecase/model"
)

type UserPresenter interface {
	ResponseLicense() (*usecasemodel.AuthResult, error)
}
