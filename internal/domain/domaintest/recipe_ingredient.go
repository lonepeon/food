package domaintest

import (
	"fmt"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func AssertRecipeIngredient(t *testing.T, want domain.RecipeIngredient, got domain.RecipeIngredient, format string, args ...interface{}) {
	AssertIngredientCode(t, want.Code(), got.Code(), "invalid ingredient code: %s", fmt.Sprintf(format, args...))
	AssertQuantity(t, want.Quantity(), got.Quantity(), "invalid quantity: %s", fmt.Sprintf(format, args...))

	wantDetails, hasWantDetails := want.Details()
	gotDetails, hasGotDetails := want.Details()
	if hasWantDetails != hasGotDetails || hasWantDetails {
		testutils.AssertEqualString(t, wantDetails, gotDetails, "invalid details")
	}
}
