package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1 string, option2 string) string {
	better := " is clearly the better choice."
	var test string
	if option1 < option2 {
		test := option1 + better
		return test
	}
	if option2 < option1 {
		test := option2 + better
		return test
	}
	return test
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice float64, age float64) float64 {
	var resellPrice float64
	if age < 3 {
		resellPrice = originalPrice * 80 / 100
	} else if age >= 10 {
		resellPrice = originalPrice * 50 / 100
	} else {
		resellPrice = originalPrice * 70 / 100
	}

	return resellPrice
}
