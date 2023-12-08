package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := (float64(productionRate) * successRate) / 100
	return int(carsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10

	individualCars := carsCount % 10

	costOfGroups := uint(groupsOfTen) * 95000

	costOfIndividual := uint(individualCars) * 10000

	totalCost := costOfGroups + costOfIndividual

	return totalCost
}