package main

//to wait for multiple go routines to finish we can use a wait group

import (
	"fmt"
	"sync"
	"time"
)

//function we'll run in every goroutine

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	
	//when we pass a weight group into functions should be done via pointer
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		//launching several routines and incremeenting the ounter by 1
		wg.Add(1)

		//wrapping the worker call in a closure that makes sure to tell the waitgroup the worker is done
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	//block until the waitgroup counter goes to 0
	wg.Wait()

	//no straightforward way to propagate errors from workers
}