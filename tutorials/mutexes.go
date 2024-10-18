package main

import (
	"fmt"
	"sync"
)

//container holds a map of counters since we want to update it concurrently from multiple goroutines we add a mutex to synchronize access
//mutexes shouldn't b copied so this must be passed around as a pointer and not by value
type Container struct {
	mu sync.Mutex
	counters map[string] int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	//unlock at end of function
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)

	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}