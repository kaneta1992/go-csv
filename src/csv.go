package kcsv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Csv struct {
	fieldNames []string
	records    []*Record
}

func New(fieldNames []string) *Csv {
	c := &Csv{
		fieldNames: fieldNames,
		records:    []*Record{},
	}
	return c
}

func NewFromFile(filepath string) *Csv {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	fieldNames, err := r.Read()
	if err == io.EOF {
		log.Fatalln("error: loaded empty text")
	}
	c := New(fieldNames)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		record := NewRecord(c, row)
		c.records = append(c.records, record)
	}
	return c
}

func (c *Csv) iterRecord(f func(*Record)) {
	for _, record := range c.records {
		if record.isEmpty() {
			continue
		}
		f(record)
	}
}

func (c *Csv) Delete() {
	for _, record := range c.records {
		record.Delete()
	}
}

func (c *Csv) fieldToIndex(field string) int {
	for index, value := range c.fieldNames {
		if value == field {
			return index
		}
	}
	return -1
}

func (c *Csv) Where(fields map[string][]string) *Csv {
	newCsv := &Csv{fieldNames: c.fieldNames}
	records := []*Record{}

	c.iterRecord(func(record *Record) {
		for key, values := range fields {
			index := c.fieldToIndex(key)
			if index == -1 {
				log.Fatalln("error: key is nothing")
			}
			for _, value := range values {
				if record.data[index] == value {
					records = append(records, record)
					continue
				}
			}

		}
	})

	newCsv.records = records
	return newCsv
}

func (c *Csv) ToArray() [][]string {
	ret := [][]string{}
	ret = append(ret, c.fieldNames)

	c.iterRecord(func(record *Record) {
		ret = append(ret, record.data)
	})

	return ret
}

func (c *Csv) Add(data []map[string]string) *Csv {
	for _, d := range data {
		row := make([]string, len(c.records))
		for key, value := range d {
			index := c.fieldToIndex(key)
			if index == -1 {
				log.Fatalln("error: key is nothing")
			}
			row[index] = value
		}
		record := NewRecord(c, row)
		c.records = append(c.records, record)
	}
	return c
}

func (c *Csv) First() *Record {
	for _, record := range c.records {
		if record.isEmpty() {
			continue
		}
		return record
	}
	return nil
}

func (c *Csv) Update(row map[string]string) *Csv {
	c.iterRecord(func(record *Record) {
		record.Update(row)
	})
	return c
}

func (c *Csv) GetRecords() []*Record {
	return c.records
}

func (c *Csv) Get(field string) []string {
	ret := []string{}
	c.iterRecord(func(record *Record) {
		ret = append(ret, record.Get(field))
	})
	return ret
}
