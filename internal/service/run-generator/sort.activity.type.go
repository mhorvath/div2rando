package runGeneratorService

import "fmt"

func sortActivityByType(a, b *Activity) bool {
	fmt.Printf("SORT TYPE \n")
	return a.TYPE > b.TYPE
}
