package runGeneratorService

type PickFilterFraction struct {
	ACTIVITY_FILTER_OPTION ActivityFilterOption
}

func (pft *PickFilterFraction) Validate(a Activity) bool {
	for _, fraction := range pft.ACTIVITY_FILTER_OPTION.FRACTIONS {
		if a.FRACTION == fraction {
			return true
		}
	}
	return false
}
