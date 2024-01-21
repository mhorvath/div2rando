package runGeneratorService

import (
	"fmt"
	"math"
)

func sortActivityByDistance(a, b *Activity) bool {
	fmt.Printf("SORT DISTANCE \n")
	distance := calculateDistance(a.LOCATION, b.LOCATION)
	return distance < 0
}

func calculateDistance(location1, location2 Location) float64 {
	return math.Sqrt(math.Pow(float64(location1.X-location2.X), 2) + math.Pow(float64(location1.Y-location2.Y), 2))
}
