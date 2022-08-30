package main

import (
	"reflect"
	"testing"
)

func Test_findAnagram(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want *map[string][]string
	}{
		{name: "Simple test",
			args: args{[]string{"пятак", "пятка", "тяпка", "тьф", "листок", "слиток", "столик", "фть", "аб", "ба", "a"}},
			want: &map[string][]string{
				"аб":     {"аб", "ба"},
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
				"тьф":    {"тьф", "фть"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagram(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
