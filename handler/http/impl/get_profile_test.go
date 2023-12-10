package impl

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/repository/db"
)

func Test_handler_GetProfile(t *testing.T) {
	validRequest := httptest.NewRequest("GET", "/get_my_profile", bytes.NewBufferString(``))
	validRequest.Header.Add("Authorization", "Bearer 1234567890")
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
			name: "failed jwt token",
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
			h.GetProfile(tt.args.w, tt.args.r)
		})
	}
}
