package impl

import (
	"github.com/bagusandrian/sawitpro/config"
	httpSawitPro "github.com/bagusandrian/sawitpro/handler/http"
	bcryptImpl "github.com/bagusandrian/sawitpro/helper/bcrypt/impl"
	"github.com/bagusandrian/sawitpro/repository/db"
	dbImpl "github.com/bagusandrian/sawitpro/repository/db/impl"
	"github.com/bagusandrian/sawitpro/resouce"
)

type handler struct {
	config       *config.Config
	dbRepository db.Repository
}

func New(
	// router *mux.Router,
	// Core
	cfg *config.Config,
	resource *resouce.Resources,
) httpSawitPro.Handler {

	// init repository
	bcryptRepository := bcryptImpl.New(cfg)
	dbRepository := dbImpl.New(resource.DBMaster, resource.DBSlave, bcryptRepository, cfg)
	h := &handler{
		config:       cfg,
		dbRepository: dbRepository,
	}
	return h
}
