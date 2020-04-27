package jobqueue

import (
	"otter/pkg/jobqueue"
)

type user struct {
	SignUp jobQueue
}

var User user = user{
	SignUp: jobQueue{jobQueue: jobqueue.New(1024)},
}
