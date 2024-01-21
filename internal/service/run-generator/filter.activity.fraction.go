package runGeneratorService

type ActivityFractionFilter struct{}

func (tf ActivityFractionFilter) Apply(option ActivityFilterOption, activities []Activity) []Activity {
	filteredActivities := []Activity{}
	for _, activity := range activities {
		for _, aF := range option.FRACTIONS {
			if aF == activity.FRACTION {
				filteredActivities = append(filteredActivities, activity)
			}
		}
	}
	return filteredActivities
}
