package impl

import (
	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/repository/db"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type repository struct {
	dbMaster *sqlx.DB
	dbSlave  *sqlx.DB
	cfg      *config.Config
}

func New(dbMaster, dbSlave *sqlx.DB, cfg *config.Config) db.Repository {
	return &repository{
		dbSlave:  dbSlave,
		dbMaster: dbMaster,
		cfg:      cfg,
	}
}
