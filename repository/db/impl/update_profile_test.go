package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/helper/bcrypt"
	"github.com/bagusandrian/sawitpro/model"
	"github.com/jmoiron/sqlx"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func Test_repository_UpdateProfile(t *testing.T) {
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
		req model.RequestUpdateProfile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes model.ResponseUpdateProfile
		mock    func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				bcrypt:   mockBycrypt,
				cfg:      &config.Config{},
			},
			args: args{
				ctx: context.Background(),
				req: model.RequestUpdateProfile{
					FullName:    "testing",
					PhoneNumber: "+1234567890",
				},
			},
			wantRes: model.ResponseUpdateProfile{
				FullName:       "testing",
				PhoneNumber:    "+1234567890",
				SuccessMessage: "success",
			},
			mock: func() {
				mockQuery.ExpectBegin()
				mockQuery.ExpectExec("UPDATE users set fullname").WillReturnResult(sqlmock.NewErrorResult(nil))
				mockQuery.ExpectExec("UPDATE users set phonenumber").WillReturnResult(sqlmock.NewErrorResult(nil))
				mockQuery.ExpectCommit()
				mockQuery.ExpectRollback()
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
			gotRes, err := r.UpdateProfile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("repository.UpdateProfile() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
