package main

import (
	"fmt"
	"path"
)

func main(){
	fmt.Println(path.Dir("makan/hello/world.go"))
	fmt.Println(path.Base("makan/hello/world.go"))
	fmt.Println(path.Ext("makan/hello/world.go"))
	fmt.Println(path.Join("makan", "hello", "world.go"))
}