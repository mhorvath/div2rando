package runGeneratorService

import "math"

func calculateLocationDistance(loc1, loc2 Location) float64 {
	return math.Sqrt(math.Pow(float64(loc2.X)-float64(loc1.X), 2) + math.Pow(float64(loc2.Y)-float64(loc1.Y), 2))
}
