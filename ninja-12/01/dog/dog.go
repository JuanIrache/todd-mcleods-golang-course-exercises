// Package dog does operations related to dogs
package dog

// YearsHumanToDog converts human years int to dog years int
func YearsHumanToDog(y int) int {
	return y * 7
}

// YearsDogToHuman converts dog years int to human years int
func YearsDogToHuman(y int) int {
	return y / 7
}
