package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("fauzan nur hida", "nur"))
	fmt.Println(strings.Split("fauzan nur hida", "hida"))
	fmt.Println(strings.ToLower("FAUZAN NUR HIDAYAT"))
	fmt.Println(strings.Trim(" Fauzan Nur ", " "))
	fmt.Println(strings.ReplaceAll("Fauzan Nur Hidayat", "Nur", "New"))
}