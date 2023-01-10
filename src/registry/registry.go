package registry

import (
	"api_client/interface/controller"
	"net/http"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
	c  *http.Client
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB, c *http.Client) Registry {
	return &registry{db, c}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
	}
}
