package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	//non blocking receives. it is non blocking because there is a default clause
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	//non blocking sends. in this case msg cant be sent because messages has no buffer and there is no receiver so default case gets selected

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("receied signals", sig)
	default:
		fmt.Println("no activity")
	}
}