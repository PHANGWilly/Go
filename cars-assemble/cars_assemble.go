package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var result float64
	result = float64(productionRate) * successRate / 100
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var result int
	result = int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
	return result
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups := carsCount / 10
	individuals := carsCount % 10
	return uint(groups*95000 + individuals*10000)
}
