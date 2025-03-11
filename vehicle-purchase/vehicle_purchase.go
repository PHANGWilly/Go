package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var finalOption string
	
	if option1 < option2 {
		finalOption = option1
	} else {
		finalOption = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", finalOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 3 && age < 10 {
		return float64(originalPrice * 0.70)
	} else if age >= 10 {
		return float64(originalPrice * 0.50)
	} else if age < 3{
		return float64(originalPrice * 0.80)
	} else {
		return originalPrice
	}
}
