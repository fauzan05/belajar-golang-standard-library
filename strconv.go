package main

import (
	"fmt"
	"strconv"
)

func main(){
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Printf("Tipe datanya : %T, valuenya : %b", boolean)
	} else {
		fmt.Println("Error", err.Error())
	}
}