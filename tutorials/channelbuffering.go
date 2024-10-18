package main

import "fmt"

func main() {

	//by default channels are unbuffered
	//meaning they only accept sends chan <- if there is a corresponding receive <- chan ready to receive the sent value

	//messages can buffer 2  values  without needing a corresponding concurrent receive
	messages := make(chan string, 2)

	messages <- "bufferred"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}