package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main(){
	input := strings.NewReader("ini adalah contoh teks yang panjang\ndengan new line spasi \n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		// fmt.Println(reader.ReadString(byte(reader.Buffered())))
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}