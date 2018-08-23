package kcsv

import (
	"log"
)

type Record struct {
	csv  *Csv
	data []string
}

func NewRecord(csv *Csv, row []string) *Record {
	r := &Record{}
	r.data = row
	r.csv = csv
	return r
}

func (r *Record) Delete() {
	r.data = nil
}

func (r *Record) isEmpty() bool {
	return r.data == nil
}

func (r *Record) Update(row map[string]string) {
	for key, value := range row {
		index := r.csv.fieldToIndex(key)
		r.data[index] = value
	}
}

func (r *Record) Get(field string) string {
	index := r.csv.fieldToIndex(field)
	if index == -1 {
		log.Fatalln("error: key is nothing")
	}
	return r.data[index]
}

func (r *Record) Set(field, value string) {
	index := r.csv.fieldToIndex(field)
	if index == -1 {
		log.Fatalln("error: key is nothing")
	}
	r.data[index] = value
}
