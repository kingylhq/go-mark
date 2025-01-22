package example

import "fmt"

func ChannelBuffering() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 10)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"
	messages <- "marklsy"
	messages <- "kinglsy"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
