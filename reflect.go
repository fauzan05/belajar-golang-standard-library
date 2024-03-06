package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Name string `required:"true" max:"10"`
	Year int
}

func readFile(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type name : ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, " with type : ", structField.Type, " tag : ", structField.Tag.Get("required"))
		// structField.Tag.Get("required")
	}
}

func isValid(value any) bool {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			// fmt.Println(data)
			if data != "" {
				fmt.Println(data)
				return true
			} else if data == "" {
				fmt.Println("data kosong")
				return false
			}
		}
	}
	return true
}

func main(){
	car := Car{"ferrari", 2011}
	// sampleCarType := reflect.TypeOf(car)
	// structField := sampleCarType.Field(1)

	// fmt.Println(sampleCarType.Name())
	// fmt.Println(structField)
	// readFile(car)
	isValid(car)
	// array1 := []string{"satu","dua"}
	// data1 := 1312312312312
	// var sample reflect.Type = reflect.TypeOf(array1)
	// var sample1 reflect.Type = reflect.TypeOf(data1)
	// fmt.Println(sample)
	// fmt.Println(sample1)


}