package runGeneratorService

import "time"

type ActivityDurationFilter struct{}

func (tf ActivityDurationFilter) Apply(option ActivityFilterOption, activities []Activity) []Activity {
	as := []Activity{}
	for _, a := range activities {
		if getAVGDurationOfActivities(as)+a.DURATION < option.DURATION {
			as = append(as, a)
		}
	}
	return as
}

func getAVGDurationOfActivities(activities []Activity) time.Duration {
	duration := time.Duration(0)
	for _, a := range activities {
		duration += a.AVGDuration()
	}

	return duration
}
