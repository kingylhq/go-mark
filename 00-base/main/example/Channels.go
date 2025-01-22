package example

import "fmt"

func Channels() {

	ch := make(chan string)

	go func() { ch <- "ping" }()

	msg := <-ch
	fmt.Println(msg)
}
