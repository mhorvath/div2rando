package runGeneratorService

type PickFilterMultiplePick struct {
	ACTIVITY_FILTER_OPTION ActivityFilterOption
	ACTIVITIES             *[]Activity
}

func (pf *PickFilterMultiplePick) Validate(a Activity) bool {
	if pf.ACTIVITY_FILTER_OPTION.MULTIPLE_PICK {
		return true
	}

	for _, vA := range *pf.ACTIVITIES {
		if a.NAME == vA.NAME && a.TYPE == vA.TYPE && a.FRACTION == vA.FRACTION {
			return false
		}
	}

	return true

}
