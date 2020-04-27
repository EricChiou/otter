package jobqueue

import (
	"otter/pkg/jobqueue"
)

type codemap struct {
	Add jobQueue
}

var Codemap codemap = codemap{
	Add: jobQueue{jobQueue: jobqueue.New(1024)},
}
