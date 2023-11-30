package testing_go

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTax2(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax2(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, tax := range table {
		result := CalculateTax(tax.amount)
		if result != tax.expect {
			t.Errorf("Expected %f but got %f", tax.expect, result)
		}
	}

}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.00)
	}
}

// This is necessary to test
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.00)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Expected 0 but got %f", result)
		}
		if amount > 2000 && result != 20 {
			t.Errorf("Expected 0 but got %f", result)
		}
	})

}
