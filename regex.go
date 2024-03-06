package main

import (
	"fmt"
	"regexp"
)

func main(){
	regex := regexp.MustCompile(`f([a-z])n`)

	fmt.Println(regex.MatchString("fauzan")) // false
	fmt.Println(regex.MatchString("fan")) // true
	fmt.Println(regex.MatchString("fun")) // true

	fmt.Println(regex.FindAllString("fan fin fun fen fon ftn fmn frn fwn fvn fgn", 10))
}