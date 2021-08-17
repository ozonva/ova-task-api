package utils

import (
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var timeFakeTask, _ = time.Parse("2006.01.02", "2021.08.16")

func BuildFakeTasks(keys ...uint64) []tasks.Task {
	var result = make([]tasks.Task, 0, len(keys))
	for _, key := range keys {
		result = append(result, BuildFakeTask(key))
	}
	return result
}

func BuildFakeTask(key uint64) tasks.Task {
	return *tasks.New(key, key, strconv.Itoa(int(key)), timeFakeTask.Add(time.Duration(key)*time.Hour))
}

func TestMapTasks(t *testing.T) {
	type args struct {
		source []tasks.Task
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint64]tasks.Task
		wantErr bool
	}{
		{
			name: "Empty source",
			args: args{
				source: []tasks.Task{},
			},
			want:    map[uint64]tasks.Task{},
			wantErr: false,
		},
		{
			name: "Non empty source",
			args: args{
				source: BuildFakeTasks(1, 2),
			},
			want:    map[uint64]tasks.Task{1: BuildFakeTask(1), 2: BuildFakeTask(2)},
			wantErr: false,
		},
		{
			name: "Non empty source, duplicates",
			args: args{
				source: BuildFakeTasks(1, 2, 3, 2),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Nil source",
			args: args{
				source: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapTasks(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
