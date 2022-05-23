package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestQuantityUnitString(t *testing.T) {
	tcs := map[string]struct {
		unit  domain.QuantityUnit
		value string
	}{
		"centiliter":  {unit: domain.QuantityUnitCentiliter, value: "cl"},
		"mililiter":   {unit: domain.QuantityUnitMililiter, value: "ml"},
		"liter":       {unit: domain.QuantityUnitLiter, value: "l"},
		"miligram":    {unit: domain.QuantityUnitMiligram, value: "mg"},
		"gram":        {unit: domain.QuantityUnitGram, value: "g"},
		"kilogram":    {unit: domain.QuantityUnitKilogram, value: "kg"},
		"tea spoon":   {unit: domain.QuantityUnitTeaSpoon, value: "cc"},
		"table spoon": {unit: domain.QuantityUnitTableSpoon, value: "cs"},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualString(t, tc.value, tc.unit.String(), "invalid string implementation")
		})
	}
}

func TestQuantityUnitZeroTrue(t *testing.T) {
	tcs := map[string]domain.QuantityUnit{
		"none":        domain.QuantityUnitNone,
		"emptyStruct": {},
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, true, unit.IsZero(), "expecting non zero value")
		})
	}
}

func TestQuantityUnitZeroFalse(t *testing.T) {
	tcs := map[string]domain.QuantityUnit{
		"centiliter":  domain.QuantityUnitCentiliter,
		"mililiter":   domain.QuantityUnitMililiter,
		"liter":       domain.QuantityUnitLiter,
		"miligram":    domain.QuantityUnitMiligram,
		"gram":        domain.QuantityUnitGram,
		"kilogram":    domain.QuantityUnitKilogram,
		"tea spoon":   domain.QuantityUnitTeaSpoon,
		"table spoon": domain.QuantityUnitTableSpoon,
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, false, unit.IsZero(), "expecting non zero value")
		})
	}
}
