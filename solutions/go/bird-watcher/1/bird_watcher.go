package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, v := range birdsPerDay {
		count += v
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := min(start+7, len(birdsPerDay))

	sum := 0

	for _, v := range birdsPerDay[start:end] {
		sum += v
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	copyBidsCount := birdsPerDay
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			copyBidsCount[i] = v + 1
		}
	}
	return copyBidsCount
}
