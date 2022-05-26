package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestPricingString(t *testing.T) {
	tcs := map[string]struct {
		unit  domain.Pricing
		value string
	}{
		"cheap":      {unit: domain.PricingCheap, value: "cheap"},
		"affordable": {unit: domain.PricingAffordable, value: "affordable"},
		"expensive":  {unit: domain.PricingExpensive, value: "expensive"},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualString(t, tc.value, tc.unit.String(), "invalid string implementation")
		})
	}
}

func TestPricingZeroTrue(t *testing.T) {
	tcs := map[string]domain.Pricing{
		"emptyStruct": {},
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, true, unit.IsZero(), "expecting zero value")
		})
	}
}

func TestPricingZeroFalse(t *testing.T) {
	tcs := map[string]domain.Pricing{
		"cheap":      domain.PricingCheap,
		"affordable": domain.PricingAffordable,
		"expensive":  domain.PricingExpensive,
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, false, unit.IsZero(), "expecting non zero value")
		})
	}
}
