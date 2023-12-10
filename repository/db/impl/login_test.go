package impl

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/helper/bcrypt"
	"github.com/bagusandrian/sawitpro/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func Test_repository_Login(t *testing.T) {
	mockDB, mockQuery, _ := sqlmock.Newx()
	mockBycrypt := bcrypt.NewMockRepository(t)
	type fields struct {
		dbMaster *sqlx.DB
		dbSlave  *sqlx.DB
		bcrypt   bcrypt.Repository
		cfg      *config.Config
	}
	type args struct {
		ctx context.Context
		req model.RequestLogin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "validation phone number is not register",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg:      &config.Config{},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestLogin{
					PhoneNumber: "+62123456789",
					Password:    "testing",
				},
			},
			mock: func() {
				mockQuery.ExpectQuery("SELECT id, password FROM users (.+)").
					WithArgs("+62123456789").
					WillReturnError(sql.ErrNoRows)

			},
			wantErr: true,
		},
		{
			name: "validation error server",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg:      &config.Config{},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestLogin{
					PhoneNumber: "+62123456789",
					Password:    "testing",
				},
			},
			mock: func() {
				mockQuery.ExpectQuery("SELECT id, password FROM users (.+)").
					WithArgs("+62123456789").
					WillReturnError(errors.New("just error"))

			},
			wantErr: true,
		},
		{
			name: "validation error password not match",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg:      &config.Config{},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestLogin{
					PhoneNumber: "+62123456789",
					Password:    "testing",
				},
			},
			mock: func() {
				mockQuery.ExpectQuery("SELECT id, password FROM users (.+)").
					WithArgs("+62123456789").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "password"}).
							AddRow(1, "testing"))
				mockBycrypt.On("ComparePassword", mock.Anything, mock.Anything).Return(false).Once()
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg: &config.Config{
					Server: config.Server{
						JWTSecretKey: "testing",
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestLogin{
					PhoneNumber: "+62123456789",
					Password:    "testing",
				},
			},
			mock: func() {
				mockQuery.ExpectQuery("SELECT id, password FROM users (.+)").
					WithArgs("+62123456789").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "password"}).
							AddRow(1, "testing"))
				mockBycrypt.On("ComparePassword", mock.Anything, mock.Anything).Return(true).Once()

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				dbMaster: tt.fields.dbMaster,
				dbSlave:  tt.fields.dbSlave,
				bcrypt:   tt.fields.bcrypt,
				cfg:      tt.fields.cfg,
			}
			if tt.mock != nil {
				tt.mock()
			}
			_, err := r.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
