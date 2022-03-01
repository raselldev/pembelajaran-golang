package golanggoroutines

import (
	"fmt"
	"testing"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	channel <- "Raka"

	data := <-channel

	fmt.Println(data)

	close(channel)
}
