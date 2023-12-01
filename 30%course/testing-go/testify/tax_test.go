package testify

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax2(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax2(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")

}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Twice() // Determining limit to call CalculateTaxAndSave (only 2 times) that should return 10
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	// repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax")) <- you can use mock.Anything

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err1 := CalculateTaxAndSave(1000.0, repository) // this second call is only to show the log noticing the amount of call exceeded
	assert.Nil(t, err1)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
