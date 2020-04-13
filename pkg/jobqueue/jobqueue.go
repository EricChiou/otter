package jobqueue

import (
	"log"
	"sync"
)

type JobQueue struct {
	queue   chan interface{}
	wg      sync.WaitGroup
	worker  func(param interface{})
	blocked bool
	running bool
	closed  bool
	run     func(jobQueue *JobQueue)
}

// Set jobqueue worker function
func (jobQueue *JobQueue) SetWorker(worker func(param interface{})) {
	jobQueue.worker = worker
}

// Set if jobqueue can be blocked. Default is false
func (jobQueue *JobQueue) SetBlocked(block bool) {
	jobQueue.blocked = block
}

// Run jobqueue
func (jobQueue *JobQueue) Run() bool {
	if jobQueue.closed {
		jobQueue.running = false
		log.Println("jobqueue is closed.")
		return false
	}

	if jobQueue.worker == nil {
		log.Println("jobqueue is not set worker function yet.")
		return false
	}

	jobQueue.running = true
	if jobQueue.run == nil {
		jobQueue.run = run
		go jobQueue.run(jobQueue)
	}
	return true
}

// Add a new job to jobqueue
func (jobQueue *JobQueue) Add(param interface{}) bool {
	if jobQueue.closed {
		log.Println("jobqueue is closed.")
		return false
	}

	if jobQueue.running {
		if jobQueue.blocked {
			jobQueue.queue <- param
			jobQueue.wg.Add(1)

		} else {
			select {
			case jobQueue.queue <- param:
				jobQueue.wg.Add(1)
				return true
			default:
				log.Println("jobqueue is full.")
				return false
			}
		}
	}

	log.Println("jobqueue is not start.")
	return false
}

// Wait for jobqueue finished
func (jobQueue *JobQueue) Wait() {
	jobQueue.wg.Wait()
}

// Start jobqueue
func (jobQueue *JobQueue) Start() {
	if jobQueue.run != nil {
		jobQueue.running = true
	} else {
		log.Println("jobqueue is not running.")
	}
}

// Stop jobqueue accept a new job
func (jobQueue *JobQueue) Stop() {
	jobQueue.running = false
}

// Close jobqueue
func (jobQueue *JobQueue) Close() {
	jobQueue.running = false
	jobQueue.closed = true
	close(jobQueue.queue)
}

// New a job queue
func New(size int) JobQueue {
	return JobQueue{
		queue:   make(chan interface{}, size),
		worker:  nil,
		blocked: false,
		running: false,
		closed:  false,
	}
}

func run(jobQueue *JobQueue) {
	for job := range jobQueue.queue {
		jobQueue.worker(job)
		jobQueue.wg.Done()
	}
}
