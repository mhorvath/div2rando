package runGeneratorService

import "time"

type ActivityType string

type Fraction string

type Location struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

type Activity struct {
	NAME        string        `json:"name"`
	DURATION    time.Duration `json:"duration"`
	TYPE        ActivityType  `json:"type"`
	FRACTION    Fraction      `json:"fraction"`
	LOCATION    Location      `json:"location"`
	FAST_TRAVEL bool          `json:"fastTravel"`
}

func (a *Activity) AVGDuration() time.Duration {
	return a.DURATION
}
