package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type CounterType int

const (
	CreateTaskCounter CounterType = iota
	UpdateTaskCounter
	RemoveTaskCounter
)

var (
	counters = map[CounterType]prometheus.Counter{
		CreateTaskCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Name: "createTaskCounter",
			},
		),
		UpdateTaskCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Name: "updateTaskCounter",
			},
		),
		RemoveTaskCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Name: "removeTaskCounter",
			},
		),
	}
)

func Inc(counterType CounterType) {
	counters[counterType].Inc()
}
