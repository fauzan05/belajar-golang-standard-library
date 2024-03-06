package main

import (
	"flag"
	"fmt"
)

func main(){
	host := flag.String("host", "localhost", "put the database hostname")
	username := flag.String("username", "root", "put the database username")
	password := flag.String("password", "root", "put the database password")
	port := flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println(*host, *username, *password, *port) // karena flag.String itu menyimpan dalam bentuk pointer, maka memanggilnya menggunnakan * (asterisk) agar diambil valuenya, bukan address memorynya
	// jika ingin mengisi manual, maka = go run flag.go -host="localhost1" -username="fauzan" -password="password1" -port=8001
}