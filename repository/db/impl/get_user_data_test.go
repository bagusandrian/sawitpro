package impl

import (
	"context"
	"reflect"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/model"
	"github.com/jmoiron/sqlx"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func Test_repository_GetUserDataByID(t *testing.T) {
	mockDB, mockQuery, _ := sqlmock.Newx()
	type fields struct {
		dbMaster *sqlx.DB
		dbSlave  *sqlx.DB
		cfg      *config.Config
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes model.ResponseGetProfile
		mock    func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "failed get data from DB",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				cfg:      &config.Config{},
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			wantRes: model.ResponseGetProfile{},
			wantErr: true,
		},
		{
			name: "Success get data",
			fields: fields{
				dbMaster: mockDB,
				dbSlave:  mockDB,
				cfg:      &config.Config{},
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			mock: func() {
				mockQuery.ExpectQuery("SELECT phonenumber, fullname FROM users (.+)").
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"phonenumber", "fullname"}).
							AddRow("+6287654321", "testing"))

			},
			wantRes: model.ResponseGetProfile{
				FullName:    "testing",
				PhoneNumber: "+6287654321",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				dbMaster: tt.fields.dbMaster,
				dbSlave:  tt.fields.dbSlave,
				cfg:      tt.fields.cfg,
			}
			if tt.mock != nil {
				tt.mock()
			}
			gotRes, err := r.GetUserDataByID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUserDataByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("repository.GetUserDataByID() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
