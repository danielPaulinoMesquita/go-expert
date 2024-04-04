package testing_go

import "time"

func CalculateTax(amount float64) float64 {
	if amount <= 0 { // if the rule was amount < 0 it wll throw an error in Fuzzy
		return 0
	}

	if amount >= 1000 && amount <= 2000 {
		return 10.0
	}

	//if amount >= 1000 { // <- If this code stay with this rule, it makes error on the Fuzzy
	//	return 10.0
	//}

	if amount >= 2000 {
		return 20.0
	}

	return 5.0
}

// CalculateTax2 it will be tested for checking what mathod is more speed.
func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
