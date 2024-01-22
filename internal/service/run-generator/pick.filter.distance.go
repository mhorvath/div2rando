package runGeneratorService

import "fmt"

type PickFilterDistance struct {
	ACTIVITY_FILTER_OPTION ActivityFilterOption
	ACTIVITIES             *[]Activity
}

func (pf *PickFilterDistance) Validate(a Activity) bool {
	activities := *pf.ACTIVITIES
	if pf.ACTIVITY_FILTER_OPTION.MAX_DISTANCE == 0 || len(activities) == 0 {
		return true
	}
	loc1 := a.EXIT_LOCATION
	loc2 := activities[len(activities)-1].ENTRY_LOCATION
	distance := calculateLocationDistance(loc1, loc2)
	cd1To2 := calculateCelestialDirection(loc1, loc2)
	cd2To1 := calculateCelestialDirection(loc2, loc1)
	distancePenalty := loc1.PENALTY[cd1To2] + loc2.PENALTY[cd2To1]
	totalDistance := distance + float64(distancePenalty)
	fmt.Printf("------ %v | %v\n", distancePenalty, totalDistance)

	return totalDistance < float64(pf.ACTIVITY_FILTER_OPTION.MAX_DISTANCE)
}
