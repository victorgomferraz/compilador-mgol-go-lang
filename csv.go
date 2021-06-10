package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func CsvToMap(filename string) map[string]map[string]string {
	csvFile, err := os.Open(filename)
	checkErr(err)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ','

	var res = make(map[string]map[string]string)
	var heading = make(map[int]string)
	i := 0

	for {
		line, err := reader.Read() //para cada linha
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr(err)
		}
		if i ==0 {
			for j, v := range line {
				heading[j] = v
			}
		} else{
			res[line[0]] = make(map[string]string)
			for j, v := range line {
				if i >0 {
					res[line[0]][heading[j]] = v
				}
			}
		}

		i++
	}
	return res
}
