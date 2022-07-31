package domaintest

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func AssertRecipeCode(t *testing.T, want domain.RecipeCode, got domain.RecipeCode, format string, args ...interface{}) {
	testutils.AssertEqualString(t, want.String(), got.String(), format, args...)
}
