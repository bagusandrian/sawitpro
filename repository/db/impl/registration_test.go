package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/helper/bcrypt"
	"github.com/bagusandrian/sawitpro/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func Test_repository_Registration(t *testing.T) {
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
		req model.RequestRegistration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes model.ResponseRegristration
		mock    func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg:      &config.Config{},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestRegistration{
					FullName:    "testing",
					PhoneNumber: "+6281234567890",
					Password:    "1234567890",
				},
			},
			mock: func() {
				mockQuery.ExpectBegin()
				mockQuery.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
				mockQuery.ExpectCommit()
				mockBycrypt.On("GeneratePassword", mock.Anything).Return("testing").Once()
			},
			wantRes: model.ResponseRegristration{
				FullName:    "testing",
				PhoneNumber: "+6281234567890",
			},
			wantErr: false,
		},
		{
			name: "duplicate phone number",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg:      &config.Config{},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestRegistration{
					FullName:    "testing",
					PhoneNumber: "+6281234567890",
					Password:    "1234567890",
				},
			},
			mock: func() {
				mockQuery.ExpectBegin()
				mockQuery.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(0, 0))
				mockQuery.ExpectCommit()
				mockQuery.ExpectRollback()
				mockBycrypt.On("GeneratePassword", mock.Anything).Return("testing").Once()
			},
			wantRes: model.ResponseRegristration{},
			wantErr: true,
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
			tt.mock()
			gotRes, err := r.Registration(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Registration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("repository.Registration() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
