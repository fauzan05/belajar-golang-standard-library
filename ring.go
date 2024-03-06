package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main(){

	
	var data *ring.Ring = ring.New(5)
	for i := 1; i <= data.Len(); i++ {
		data.Value = "Value ke - " + strconv.Itoa(i)
		data = data.Next()
	}
	data.Do(func(a any) {
		fmt.Println(a)
	})
}