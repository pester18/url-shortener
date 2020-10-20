package registry

import (
	"github.com/pester18/url-shortener/interface/controller"
	ir "github.com/pester18/url-shortener/interface/repository"
	"github.com/pester18/url-shortener/usecase/interactor"
	ur "github.com/pester18/url-shortener/usecase/repository"
	mgo "gopkg.in/mgo.v2"
)

type registry struct {
	db *mgo.Database
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *mgo.Database) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.NewAppController(r.NewInteractor())
}

func (r *registry) NewInteractor() interactor.Interactor {
	return interactor.NewInteractor(r.NewRepository())
}

func (r *registry) NewRepository() ur.Repository {
	return ir.NewRepository(r.db)
}
