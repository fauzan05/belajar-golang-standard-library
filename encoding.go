package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	encoded := base64.StdEncoding.EncodeToString([]byte("Fauzan Nur Hidayat"))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}