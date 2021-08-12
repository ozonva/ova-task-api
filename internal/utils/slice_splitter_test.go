package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	type args struct {
		source []string
		batch  int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"empty source, batch > 0", args{source: []string{}, batch: 5}, [][]string{}},
		{"empty source, batch > 0", args{source: nil, batch: 5}, [][]string{}},
		{"non empty source, batch = 1", args{source: []string{"1", "2", "3"}, batch: 1}, [][]string{{"1"}, {"2"}, {"3"}}},
		{"non empty source, batch < len(source)/2", args{source: []string{"1", "2", "3", "4", "5"}, batch: 2}, [][]string{{"1", "2"}, {"3", "4"}, {"5"}}},
		{"non empty source, batch = len(source)/2", args{source: []string{"1", "2", "3", "4"}, batch: 2}, [][]string{{"1", "2"}, {"3", "4"}}},
		{"non empty source, batch > len(source)/2", args{source: []string{"1", "2", "3", "4", "5"}, batch: 3}, [][]string{{"1", "2", "3"}, {"4", "5"}}},
		{"non empty source, batch = len(source)", args{source: []string{"1", "2", "3"}, batch: 3}, [][]string{{"1", "2", "3"}}},
		{"non empty source, batch > len(source)", args{source: []string{"1", "2", "3"}, batch: 5}, [][]string{{"1", "2", "3"}}},
	}
	panicTests := []struct {
		name string
		args args
	}{
		{"empty source, batch = 0", args{source: []string{}, batch: 0}},
		{"non empty source, batch = 0", args{source: []string{}, batch: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slices := SplitSlice(tt.args.source, tt.args.batch)
			if cap(slices) != cap(tt.want) {
				t.Errorf("cap(actual) = %v, cap(expected) %v", cap(slices), cap(tt.want))
			}
			if !reflect.DeepEqual(slices, tt.want) {
				t.Errorf("actual = %v, expected %v", slices, tt.want)
			}
		})
	}
	for _, tt := range panicTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panicsf(t, func() {
				SplitSlice(tt.args.source, tt.args.batch)
			}, "Method should panic")
		})
	}
}
