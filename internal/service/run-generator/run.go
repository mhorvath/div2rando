package runGeneratorService

import "time"

type ActivityFilterOption struct {
	DURATION       time.Duration
	ACTIVITY_TYPES []ActivityType
	FRACTIONS      []Fraction
}

type SortFunction func(a, b *Activity) bool

type SortFunctionName string

type RunOption struct {
	ACTIVITY_FILTER_OPTIONS []ActivityFilterOption
	SORT_FUNCTION_NAMES     []SortFunctionName
}

type Run struct {
	FILTERED_ACTIVITIES [][]Activity
	DATE                time.Time
	OPTION              RunOption
	ACTIVITIES          []Activity
}

func (s *Run) GetFilteredActivities() []Activity {
	as := []Activity{}
	for _, fA := range s.FILTERED_ACTIVITIES {
		as = append(as, fA...)
	}
	return as
}

func (s *Run) Duration() time.Duration {
	var totalDuration time.Duration
	for _, a := range s.GetFilteredActivities() {
		totalDuration += a.DURATION
	}
	return totalDuration
}

func (s *Run) AVGDuration() time.Duration {
	var totalDuration time.Duration
	for _, a := range s.GetFilteredActivities() {
		totalDuration += a.AVGDuration()
	}
	return totalDuration
}
