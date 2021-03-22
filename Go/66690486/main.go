package main

import (
	"fmt"
	csvutils "github.com/alessiosavi/GoGPUtils/csv"
	"io/ioutil"
)

type Asset struct {
	Id    string `json:"id"`
	Color string `json:"color"`
	Owner string `json:"owner"`
}

func main() {
	data, err := ioutil.ReadFile("assets.csv")
	if err != nil {
		panic(err)
	}
	headers, csvData := csvutils.ReadCsv(data, ',')
	var assets []Asset = make([]Asset, len(csvData))
	fmt.Println("The csv file have the following headers: ", headers)

	for i, line := range csvData {
		asst := Asset{
			Id:    line[0],
			Color: line[1],
			Owner: line[2],
		}
		assets[i] = asst
	}
}
