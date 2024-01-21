package runGeneratorService

import "fmt"

func sortActivityByDuration(a, b *Activity) bool {
	fmt.Printf("SORT DURATION \n")
	return a.DURATION > b.DURATION
}
