package handler

type Work struct {
	MaxSteps int
	Name     string
}

// key: WorkId
// value: Work
var Works = map[string]*Work{
	"e454b2cc-7620-4999-9336-cf10a41aebb6": {
		MaxSteps: 5,
		Name:     "action-checker",
	},
}
