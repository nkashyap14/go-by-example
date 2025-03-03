package main

//timer and ticker features can be used to execute go at some point in the future or repeatedly at some interval
import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	//blocks on timer's channel c until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	//timer's are different than time.sleep as we can cancel the timer before it fires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}