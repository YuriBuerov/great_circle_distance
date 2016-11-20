package csv

import (
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("Given a set of records", t, func() {
		var filePath = "testdata/sample.csv"
		var rowCount = 7

		Convey("I expect to process all valid rows", func() {
			var lineCount int
			path, err := filepath.Abs(filePath)

			parser := NewParser(path)
			lines, err := parser.Parse()
			So(err, ShouldBeNil)

			for _ = range lines {
				lineCount++
			}

			So(lineCount, ShouldEqual, rowCount)
		})
	})
}
