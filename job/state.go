package job

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)
