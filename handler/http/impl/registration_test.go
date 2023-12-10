package impl

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/repository/db"
)

func Test_handler_Registration(t *testing.T) {
	form := url.Values{}
	form.Add("fullname", "testing")
	form.Add("phonenumber", "+62123456789")
	form.Add("password", "t3St!ngAj4h")
	validRequest, _ := http.NewRequest("POST", "/registration", strings.NewReader(form.Encode()))
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
	}{
		// TODO: Add test cases.
		{
			name: "failed validation data",
			fields: fields{
				config:       &config.Config{},
				dbRepository: mockDB,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: validRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				config:       tt.fields.config,
				dbRepository: tt.fields.dbRepository,
			}
			h.Registration(tt.args.w, tt.args.r)
		})
	}
}
