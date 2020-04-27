package jobqueue

import "otter/pkg/jobqueue"

type jobQueue struct {
	jobQueue jobqueue.JobQueue
}

func (j *jobQueue) Init() {
	j.jobQueue.SetWorker(func(worker interface{}) {
		if f, ok := worker.(func()); ok {
			f()
		}
	})
	j.jobQueue.Run()
}

func (j *jobQueue) Wait() {
	j.jobQueue.Wait()
}

func (j *jobQueue) Add(f func()) {
	j.jobQueue.Add(f)
}

func Init() {
	User.SignUp.Init()
	Codemap.Add.Init()
}

func Wait() {
	User.SignUp.Wait()
	Codemap.Add.Wait()
}
