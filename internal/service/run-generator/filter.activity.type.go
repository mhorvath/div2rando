package runGeneratorService

type ActivityTypeFilter struct{}

func (tf ActivityTypeFilter) Apply(option ActivityFilterOption, activities []Activity) []Activity {
	filteredActivities := []Activity{}
	for _, activity := range activities {
		for _, aT := range option.ACTIVITY_TYPES {
			if aT == activity.TYPE {
				filteredActivities = append(filteredActivities, activity)
			}
		}
	}
	return filteredActivities
}
