package main
import "fmt"

//can specify if a channel is meant to only send or receive values thus increasing type-safety of the program
//ping accepts a channel for sending
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//accepts a channel for receiving (pings) and a channel for sending (pongs)
func pong(pings <-chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}