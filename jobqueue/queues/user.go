package queues

import (
	"otter/pkg/jobqueue"
)

type user struct {
	SignUp jobqueue.Queue
}

var User user = user{
	SignUp: jobqueue.New(1024),
}
