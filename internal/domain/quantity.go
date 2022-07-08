package domain

import (
	"fmt"
)

var (
	ErrQuantityTooSmall = EInvalid("quantity must be greater than 0")
)

type Quantity struct {
	value float64
	unit  QuantityUnit
}

func NewQuantity(value float64, unit QuantityUnit) (Quantity, error) {
	if value <= 0 {
		return Quantity{}, ErrQuantityTooSmall
	}

	return Quantity{value: value, unit: unit}, nil
}

func (q Quantity) Value() float64 {
	return q.value
}

func (q Quantity) Unit() QuantityUnit {
	return q.unit
}

func (q Quantity) IsZero() bool {
	return q.value == 0
}

func (q Quantity) String() string {
	return fmt.Sprintf("%g%s", q.value, q.unit)
}
