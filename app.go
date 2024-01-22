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
		runGenerator: &runGeneratorService.RunGenerator{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GenerateRun(runOption runGeneratorService.RunOption) runGeneratorService.Run {
	return a.runGenerator.GenerateRun(runOption)
}

func (a *App) GetActivityTypes() []runGeneratorService.ActivityType {
	return a.runGenerator.GetActivityTypes()
}

func (a *App) GetFractions() []runGeneratorService.Fraction {
	return a.runGenerator.GetFractions()
}

func (a *App) SaveRun(run runGeneratorService.Run) {
	utils.SaveJSON(fmt.Sprintf("run-%v.json", run.DATE), run)
}
