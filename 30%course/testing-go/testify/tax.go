package testify

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 && amount <= 2000 {
		return 10.0
	}

	if amount >= 2000 {
		return 20.0
	}

	return 5.0
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	return repository.SaveTax(tax)
}

func CalculateTax2(amount float64) (float64, error) {
	if amount <= 0 { // if the rule was amount < 0 it wll throw an error in Fuzzy
		return 0.0, errors.New("amount must be greater than 0")
	}

	if amount >= 1000 && amount <= 2000 {
		return 10.0, nil
	}

	if amount >= 2000 {
		return 20.0, nil
	}

	return 5.0, nil
}
