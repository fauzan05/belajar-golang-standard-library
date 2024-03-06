package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	// csv reader
	csvString := "fauzan, nur, hidayat\n" + "nurhadi, aldo\n" + "susi, anjar, sari\n"
	// csvString := ""
	reader := csv.NewReader(strings.NewReader(csvString)) 
	// fmt.Println(reader)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			// fmt.Println(io.EOF)
			fmt.Println("tidak ada data yang dibaca (kosong)")
			break
		}
		fmt.Println(record)
	}

	// csv writer
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"fauzan", "nur", "hidayat"})
	_ = writer.Write([]string{"nurhadi", "aldo"})
	_ = writer.Write([]string{"susi", "anjar", "sari"})
	writer.Flush()

}