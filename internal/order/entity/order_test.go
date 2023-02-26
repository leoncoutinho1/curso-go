package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}
func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid price")
}
func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenCreateANewOrder_ThenShouldReceiveCreatedOrderWithAllParams(t *testing.T) {
	order := Order{
		Id:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Nil(t, order.IsValid())
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenValidParams_WhenCallNewOrder_ThenShouldReceiveCreatedOrderAndNoError(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenValidParams_WhenCallCalculateFinalPrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
