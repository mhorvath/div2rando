package runGeneratorService

import "time"

type ActivityFilterOption struct {
	DURATION       time.Duration
	ACTIVITY_TYPES []ActivityType
	FRACTIONS      []Fraction
	MULTIPLE_PICK  bool
	MAX_DISTANCE   uint
}

type SortFunction func(a, b *Activity) bool

type SortFunctionName string

type RunOption struct {
	ACTIVITY_FILTER_OPTIONS []ActivityFilterOption
	SORT_FUNCTION_NAMES     []SortFunctionName
}

func (ro *RunOption) GetTotalDuration() time.Duration {
	duration := time.Duration(0)
	for _, afo := range ro.ACTIVITY_FILTER_OPTIONS {
		duration += afo.DURATION
	}
	return duration
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
