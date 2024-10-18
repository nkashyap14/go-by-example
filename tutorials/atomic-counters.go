package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//primary mechanism for managing state in Go is communication through channels


//other mechanism exists. in this case sync/atomic has atomic counters

func main() {

	//atomic integer to represent our counter
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			
			wg.Done()
		}()
	}

	wg.Wait()

	//get 50000 because this is an atomic integer. if we had used a non atomic integer we would have got a different number changing between runs because the goroutines would interfere with each other
	fmt.Println("ops:", ops.Load())
}