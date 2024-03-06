package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}


type UserSlice []User
// agar bisa menggunakan method sort.Sort(), harus mengimplementasikan interface Len(), Less(i,j int) bool, Swap(i, j int)

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	// return s[i].Name < s[j].Name //jika ingin sort by name
	return s[i].Age < s[j].Age // jika ingin sort by age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	// sama saja dengan 
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp
}

func main(){
	data := []User{
		{"Fauzan", 23},
		{"Susi", 22},
		{"Rudi", 23},
		{"Adit", 20},
	}
	sort.Sort(UserSlice(data))
	for _, v := range data {
		fmt.Println(v.Age)
		// fmt.Println(v.Name)
	}
	
}