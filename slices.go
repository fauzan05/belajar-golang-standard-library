package main

import (
	"fmt"
	"slices"
)

func main(){
	names := []string{"fauzan", "rudi", "heri", "susi", "fida"}
	values := []int{23, 34, 21, 30, 20}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
}