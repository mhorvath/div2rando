package runGeneratorService

import (
	"changeme/internal/utils"
	"fmt"
	"math/rand"
)

var (
	DB_KEY_ACTIVITIES DBKey = "ACTIVITIES"
	ALL_DB_ACTIVITIES []Activity
	DB_FILES          = map[DBKey]string{
		DB_KEY_ACTIVITIES: "activities.json",
	}

	ACTIVITY_TYPE_MISSION    ActivityType = "MISSION"
	ACTIVITY_TYPE_BOUNTY     ActivityType = "BOUNTY"
	ACTIVITY_TYPE_CHECKPOINT ActivityType = "CHECKPOINT"

	FRACTION_HYENAS    Fraction = "Hyenas"
	FRACTION_OUTCASTS  Fraction = "Outcasts"
	FRACTION_TRUESONS  Fraction = "True Sons"
	FRACTION_BLACKTUSK Fraction = "Black Tusk"
	FRACTION_CLEANERS  Fraction = "Cleaners"
)

type DBKey string

type PickFilterFunc interface {
	Validate(activity Activity) bool
}

type RunGenerator struct{}

func (rg *RunGenerator) GenerateRun(option RunOption) Run {
	run := &Run{}

	fmt.Printf("%v", option)

	totalActivities := []Activity{}

	for _, activityFilterOption := range option.ACTIVITY_FILTER_OPTIONS {
		activities := []Activity{}

		pickFilterDuration := &PickFilterDuration{
			ACTIVITY_FILTER_OPTION: activityFilterOption,
			ACTIVITIES:             &activities,
		}
		pickFilterType := &PickFilterType{
			ACTIVITY_FILTER_OPTION: activityFilterOption,
		}
		pickFilterFraction := &PickFilterFraction{
			ACTIVITY_FILTER_OPTION: activityFilterOption,
		}
		pickFilterMultiplePick := &PickFilterMultiplePick{
			ACTIVITY_FILTER_OPTION: activityFilterOption,
			ACTIVITIES:             &activities,
		}

		allPickFilters := []PickFilterFunc{
			pickFilterMultiplePick,
			pickFilterType,
			pickFilterFraction,
			pickFilterDuration,
		}
		maxIV := 5000
		iv := 0

	RandomActivityLoop:
		for {
			iv += 1
			if iv >= maxIV {
				break RandomActivityLoop
			}
			randomActivity := rg.RandomActivityPick()
			for _, pf := range allPickFilters {
				isValid := pf.Validate(randomActivity)
				if !isValid && pf == pickFilterDuration {
					break RandomActivityLoop
				}
				if !isValid {
					continue RandomActivityLoop
				}
			}
			activities = append(activities, randomActivity)
		}
		totalActivities = append(totalActivities, activities...)
	}
	run.ACTIVITIES = totalActivities
	return *run
}

func (rg *RunGenerator) RandomActivityPick() Activity {
	randomIndex := rand.Intn(len(ALL_DB_ACTIVITIES))
	return ALL_DB_ACTIVITIES[randomIndex]
}

func (rg *RunGenerator) GetActivityTypes() []ActivityType {
	return []ActivityType{
		ACTIVITY_TYPE_BOUNTY,
		ACTIVITY_TYPE_CHECKPOINT,
		ACTIVITY_TYPE_MISSION,
	}
}

func (rg *RunGenerator) GetFractions() []Fraction {
	return []Fraction{
		FRACTION_BLACKTUSK,
		FRACTION_CLEANERS,
		FRACTION_HYENAS,
		FRACTION_OUTCASTS,
		FRACTION_TRUESONS,
	}
}

func init() {
	utils.LoadJSON(DB_FILES[DB_KEY_ACTIVITIES], &ALL_DB_ACTIVITIES)
	/*
		fmt.Println("--------------------------------------------------------------")
		rg := RunGenerator{}

		run := rg.GenerateRun(RunOption{
			ACTIVITY_FILTER_OPTIONS: []ActivityFilterOption{
				{
					ACTIVITY_TYPES: []ActivityType{ACTIVITY_TYPE_MISSION},
					FRACTIONS:      []Fraction{FRACTION_BLACKTUSK, FRACTION_CLEANERS, FRACTION_HYENAS, FRACTION_OUTCASTS, FRACTION_TRUESONS},
					DURATION:       time.Duration(30000 * time.Minute),
					MULTIPLE_PICK:  false,
				},
				{
					ACTIVITY_TYPES: []ActivityType{ACTIVITY_TYPE_BOUNTY},
					FRACTIONS:      []Fraction{FRACTION_BLACKTUSK},
					DURATION:       time.Duration(9 * time.Minute),
				},
				{
					ACTIVITY_TYPES: []ActivityType{ACTIVITY_TYPE_CHECKPOINT},
					FRACTIONS:      []Fraction{FRACTION_BLACKTUSK, FRACTION_CLEANERS, FRACTION_HYENAS},
					DURATION:       time.Duration(12 * time.Minute),
				},
			},
		})

		fmt.Printf("%v", run.ACTIVITIES)
	*/
}
