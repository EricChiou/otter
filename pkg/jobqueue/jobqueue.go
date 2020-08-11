package jobqueue

import (
	"errors"
	"sync"
)

type Queue struct {
	channel chan interface{}
	wg      sync.WaitGroup
	worker  func(param interface{})
	blocked bool
	running bool
	closed  bool
	run     func(jobQueue *Queue)
}

// Set jobqueue worker function
func (queue *Queue) SetWorker(worker func(param interface{})) {
	queue.worker = worker
}

// Set if jobqueue can be blocked. Default is false
func (queue *Queue) SetBlocked(block bool) {
	queue.blocked = block
}

// Run jobqueue
func (queue *Queue) Run() error {
	if queue.closed {
		queue.running = false
		return errors.New("jobqueue is closed.")
	}

	if queue.worker == nil {
		return errors.New("jobqueue is not set worker function yet.")
	}

	queue.running = true
	if queue.run == nil {
		queue.run = run
		go queue.run(queue)
	}
	return nil
}

// Add a new job to jobqueue
func (queue *Queue) Add(param interface{}) error {
	if queue.closed {
		return errors.New("jobqueue was closed.")
	}

	if queue.running {
		if queue.blocked {
			queue.channel <- param
			queue.wg.Add(1)

		} else {
			select {
			case queue.channel <- param:
				queue.wg.Add(1)
				return nil
			default:
				return errors.New("jobqueue is full.")
			}
		}
	}

	return errors.New("jobqueue is not start yet.")
}

// Wait for jobqueue finished
func (queue *Queue) Wait() {
	queue.wg.Wait()
}

// Start jobqueue
func (queue *Queue) Start() error {
	if queue.run != nil {
		queue.running = true
	} else {
		return errors.New("jobqueue is not running.")
	}

	return nil
}

// Stop jobqueue accept a new job
func (queue *Queue) Stop() {
	queue.running = false
}

// Close jobqueue
func (queue *Queue) Close() {
	queue.running = false
	queue.closed = true
	close(queue.channel)
}

// New a job queue
func New(size int) Queue {
	return Queue{
		channel: make(chan interface{}, size),
		wg:      sync.WaitGroup{},
		worker:  nil,
		blocked: false,
		running: false,
		closed:  false,
		run:     nil,
	}
}

func run(queue *Queue) {
	for job := range queue.channel {
		queue.worker(job)
		queue.wg.Done()
	}
}
