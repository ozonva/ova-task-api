package utils

import (
	"github.com/stretchr/testify/assert"
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
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
		{"empty source, butchSize > 0", args{source: []string{}, batch: 5}, [][]string{}},
		{"empty source, butchSize > 0", args{source: nil, batch: 5}, [][]string{}},
		{"non empty source, butchSize = 1", args{source: []string{"1", "2", "3"}, batch: 1}, [][]string{{"1"}, {"2"}, {"3"}}},
		{"non empty source, butchSize < len(source)/2", args{source: []string{"1", "2", "3", "4", "5"}, batch: 2}, [][]string{{"1", "2"}, {"3", "4"}, {"5"}}},
		{"non empty source, butchSize = len(source)/2", args{source: []string{"1", "2", "3", "4"}, batch: 2}, [][]string{{"1", "2"}, {"3", "4"}}},
		{"non empty source, butchSize > len(source)/2", args{source: []string{"1", "2", "3", "4", "5"}, batch: 3}, [][]string{{"1", "2", "3"}, {"4", "5"}}},
		{"non empty source, butchSize = len(source)", args{source: []string{"1", "2", "3"}, batch: 3}, [][]string{{"1", "2", "3"}}},
		{"non empty source, butchSize > len(source)", args{source: []string{"1", "2", "3"}, batch: 5}, [][]string{{"1", "2", "3"}}},
	}
	panicTests := []struct {
		name string
		args args
	}{
		{"empty source, butchSize = 0", args{source: []string{}, batch: 0}},
		{"non empty source, butchSize = 0", args{source: []string{}, batch: 0}},
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

func TestSplitToBulks(t *testing.T) {
	type args struct {
		source    []tasks.Task
		butchSize int
	}
	tests := []struct {
		name string
		args args
		want [][]tasks.Task
	}{
		{"empty source, butchSize > 0", args{source: []tasks.Task{}, butchSize: 5}, [][]tasks.Task{}},
		{"empty source, butchSize > 0", args{source: nil, butchSize: 5}, [][]tasks.Task{}},
		{"non empty source, butchSize = 1", args{source: BuildFakeTasks(1, 2, 3), butchSize: 1}, [][]tasks.Task{BuildFakeTasks(1), BuildFakeTasks(2), BuildFakeTasks(3)}},
		{"non empty source, butchSize < len(source)/2", args{source: BuildFakeTasks(1, 2, 3, 4, 5), butchSize: 2}, [][]tasks.Task{BuildFakeTasks(1, 2), BuildFakeTasks(3, 4), BuildFakeTasks(5)}},
		{"non empty source, butchSize = len(source)/2", args{source: BuildFakeTasks(1, 2, 3, 4), butchSize: 2}, [][]tasks.Task{BuildFakeTasks(1, 2), BuildFakeTasks(3, 4)}},
		{"non empty source, butchSize > len(source)/2", args{source: BuildFakeTasks(1, 2, 3, 4, 5), butchSize: 3}, [][]tasks.Task{BuildFakeTasks(1, 2, 3), BuildFakeTasks(4, 5)}},
		{"non empty source, butchSize = len(source)", args{source: BuildFakeTasks(1, 2, 3), butchSize: 3}, [][]tasks.Task{BuildFakeTasks(1, 2, 3)}},
		{"non empty source, butchSize > len(source)", args{source: BuildFakeTasks(1, 2, 3), butchSize: 5}, [][]tasks.Task{BuildFakeTasks(1, 2, 3)}},
	}
	panicTests := []struct {
		name string
		args args
	}{
		{"empty source, butchSize = 0", args{source: []tasks.Task{}, butchSize: 0}},
		{"non empty source, butchSize = 0", args{source: []tasks.Task{}, butchSize: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slices := SplitTasksSlice(tt.args.source, tt.args.butchSize)
			if cap(slices) != cap(tt.want) {
				t.Errorf("cap(actual) = %v, cap(expected) %v", cap(slices), cap(tt.want))
			}
			if !reflect.DeepEqual(slices, tt.want) {
				t.Errorf("actual = %v, \r\nexpected %v", slices, tt.want)
			}
		})
	}
	for _, tt := range panicTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panicsf(t, func() {
				SplitTasksSlice(tt.args.source, tt.args.butchSize)
			}, "Method should panic")
		})
	}
}
