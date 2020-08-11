package queues

import (
	"otter/pkg/jobqueue"
)

type codemap struct {
	Add jobqueue.Queue
}

var Codemap codemap = codemap{
	Add: jobqueue.New(1024),
}
