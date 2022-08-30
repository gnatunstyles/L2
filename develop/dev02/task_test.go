package main

import "testing"

func Test_unpack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test_simple",
			args:    args{"a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "test_rus",
			args:    args{"а4ку2и5ч"},
			want:    "аааакууииииич",
			wantErr: false,
		},
		{
			name:    "test_error",
			args:    args{"45"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpack(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
