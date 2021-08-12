package utils

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "empty source", args: args{source: []int{}}, want: []int{}},
		{name: "non empty source, nothing to exclude", args: args{source: []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}}, want: []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}},
		{name: "non empty source, exclude part", args: args{source: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, want: []int{1, 4, 6, 8, 9, 10}},
		{name: "non empty source, exclude all", args: args{source: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
