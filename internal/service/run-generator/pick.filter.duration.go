package runGeneratorService

import "time"

type PickFilterDuration struct {
	ACTIVITY_FILTER_OPTION ActivityFilterOption
	ACTIVITIES             *[]Activity
}

func (pf *PickFilterDuration) Validate(a Activity) bool {
	return pf.getAVGDurationOfActivities()+a.AVGDuration() < pf.ACTIVITY_FILTER_OPTION.DURATION
}

func (pf *PickFilterDuration) getAVGDurationOfActivities() time.Duration {
	duration := time.Duration(0)
	for _, a := range *pf.ACTIVITIES {
		duration += a.AVGDuration()
	}
	return duration
}
