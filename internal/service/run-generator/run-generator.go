package runGeneratorService

import (
	"changeme/internal/utils"
	"errors"
	"math/rand"
	"time"
)

const (
	DB_KEY_ACTIVITIES        DBKey        = "ACTIVITIES"
	ACTIVITY_TYPE_MISSION    ActivityType = "MISSION"
	ACTIVITY_TYPE_BOUNTY     ActivityType = "BOUNTY"
	ACTIVITY_TYPE_CHECKPOINT ActivityType = "CHECKPOINT"
	FRACTION_HYENAS          Fraction     = "Hyenas"
	FRACTION_OUTCASTS        Fraction     = "Outcasts"
	FRACTION_TRUESONS        Fraction     = "True Sons"
	FRACTION_BLACKTUSK       Fraction     = "Black Tusk"
	FRACTION_CLEANERS        Fraction     = "Cleaners"
)

var (
	ALL_FILTERS []FilterFunc = []FilterFunc{
		ActivityTypeFilter{},
		ActivityDurationFilter{},
		ActivityFractionFilter{},
	}

	DB_FILE = map[DBKey]string{
		DB_KEY_ACTIVITIES: "activities.json",
	}

	DB_ACTIVITIES []Activity
)

type DBKey string

type FilterFunc interface {
	Apply(option ActivityFilterOption, activities []Activity) []Activity
}

type RunGenerator struct{}

func NewRunGenerator() *RunGenerator {
	rg := RunGenerator{}
	if err := utils.LoadJSON(DB_FILE[DB_KEY_ACTIVITIES], &DB_ACTIVITIES); err != nil {
		utils.SaveJSON(DB_FILE[DB_KEY_ACTIVITIES], &[]Activity{})
	}
	return &rg
}

func (rg *RunGenerator) GenerateRun(option RunOption) *Run {
	run := &Run{
		OPTION: option,
		DATE:   time.Now(),
	}
	for _, o := range option.ACTIVITY_FILTER_OPTIONS {
		activities := rg.filterActivities(o)
		run.FILTERED_ACTIVITIES = append(run.FILTERED_ACTIVITIES, activities)
	}
	activities := run.GetFilteredActivities()

	run.ACTIVITIES = randomizeActivities(activities)
	return run
}

func (rg *RunGenerator) filterActivities(option ActivityFilterOption) []Activity {
	ra := DB_ACTIVITIES
	for i := 0; i < 5; i++ {
		ra = randomizeActivities(ra)
	}
	for _, f := range ALL_FILTERS {
		ra = f.Apply(option, ra)
	}

	return ra
}

// UTIL
func randomizeActivities(slice []Activity) []Activity {
	n := len(slice)
	result := make([]Activity, n)
	copy(result, slice)

	// Fisher-Yates Shuffle Algorithmus
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func getSortFunctionByName(name SortFunctionName) (SortFunction, error) {
	switch name {
	case "duration":
		return sortActivityByDuration, nil
	case "type":
		return sortActivityByType, nil
	case "distance":
		return sortActivityByDistance, nil
	}
	return nil, errors.New("invalid sort name")
}

func (rs *RunGenerator) GetActivityTypes() []ActivityType {
	return []ActivityType{
		ACTIVITY_TYPE_BOUNTY,
		ACTIVITY_TYPE_CHECKPOINT,
		ACTIVITY_TYPE_MISSION,
	}
}
