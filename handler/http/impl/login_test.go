package impl

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/model"
	"github.com/bagusandrian/sawitpro/repository/db"
	"github.com/stretchr/testify/mock"
)

func Test_handler_Login(t *testing.T) {
	form := url.Values{}
	form.Add("phonenumber", "+62123456789")
	form.Add("password", "t3St!ngAj4h")
	validRequest, _ := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	mockDB := db.NewMockRepository(t)
	type fields struct {
		config       *config.Config
		dbRepository db.Repository
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func()
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				config:       &config.Config{},
				dbRepository: mockDB,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: validRequest,
			},
			mock: func() {
				resp := model.ResponseLogin{
					JWTToken: "1234567890",
				}
				mockDB.On("Login", mock.Anything, mock.Anything).Return(resp, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				config:       tt.fields.config,
				dbRepository: tt.fields.dbRepository,
			}
			if tt.mock != nil {
				tt.mock()
			}
			h.Login(tt.args.w, tt.args.r)
		})
	}
}
