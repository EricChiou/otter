package jobqueue

import (
	"otter/service/apihandler"

	"github.com/EricChiou/jobqueue"
)

type user struct {
	signUp jobqueue.Queue
}

func (u *user) NewSignUpJob(run func() apihandler.ResponseEntity) apihandler.ResponseEntity {
	wait := make(chan apihandler.ResponseEntity)
	u.signUp.Add(worker{run: run, wait: &wait})

	return <-wait
}

var User user = user{
	signUp: jobqueue.New(1024),
}
