package config

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(context.Background())
			_, _ = got, err
		})
	}
}
