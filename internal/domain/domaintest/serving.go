package domaintest

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func AssertServingType(t *testing.T, want domain.ServingType, got domain.ServingType, format string, args ...interface{}) {
	testutils.AssertEqualString(t, want.String(), got.String(), format, args...)
}
