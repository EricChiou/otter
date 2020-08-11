package jobqueue

import (
	"otter/jobqueue/queues"
	"otter/pkg/jobqueue"
)

func Init() {
	// user job queues
	run(&queues.User.SignUp)

	// codemap job queues
	run(&queues.Codemap.Add)
}

func Wait() {
	// user job queues
	queues.User.SignUp.Wait()

	// codemap job queues
	queues.Codemap.Add.Wait()
}

func run(queue *jobqueue.Queue) {
	queue.SetWorker(func(worker interface{}) {
		if f, ok := worker.(func()); ok {
			f()
		}
	})
	queue.Run()
}
