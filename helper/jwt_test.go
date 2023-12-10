package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	type args struct {
		userID    int64
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success generate",
			args: args{
				userID:    1,
				secretKey: "testing",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateJWT(tt.args.userID, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestExtractTokenFromHeader(t *testing.T) {

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "success extract token",
			args: args{
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/testing", nil)
					req.Header.Set("Authorization", "Bearer 1234567890")
					return req
				}(),
			},
			want: "1234567890",
		},
		{
			name: "no auth",
			args: args{
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/testing", nil)
					req.Header.Set("failed", "")
					return req
				}(),
			},
		},
		{
			name: "no Bearer without key",
			args: args{
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/testing", nil)
					req.Header.Set("Authorization", "Bearer")
					return req
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractTokenFromHeader(tt.args.r); got != tt.want {
				t.Errorf("ExtractTokenFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
