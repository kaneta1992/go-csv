package main

import (
	"fmt"

	"github.com/kaneta1992/go-csv/src"
)

func main() {
	csv := kcsv.NewFromFile("strings_m.csv")
	fmt.Println(csv.Where(map[string][]string{
		"string_key": {"support"},
	}).ToArray())
}
