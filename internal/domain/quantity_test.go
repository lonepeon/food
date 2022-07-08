package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/food/internal/domain/domaintest"
	"github.com/lonepeon/golib/testutils"
)

func TestNewQuantitySuccess(t *testing.T) {
	tcs := map[string]struct {
		value float64
		unit  domain.QuantityUnit
	}{
		"postiveCentiliter": {
			value: 13.37,
			unit:  domain.QuantityUnitCentiliter,
		},
		"positiveNoUnit": {
			value: 42,
			unit:  domain.QuantityUnitNone,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			domain, err := domain.NewQuantity(tc.value, tc.unit)
			testutils.RequireNoError(t, err, "didn't expect an error")

			testutils.AssertEqualFloat64(t, tc.value, domain.Value(), "unexpected domain value")
			domaintest.AssertQuantityUnit(t, tc.unit, domain.Unit(), "unexpected domaine unit")
		})
	}
}

func TestNewQuantityError(t *testing.T) {
	tcs := map[string]struct {
		value float64
		unit  domain.QuantityUnit
	}{
		"zeroQuantity": {
			value: 0,
			unit:  domain.QuantityUnitCentiliter,
		},
		"negativeQuantity": {
			value: -5,
			unit:  domain.QuantityUnitNone,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewQuantity(tc.value, tc.unit)
			testutils.RequireHasError(t, err, "expecting an error")

			testutils.AssertErrorIs(t, domain.ErrQuantityTooSmall, err, "unexpected error")
		})
	}
}

func TestQuantityString(t *testing.T) {
	tcs := map[string]struct {
		value    float64
		unit     domain.QuantityUnit
		expected string
	}{
		"withUnit": {
			value:    15.23,
			unit:     domain.QuantityUnitKilogram,
			expected: "15.23kg",
		},
		"withoutUnit": {
			value:    12,
			unit:     domain.QuantityUnitNone,
			expected: "12",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			domain, err := domain.NewQuantity(tc.value, tc.unit)
			testutils.RequireNoError(t, err, "expecting no error")

			testutils.AssertEqualString(t, tc.expected, domain.String(), "unexpected stringer implementation")
		})
	}
}

func TestQuantityIsZeroTrue(t *testing.T) {
	var domain domain.Quantity
	testutils.AssertEqualBool(t, true, domain.IsZero(), "expeting zero value")
}

func TestQuantityIsZeroFalse(t *testing.T) {
	domain, err := domain.NewQuantity(12, domain.QuantityUnitGram)
	testutils.RequireNoError(t, err, "unexpecter error")

	testutils.AssertEqualBool(t, false, domain.IsZero(), "not expecting zero value")
}
