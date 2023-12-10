package resouce

import (
	"fmt"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/jmoiron/sqlx"
)

type Resources struct {
	DBMaster *sqlx.DB
	DBSlave  *sqlx.DB
}

func InitResource(conf *config.Config) (*Resources, error) {
	// initDB slave
	DBSlaveInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Slave.Host, conf.Database.Slave.Port, conf.Database.Slave.User, conf.Database.Slave.Password, conf.Database.Slave.DDname)
	DBSlave, err := sqlx.Connect("postgres", DBSlaveInfo)
	if err != nil {
		return nil, err
	}
	err = DBSlave.Ping()
	if err != nil {
		return nil, err
	}
	// init db master
	DBSMasterInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Master.Host, conf.Database.Master.Port, conf.Database.Master.User, conf.Database.Master.Password, conf.Database.Master.DDname)
	DBMaster, err := sqlx.Connect("postgres", DBSMasterInfo)
	if err != nil {
		return nil, err
	}
	err = DBMaster.Ping()
	if err != nil {
		return nil, err
	}
	return &Resources{
		DBMaster: DBMaster,
		DBSlave:  DBSlave,
	}, nil
}
