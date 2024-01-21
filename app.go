package main

import (
	runGeneratorService "changeme/internal/service/run-generator"
	"changeme/internal/utils"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx          context.Context
	runGenerator *runGeneratorService.RunGenerator
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		runGenerator: runGeneratorService.NewRunGenerator(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GenerateRun(runOption runGeneratorService.RunOption) *runGeneratorService.Run {
	fmt.Printf("----- RUN GENERATION CONFIG: %v", runOption)
	return a.runGenerator.GenerateRun(runOption)
}

func (a *App) GetActivityTypes() []runGeneratorService.ActivityType {
	return a.runGenerator.GetActivityTypes()
}

func (a *App) SaveRun(run runGeneratorService.Run) {
	utils.SaveJSON(fmt.Sprintf("run-%v.json", run.DATE), run)
}

/*

runGeneratorService.RunOption{
		ACTIVITY_FILTER_OPTIONS: []runGeneratorService.ActivityFilterOption{
			{
				ACTIVITY_TYPES: []runGeneratorService.ActivityType{runGeneratorService.ACTIVITY_TYPE_MISSION},
				FRACTIONS:      []runGeneratorService.Fraction{runGeneratorService.FRACTION_BLACKTUSK, runGeneratorService.FRACTION_CLEANERS, runGeneratorService.FRACTION_HYENAS, runGeneratorService.FRACTION_OUTCASTS, runGeneratorService.FRACTION_TRUESONS},
				DURATION:       time.Duration(40 * float64(time.Minute)),
			},
			{
				ACTIVITY_TYPES: []runGeneratorService.ActivityType{runGeneratorService.ACTIVITY_TYPE_BOUNTY},
				FRACTIONS:      []runGeneratorService.Fraction{runGeneratorService.FRACTION_BLACKTUSK, runGeneratorService.FRACTION_CLEANERS, runGeneratorService.FRACTION_HYENAS, runGeneratorService.FRACTION_OUTCASTS, runGeneratorService.FRACTION_TRUESONS},
				DURATION:       time.Duration(15 * float64(time.Minute)),
			},
			{
				ACTIVITY_TYPES: []runGeneratorService.ActivityType{runGeneratorService.ACTIVITY_TYPE_CHECKPOINT},
				FRACTIONS:      []runGeneratorService.Fraction{runGeneratorService.FRACTION_BLACKTUSK, runGeneratorService.FRACTION_CLEANERS, runGeneratorService.FRACTION_HYENAS, runGeneratorService.FRACTION_OUTCASTS, runGeneratorService.FRACTION_TRUESONS},
				DURATION:       time.Duration(20 * float64(time.Minute)),
			},
		},
		SORT_FUNCTION_NAMES: []runGeneratorService.SortFunctionName{
			// "type",
			// "duration",
			//"distance",
		},
	}

*/
