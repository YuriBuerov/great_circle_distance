package main

import (
	"fmt"
	"github.com/YuriBuerov/great_circle_distance/csv"
	"github.com/YuriBuerov/great_circle_distance/point"
	"github.com/YuriBuerov/great_circle_distance/result"

	"path/filepath"
)

const filePath = "csv/testdata/geoData.csv"

func main() {
	path, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	parser := csv.NewParser(path)
	points, err := parser.Parse()

	if err != nil {
		fmt.Println(err)
		return
	}

	targetPoint := point.NewPointFromCoords(51.9204549, 4.5099984)

	topResults := result.NewResult(result.TopResultType, 5, targetPoint)
	bottomResults := result.NewResult(result.BottomResultType, 5, targetPoint)

	ready := make(chan bool)
	defer func() {
		close(ready)
	}()

	go func() {
		for p := range points {
			fmt.Println("for p := range points")
			fmt.Println(p)
			topResults.Check(p)
			bottomResults.Check(p)
		}

		ready <- true
	}()

	<-ready

	// print results

	fmt.Println(topResults)
	fmt.Println(bottomResults)
}
