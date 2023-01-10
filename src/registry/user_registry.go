package registry

import (
	"api_client/interface/controller"
	ip "api_client/interface/presenter"
	ir "api_client/interface/repository"
	"api_client/usecase/interactor"
	up "api_client/usecase/presenter"
	ur "api_client/usecase/repository"

	"github.com/go-redis/redis/v8"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserPresenter(), r.NewRedisRepository())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

func (r *registry) NewRedisRepository() ur.RedisRepository {
	var rc *redis.Client
	return ir.NewRedisRepository(rc)
}
