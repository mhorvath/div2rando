package runGeneratorService

import "math"

type CelestialDirection string

func calculateCelestialDirection(locA, locB Location) CelestialDirection {
	// Berechne den Winkel zwischen den beiden Punkten
	angle := math.Atan2(float64(locB.Y-locA.Y), float64(locB.X-locA.X)) * (180.0 / math.Pi)

	// Normalisiere den Winkel auf den Bereich von 0 bis 360 Grad
	if angle < 0 {
		angle += 360
	}

	// Bestimme die Himmelsrichtung basierend auf dem Winkel
	var direction CelestialDirection
	switch {
	case angle >= 315 || angle < 45:
		direction = "n"
	case angle >= 45 && angle < 135:
		direction = "o"
	case angle >= 135 && angle < 225:
		direction = "s"
	case angle >= 225 && angle < 315:
		direction = "w"
	}

	return direction
}
