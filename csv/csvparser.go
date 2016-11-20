package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/YuriBuerov/great_circle_distance/point"
	"io"
	"log"
	"os"
)

// Parser represents structure with file path
type Parser struct {
	filePath string
}

// NewParser initialize Parse
func NewParser(filePath string) *Parser {
	return &Parser{filePath}
}

// Parse - parse csv file
func (c *Parser) Parse() (<-chan *point.Point, error) {
	file, err := os.Open(c.filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)

	// get CSV Header
	_, err = csvReader.Read()

	if err != nil {
		return nil, err
	}

	out := make(chan *point.Point, 100)

	go func() {
		defer func() {
			close(out)
		}()
		defer file.Close()

		// For each record
		for {
			line, err := csvReader.Read()

			fmt.Println("line")
			fmt.Println(line)

			if err != nil {
				// abort reading on parse error
				if cerr, ok := err.(*csv.ParseError); ok {
					log.Println(cerr)
					return
				} else if err == io.EOF {
					log.Println("Read EOF from file")
					return
				} else {
					log.Println("Some error")
					return
				}
			}

			if len(line) == 0 {
				continue
			}

			out <- point.NewPoint(line)
		}
	}()

	return out, nil
}
