package impl

import (
	"testing"

	"github.com/bagusandrian/sawitpro/config"
	"github.com/bagusandrian/sawitpro/resouce"
)

func TestNew(t *testing.T) {
	// mockBcrypt := bcrypt.NewMockRepository(t)
	type args struct {
		cfg      *config.Config
		resource *resouce.Resources
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				cfg:      &config.Config{},
				resource: &resouce.Resources{},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cfg, tt.args.resource)
		})
	}
}
