package goric

import (
	"testing"
)

func TestCpfIsValid(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test remove special characters",
			args: args{
				d: "111.111.111-11",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test verifies that the cpf size is 11",
			args: args{
				d: "11111111111",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test verifies that the cpf is in the known invalid format",
			args: args{
				d: "11111111111444",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test check if the cpf digits are valid",
			args: args{
				d: "432.484.831-90",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test cpf is valid",
			args: args{
				d: "432.484.831-99",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CpfIsValid(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("CpfIsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CpfIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateCpfFirstDigit(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test generate first cpf digit ok",
			args: args{
				d: "44504356724",
			},
			want: 2,
		},
		{
			name: "test generate first cpf digit not ok",
			args: args{
				d: "82607236207",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateCpfFirstDigit(tt.args.d); got != tt.want {
				t.Errorf("GenerateCpfFirstDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateCpfSecondDigit(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test generate second cpf digit ok",
			args: args{
				d: "44504356724",
			},
			want: 4,
		},
		{
			name: "test generate second cpf not ok",
			args: args{
				d: "56043299160",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateCpfSecondDigit(tt.args.d); got != tt.want {
				t.Errorf("GenerateCpfSecondDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
