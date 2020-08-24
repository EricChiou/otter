package queues

import (
	"github.com/EricChiou/jobqueue"
)

type codemap struct {
	Add jobqueue.Queue
}

var Codemap codemap = codemap{
	Add: jobqueue.New(1024),
}
