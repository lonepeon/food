package domaintest

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func AssertQuantityUnit(t *testing.T, want domain.QuantityUnit, got domain.QuantityUnit, format string, args ...interface{}) {
	testutils.AssertEqualString(t, want.String(), got.String(), format, args...)
}
