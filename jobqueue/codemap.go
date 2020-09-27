package jobqueue

import (
	"github.com/EricChiou/jobqueue"
)

type codemap struct {
	add jobqueue.Queue
}

func (u *codemap) NewAddJob(run func() interface{}) interface{} {
	wait := make(chan interface{})
	u.add.Add(worker{run: run, wait: &wait})

	return <-wait
}

var Codemap codemap = codemap{
	add: jobqueue.New(1024),
}
