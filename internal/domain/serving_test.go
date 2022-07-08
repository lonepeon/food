package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/food/internal/domain/domaintest"
	"github.com/lonepeon/golib/testutils"
)

func TestNewGuestServingSuccess(t *testing.T) {
	serving, err := domain.NewGuestServing(5)
	testutils.RequireNoError(t, err, "can't create guest serving")
	testutils.AssertEqualInt(t, 5, serving.Number(), "invalid number of guests")
	domaintest.AssertServingType(t, domain.ServingTypeGuest, serving.Type(), "invalid serving type")
}

func TestNewGuestServingLessThanZero(t *testing.T) {
	_, err := domain.NewGuestServing(-1)
	testutils.RequireHasError(t, err, "expecting an error")
	testutils.AssertErrorIs(t, domain.ErrServingTooSmall, err, "invalid error type")
}

func TestNewUnitServingSuccess(t *testing.T) {
	serving, err := domain.NewUnitServing(5)
	testutils.RequireNoError(t, err, "can't create unit serving")
	testutils.AssertEqualInt(t, 5, serving.Number(), "invalid number of units")
	domaintest.AssertServingType(t, domain.ServingTypeUnit, serving.Type(), "invalid serving type")
}

func TestNewUnitServingLessThanZero(t *testing.T) {
	_, err := domain.NewUnitServing(-1)
	testutils.RequireHasError(t, err, "expecting an error")
	testutils.AssertErrorIs(t, domain.ErrServingTooSmall, err, "invalid error type")
}

func TestServingString(t *testing.T) {
	tcs := map[string]struct {
		number      int
		constructor func(int) (domain.Serving, error)
		expected    string
	}{
		"unit": {
			number:      1,
			constructor: domain.NewUnitServing,
			expected:    "1 unit",
		},
		"units": {
			number:      12,
			constructor: domain.NewUnitServing,
			expected:    "12 units",
		},
		"guest": {
			number:      1,
			constructor: domain.NewGuestServing,
			expected:    "1 guest",
		},
		"guests": {
			number:      12,
			constructor: domain.NewGuestServing,
			expected:    "12 guests",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			serving, err := tc.constructor(tc.number)
			testutils.RequireNoError(t, err, "can't build serving")
			testutils.AssertEqualString(t, tc.expected, serving.String(), "invalid string representation")

		})
	}
}

func TestServingZeroTrue(t *testing.T) {
	var serving domain.Serving
	testutils.AssertEqualBool(t, true, serving.IsZero(), "expecting zero value")
}

func TestServingZeroFalse(t *testing.T) {
	tcs := map[string]func(int) (domain.Serving, error){
		"guest": domain.NewGuestServing,
		"unit":  domain.NewUnitServing,
	}

	for name, constructor := range tcs {
		t.Run(name, func(t *testing.T) {
			serving, err := constructor(12)
			testutils.RequireNoError(t, err, "can't build serving")
			testutils.AssertEqualBool(t, false, serving.IsZero(), "expecting non zero value")
		})
	}
}
