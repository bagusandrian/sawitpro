package impl

import (
	"testing"

	"github.com/bagusandrian/sawitpro/config"
)

func Test_repository_GeneratePassword(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				cfg: &config.Config{},
			},
			args: args{
				password: "1234567890",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				cfg: tt.fields.cfg,
			}
			if got := r.GeneratePassword(tt.args.password); len(got) != 0 {
				t.Errorf("repository.GeneratePassword() error test")
			}
		})
	}
}

func Test_repository_ComparePassword(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		passwordHash string
		passwordReq  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "failed",
			fields: fields{
				cfg: &config.Config{},
			},
			args: args{
				passwordHash: "1234567890",
				passwordReq:  "1234567890",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				cfg: tt.fields.cfg,
			}
			if got := r.ComparePassword(tt.args.passwordHash, tt.args.passwordReq); got != tt.want {
				t.Errorf("repository.ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
