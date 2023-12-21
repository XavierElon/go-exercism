package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birdCount int
	for i := 0; i < len(birdsPerDay); i++ {
		birdCount += birdsPerDay[i]
	}
	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdCountInWeek int

	for i := (week - 1) * 7; i < week*7; i++ {
		birdCountInWeek += birdsPerDay[i]
	}
	return birdCountInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	newBirdsPerDay := birdsPerDay
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			newBirdsPerDay[i] += 1
		}
	}
	return newBirdsPerDay
}
