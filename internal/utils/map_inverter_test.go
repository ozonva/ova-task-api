package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInvertMap(t *testing.T) {
	type args struct {
		source map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"empty map", args{source: map[string]string{}}, map[string]string{}},
		{"non empty map", args{source: map[string]string{"key1": "value1", "key2": "value2"}}, map[string]string{"value1": "key1", "value2": "key2"}},
	}
	panicTests := []struct {
		name string
		args args
	}{
		{"map with duplicated values", args{source: map[string]string{"1": "1", "2": "1"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InvertMap(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertMap() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range panicTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panicsf(t, func() {
				InvertMap(tt.args.source)
			}, "Method should panic")
		})
	}
}
