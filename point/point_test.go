package point

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPoint(t *testing.T) {
	Convey("Given a set of coordinats", t, func() {
		targetPoint := NewPointFromCoords(51.9204549, 4.5099984)
		otherPoint := NewPointFromCoords(51.9205376, 4.5100435)
		expectedDistance := 0.00970203531780931

		Convey("I expect to calculate correct distance", func() {
			targetPoint.CalculateDistance(otherPoint)
			distance := targetPoint.Distance

			So(distance, ShouldEqual, expectedDistance)
		})
	})
}
