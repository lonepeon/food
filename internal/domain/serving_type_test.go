package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestServingTypeString(t *testing.T) {
	tcs := map[string]struct {
		unit  domain.ServingType
		value string
	}{
		"guest": {unit: domain.ServingTypeGuest, value: "guest"},
		"unit":  {unit: domain.ServingTypeUnit, value: "unit"},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualString(t, tc.value, tc.unit.String(), "invalid string implementation")
		})
	}
}

func TestServingTypeZeroTrue(t *testing.T) {
	var pricing domain.ServingType
	testutils.AssertEqualBool(t, true, pricing.IsZero(), "expecting zero value")
}

func TestServingTypeZeroFalse(t *testing.T) {
	tcs := map[string]domain.ServingType{
		"guest": domain.ServingTypeGuest,
		"unit":  domain.ServingTypeUnit,
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, false, unit.IsZero(), "expecting non zero value")
		})
	}
}
