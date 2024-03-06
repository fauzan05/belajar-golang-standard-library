package main

import (
	"container/list"
	"fmt"
)

func main(){
	data := list.New()
	data.PushBack("Fauzan")
	data.PushBack("Nur")
	data.PushBack("Hidayat")

	var firstData *list.Element = data.Front()
	fmt.Println(firstData.Value)
	// fmt.Println(data.Front().Value)
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}