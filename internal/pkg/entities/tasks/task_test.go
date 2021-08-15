package tasks

import (
	"testing"
	"time"
)

func TestTask_String(t *testing.T) {
	type fields struct {
		UserId      uint64
		TaskId      uint64
		Text        string
		DateCreated time.Time
	}
	parseTimeLight := func(layout, value string) time.Time {
		dateTime, _ := time.Parse(layout, value)
		return dateTime
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String() test",
			fields: fields{
				UserId:      1,
				TaskId:      2,
				Text:        "Задеплоить сервис",
				DateCreated: parseTimeLight(time.RFC3339, "2021-08-15T08:34:00Z"),
			},
			want: "[User:1, task:2] 2021.08.15 08:34:00 - Задеплоить сервис",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := New(
				tt.fields.UserId,
				tt.fields.TaskId,
				tt.fields.Text,
				tt.fields.DateCreated,
			)
			if got := task.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
