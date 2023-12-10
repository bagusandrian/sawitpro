package impl

import (
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/helper/bcrypt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestNew(t *testing.T) {
	dbMock, _, _ := sqlmock.NewxWithDSN("testing")
	type args struct {
		dbMaster *sqlx.DB
		dbSlave  *sqlx.DB
		bcrypt   bcrypt.Repository
		cfg      *config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				dbMaster: dbMock,
				dbSlave:  dbMock,
				bcrypt:   bcrypt.NewMockRepository(t),
				cfg:      &config.Config{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.dbMaster, tt.args.dbSlave, tt.args.bcrypt, tt.args.cfg)
		})
	}
}
