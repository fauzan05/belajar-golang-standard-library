package main

import (
	"fmt"
	"errors"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error{
	if id == "" {
		return ValidationError
	}

	if id != "kode" {
		return NotFoundError
	}

	return nil
}

func main(){
	err := GetById("kode")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
