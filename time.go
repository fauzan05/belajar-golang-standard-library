package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// fmt.Printf("Formatnya %T, valuenya : %v\n", now, now.Local())

	// utc := time.Date(2023, time.March, 16, 30, 0, 0, 0, time.UTC)
	// fmt.Printf("Formatnya %T, valuenya : %v\n", utc, utc)

	// // time1 := "05:00:00PM"
	// // time2 := "17:00:00"
	// // t, err := time.Parse(time1, "07:00:00PM")
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	return
	// // }
	// // fmt.Println(t.Format(time1))
	// // fmt.Println(t.Format(time2))

	// layout1 := "03:04:05PM"
    // layout2 := "15:04:05"
    // t, err := time.Parse(layout1, "07:05:45PM")
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(t.Format(layout1))
    // fmt.Println(t.Format(layout2))

	// parse, _ := time.Parse(time.RFC3339, "2023-10-01T15:04:05Z")
	// fmt.Println(parse)

	currentTime := time.Now().Format("03:04:05PM")
	// fmt.Println(currentTime)
	resultParse, err := time.Parse("03:04:05PM", currentTime)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	} else {
		fmt.Println(resultParse.Format("15:04:05PM"))
	}
}
