package task_states

type TaskState int

const (
	Unspecified TaskState = iota
	New                   = iota
	Active                = iota
	Resolved              = iota
)
