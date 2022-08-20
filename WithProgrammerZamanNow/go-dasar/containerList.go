package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Raka")
	data.PushBack("Rasell")
	data.PushBack("RakaRasell")

	fmt.Println(data.Front().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
