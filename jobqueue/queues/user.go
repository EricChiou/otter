package queues

import (
	"github.com/EricChiou/jobqueue"
)

type user struct {
	SignUp jobqueue.Queue
}

var User user = user{
	SignUp: jobqueue.New(1024),
}
