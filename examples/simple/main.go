package main

import (
	"fmt"
	"github.com/arjunsk/workerpool"
)

func main() {
	wp := workerpool.NewWorkerPool(2)
	wp.Start()

	for i := int32(0); i < 10; i++ {
		val := i // create a copy of i
		err := wp.Submit(func() {
			fmt.Printf(" %v ", val)
		})

		if err != nil {
			panic(err)
		}
	}
	wp.Stop()
}

// Output:
// 0  2  3  4  5  6  7  8  9  1
