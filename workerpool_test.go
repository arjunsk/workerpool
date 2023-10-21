package workerpool

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewWorkerPoolStartStop(t *testing.T) {
	wp := NewWorkerPool(runtime.NumCPU())

	wp.Start()
	assert.Equal(t, false, wp.stopped.Load())

	wp.Stop()
	assert.Equal(t, true, wp.stopped.Load())
}

func TestWorkerPool_Submit(t *testing.T) {
	wp := NewWorkerPool(2)
	wp.Start()

	var expected int32 = 100
	var actual atomic.Int32

	for i := int32(0); i < expected; i++ {
		err := wp.Submit(func() {
			// adding an extra delay so that worker pool stop() is called before the jobs are finished.
			time.Sleep(10 * time.Millisecond)
			actual.Add(1)
		})

		assert.Equal(t, nil, err)
	}

	wp.Stop()
	assert.Equal(t, expected, actual.Load())
}
