package helper

import (
	"testing"
)

func TestContainsUppercase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "contains upper case postive",
			args: args{
				s: "TestinG",
			},
			want: true,
		},
		{
			name: "contains upper case negative",
			args: args{
				s: "testing",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUppercase(tt.args.s); got != tt.want {
				t.Errorf("ContainsUppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "contains digit postive",
			args: args{
				s: "Testi4G",
			},
			want: true,
		},
		{
			name: "contains digit case negative",
			args: args{
				s: "testing",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDigit(tt.args.s); got != tt.want {
				t.Errorf("ContainsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsSpecialChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "contains special char postive",
			args: args{
				s: "Testi@G",
			},
			want: true,
		},
		{
			name: "contains special char negative",
			args: args{
				s: "testing",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsSpecialChar(tt.args.s); got != tt.want {
				t.Errorf("ContainsSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
