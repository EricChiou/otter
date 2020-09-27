package jobqueue

import (
	"github.com/EricChiou/jobqueue"
)

type user struct {
	signUp jobqueue.Queue
}

func (u *user) NewSignUpJob(run func() interface{}) interface{} {
	wait := make(chan interface{})
	u.signUp.Add(worker{run: run, wait: &wait})

	return <-wait
}

var User user = user{
	signUp: jobqueue.New(1024),
}
