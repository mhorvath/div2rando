package runGeneratorService

type PickFilterType struct {
	ACTIVITY_FILTER_OPTION ActivityFilterOption
}

func (pf *PickFilterType) Validate(a Activity) bool {
	for _, aType := range pf.ACTIVITY_FILTER_OPTION.ACTIVITY_TYPES {
		if a.TYPE == aType {
			return true
		}
	}
	return false
}
