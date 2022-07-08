package domain

import (
	"errors"
	"fmt"
)

var (
	ErrServingLessThanZero = errors.New("must be greater than zero")
)

type ServingType struct {
	code string
}

var (
	ServingTypeGuest = ServingType{"guest"}
	ServingTypeUnit  = ServingType{"unit"}
)

func (t ServingType) IsZero() bool {
	return t.code == ""
}

func (t ServingType) String() string {
	return t.code
}

type Serving struct {
	kind   ServingType
	number int
}

func NewGuestServing(guests int) (Serving, error) {
	if guests <= 0 {
		return Serving{}, ErrServingLessThanZero
	}

	return Serving{kind: ServingTypeGuest, number: guests}, nil
}

func NewUnitServing(units int) (Serving, error) {
	if units <= 0 {
		return Serving{}, ErrServingLessThanZero
	}

	return Serving{kind: ServingTypeUnit, number: units}, nil
}

func (s Serving) Number() int {
	return s.number
}

func (s Serving) Type() ServingType {
	return s.kind
}

func (s Serving) String() string {
	str := fmt.Sprintf("%d %s", s.number, s.kind)
	if s.number > 1 {
		str += "s"
	}

	return str
}

func (s Serving) IsZero() bool {
	return s.number == 0 || s.kind.IsZero()
}
