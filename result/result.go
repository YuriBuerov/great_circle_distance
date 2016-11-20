package result

import (
	"fmt"
	"github.com/YuriBuerov/great_circle_distance/point"
	"sort"
)

// Result represents structure with set of points and target point
type Result struct {
	points      []*point.Point
	targetPoint *point.Point
	amount      int
	Type
}

// NewResult initialize Result structure
func NewResult(t Type, amount int, targetPoint *point.Point) *Result {
	return &Result{make([]*point.Point, 0), targetPoint, amount, t}
}

// Check compare Result with Point
func (r *Result) Check(p *point.Point) {
	// calculate distance
	p.CalculateDistance(r.targetPoint)

	r.points = append(r.points, p)

	fmt.Println("r.points")
	fmt.Println(r.points)

	if r.Type == TopResultType {
		sort.Sort(point.ByTop(r.points))
	} else {
		sort.Sort(point.ByBottom(r.points))
	}

	if len(r.points) > 5 {
		r.points = r.points[:5]
	}
}

func (r *Result) String() string {
	var resultsName string

	if r.Type == TopResultType {
		resultsName = "Top"
	} else {
		resultsName = "Bottom"
	}

	fmt.Println(r.points)

	return fmt.Sprintf("Top %s Results %v", resultsName, r.points[:5])
}
