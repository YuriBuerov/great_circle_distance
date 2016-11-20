package point

import (
	"fmt"
	"math"
	"strconv"
)

const (
	earthRadius = 6371
)

// Point represents structure with coords, id and distance to other point
type Point struct {
	ID       int
	Lat      float64
	Lng      float64
	Distance float64
}

// NewPointFromCoords initialize Point structure with lat and lng fields only
func NewPointFromCoords(lat, lng float64) *Point {
	return &Point{Lat: lat, Lng: lng}
}

// NewPoint initialize new point and convert fields
func NewPoint(line []string) *Point {
	id, _ := strconv.Atoi(line[0])
	lat, _ := strconv.ParseFloat(line[1], 64)
	lng, _ := strconv.ParseFloat(line[2], 64)

	return &Point{id, lat, lng, 0}
}

// CalculateDistance great circle distance calculation
func (p *Point) CalculateDistance(other *Point) {
	dLat := (other.Lat - p.Lat) * (math.Pi / 180.0)
	dLon := (other.Lng - p.Lng) * (math.Pi / 180.0)

	lat1 := p.Lat * (math.Pi / 180.0)
	lat2 := other.Lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	p.Distance = earthRadius * c
}

func (p *Point) String() string {
	return fmt.Sprintf("Distance:  %v, id: %v", p.Distance, p.ID)
}
