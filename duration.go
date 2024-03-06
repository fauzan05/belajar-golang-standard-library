package main

import (
	"fmt"
	"time"
)

func main(){
	var duration1 time.Duration = time.Second * 60
	var duration2 time.Duration = time.Second * 10
	var result time.Duration = duration1 - duration2
	fmt.Println(result)

}