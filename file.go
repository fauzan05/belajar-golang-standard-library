package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666) // 666
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}
func addMessageToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666) // 666
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readNewFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "error", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// createNewFile("contoh.txt", "ini adalah contoh sampel file txt")
	addMessageToFile("contoh.txt", "\nini adalah pesan yang baru 1")

	read, err := readNewFile("contoh.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(read)
	}
}
