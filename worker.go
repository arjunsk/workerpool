package workerpool

import (
	"sync"
)

// Worker is responsible for executing the work. There could be many workers in a worker pool.
type Worker struct {
	wg               *sync.WaitGroup
	workerCh         chan Work
	freeWorkerPoolCh chan chan Work
	stopCh           chan bool
}

func NewWorker(freeWorkerPoolCh chan chan Work, wg *sync.WaitGroup) *Worker {
	return &Worker{
		wg:               wg,
		workerCh:         make(chan Work),
		freeWorkerPoolCh: freeWorkerPoolCh,
		stopCh:           make(chan bool),
	}
}

func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		for {
			// Add workerCh back to freeWorkerPoolCh when it is not doing any work (ie completed the old work).
			w.freeWorkerPoolCh <- w.workerCh

			select {
			case work := <-w.workerCh:
				work()
			case <-w.stopCh:
				w.wg.Done()
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.stopCh <- true
}
