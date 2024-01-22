package runGeneratorService

import (
	"encoding/json"
	"time"
)

type ActivityType string

type Fraction string

type Location struct {
	X       uint                        `json:"x"`
	Y       uint                        `json:"y"`
	PENALTY map[CelestialDirection]uint `json:"penalty"`
}

type Activity struct {
	NAME           string        `json:"name"`
	DURATION       time.Duration `json:"duration"`
	TYPE           ActivityType  `json:"type"`
	FRACTION       Fraction      `json:"fraction"`
	ENTRY_LOCATION Location      `json:"entryLocation"`
	EXIT_LOCATION  Location      `json:"exitLocation"`
	FAST_TRAVEL    bool          `json:"fastTravel"`
}

func (a *Activity) AVGDuration() time.Duration {
	return a.DURATION
}

func (a *Activity) UnmarshalJSON(data []byte) error {
	type Alias Activity
	aux := &struct {
		*Alias
		ExitLocationEmpty bool `json:"exitLocation,omitempty"`
	}{Alias: (*Alias)(a)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Falls exitLocation nicht im JSON vorhanden ist, setze es auf entryLocation
	if aux.ExitLocationEmpty {
		aux.EXIT_LOCATION = aux.ENTRY_LOCATION
	}

	return nil
}
