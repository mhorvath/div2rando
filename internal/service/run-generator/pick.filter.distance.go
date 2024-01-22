package runGeneratorService

type PickFilterDistance struct {
	ACTIVITY_FILTER_OPTION ActivityFilterOption
	ACTIVITIES             *[]Activity
}

func (pf *PickFilterDistance) Validate(a Activity) bool {
	activities := *pf.ACTIVITIES
	if pf.ACTIVITY_FILTER_OPTION.MAX_DISTANCE == 0 || len(activities) == 0 {
		return true
	}
	loc1 := a.LOCATION
	loc2 := activities[len(activities)-1].LOCATION
	return calculateLocationDistance(loc1, loc2) < float64(pf.ACTIVITY_FILTER_OPTION.MAX_DISTANCE)
}
