package domaintest

import (
	"fmt"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func AssertQuantityUnit(t *testing.T, want domain.QuantityUnit, got domain.QuantityUnit, format string, args ...interface{}) {
	testutils.AssertEqualString(t, want.String(), got.String(), format, args...)
}

func AssertQuantity(t *testing.T, want domain.Quantity, got domain.Quantity, format string, args ...interface{}) {
	testutils.AssertEqualFloat64(t, want.Value(), got.Value(), fmt.Sprintf("invalid value: %s", fmt.Sprintf(format, args...)))
	testutils.AssertEqualString(t, want.Unit().String(), got.Unit().String(), fmt.Sprintf("invalid unit: %s", fmt.Sprintf(format, args...)))
}
