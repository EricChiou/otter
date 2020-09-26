package jobqueue

import (
	"otter/service/apihandler"

	"github.com/EricChiou/jobqueue"
)

type worker struct {
	run  func() apihandler.ResponseEntity
	wait *chan apihandler.ResponseEntity
}

func Init() {
	// user job queues
	run(&User.signUp)

	// codemap job queues
	run(&Codemap.add)
}

func Wait() {
	// user job queues
	User.signUp.Wait()

	// codemap job queues
	Codemap.add.Wait()
}

func run(queue *jobqueue.Queue) {
	queue.SetWorker(func(w interface{}) {
		if w, ok := w.(worker); ok {
			*w.wait <- w.run()
		} else {
			*w.wait <- apihandler.ResponseEntity{}
		}
	})
	queue.Run()
}
