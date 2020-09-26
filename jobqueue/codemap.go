package jobqueue

import (
	"otter/service/apihandler"

	"github.com/EricChiou/jobqueue"
)

type codemap struct {
	add jobqueue.Queue
}

func (u *codemap) NewAddJob(run func() apihandler.ResponseEntity) apihandler.ResponseEntity {
	wait := make(chan apihandler.ResponseEntity)
	u.add.Add(worker{run: run, wait: &wait})

	return <-wait
}

var Codemap codemap = codemap{
	add: jobqueue.New(1024),
}
