package domaintest

import (
	"fmt"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func AssertIngredient(t *testing.T, want *domain.Ingredient, got *domain.Ingredient, format string, args ...interface{}) {
	AssertIngredientCode(t, want.Code(), got.Code(), "invalid code: %s", fmt.Sprintf(format, args...))
	testutils.AssertEqualString(t, want.Name(), got.Name(), "invalid name: %s", fmt.Sprintf(format, args...))
}

func AssertIngredientCode(t *testing.T, want domain.IngredientCode, got domain.IngredientCode, format string, args ...interface{}) {
	testutils.AssertEqualString(t, want.String(), got.String(), format, args...)
}
