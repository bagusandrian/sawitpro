package impl

import (
	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/helper/bcrypt"

	_ "github.com/lib/pq"
)

type repository struct {
	cfg *config.Config
}

func New(cfg *config.Config) bcrypt.Repository {
	return &repository{
		cfg: cfg,
	}
}
