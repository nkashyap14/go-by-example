package main

import "fmt"

//closing channels indicates no more values will be sent on it. can be useful to communicate completion to the channel's receivers

func main() {

	//communicates work to be done from the main() go routine to a worker goroutine. close jobs when no more jobs for worker to do
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//more value will be false if jobs has been closed and all values in the channel have been received
			j, more := <- jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	//awaiting on the worker to be done
	<-done

	//reading from a closed channel succeeds immediately. returning zero value of the undelrying type
	//ok will be false to inidicate it was a zero value generated because the channel is closed and empty
	_, ok := <- jobs
	fmt.Println("received more jobs:", ok)
}