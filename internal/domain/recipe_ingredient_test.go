package domain_test

import (
	"strings"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/food/internal/domain/domaintest"
	"github.com/lonepeon/golib/testutils"
)

func TestNewRecipeIngredientSuccess(t *testing.T) {
	code, err := domain.NewIngredientCode("spinach")
	testutils.RequireNoError(t, err, "cannot build ingredient code")
	quantity, err := domain.NewQuantity(12, domain.QuantityUnitGram)
	testutils.RequireNoError(t, err, "cannot build quantity")
	ingredient, err := domain.NewRecipeIngredient(code, quantity)
	testutils.RequireNoError(t, err, "cannot build recipe ingredient")

	domaintest.AssertQuantity(t, quantity, ingredient.Quantity(), "invalid quantity")
	testutils.AssertEqualString(t, code.String(), ingredient.Code().String(), "invalid ingredient code")
	details, hasDetails := ingredient.Details()
	testutils.AssertEqualBool(t, false, hasDetails, "unexpected details: %v", details)
}

func TestNewRecipeIngredientValidDetails(t *testing.T) {
	code, err := domain.NewIngredientCode("spinach")
	testutils.RequireNoError(t, err, "cannot build ingredient code")
	quantity, err := domain.NewQuantity(12, domain.QuantityUnitGram)
	testutils.RequireNoError(t, err, "cannot build quantity")
	ingredient, err := domain.NewRecipeIngredientWithDetails(code, quantity, "fresh")
	testutils.RequireNoError(t, err, "expect to build recipe ingredient")

	domaintest.AssertQuantity(t, quantity, ingredient.Quantity(), "invalid quantity")
	testutils.AssertEqualString(t, code.String(), ingredient.Code().String(), "invalid ingredient code")
	details, hasDetails := ingredient.Details()
	testutils.AssertEqualBool(t, true, hasDetails, "expecting some details")
	testutils.AssertEqualString(t, "fresh", details, "invalid details")
}

func TestNewRecipeIngredientEmptyIngredientCode(t *testing.T) {
	var code domain.IngredientCode
	quantity, err := domain.NewQuantity(12, domain.QuantityUnitGram)
	testutils.RequireNoError(t, err, "cannot build quantity")
	_, err = domain.NewRecipeIngredient(code, quantity)
	testutils.RequireHasError(t, err, "didn't expect to build recipe ingredient")
	testutils.AssertErrorIs(t, domain.ErrRecipeIngredientCodeNotInitialized, err, "invalid error type")
}

func TestNewRecipeIngredientEmptyQuantity(t *testing.T) {
	code, err := domain.NewIngredientCode("spinach")
	testutils.RequireNoError(t, err, "cannot build ingredient code")
	var quantity domain.Quantity
	_, err = domain.NewRecipeIngredient(code, quantity)
	testutils.RequireHasError(t, err, "didn't expect to build recipe ingredient")
	testutils.AssertErrorIs(t, domain.ErrRecipeIngredientQuantityNotInitialized, err, "invalid error type")
}

func TestNewRecipeIngredientTooShortDetails(t *testing.T) {
	code, err := domain.NewIngredientCode("spinach")
	testutils.RequireNoError(t, err, "cannot build ingredient code")
	quantity, err := domain.NewQuantity(12, domain.QuantityUnitGram)
	testutils.RequireNoError(t, err, "cannot build quantity")
	_, err = domain.NewRecipeIngredientWithDetails(code, quantity, "")
	testutils.RequireHasError(t, err, "didn't expect to build recipe ingredient")
	testutils.AssertErrorIs(t, domain.ErrRecipeIngredientDetailsTooSmall, err, "invalid error type")
}

func TestNewRecipeIngredientTooBigDetails(t *testing.T) {
	code, err := domain.NewIngredientCode("spinach")
	testutils.RequireNoError(t, err, "cannot build ingredient code")
	quantity, err := domain.NewQuantity(12, domain.QuantityUnitGram)
	testutils.RequireNoError(t, err, "cannot build quantity")
	_, err = domain.NewRecipeIngredientWithDetails(code, quantity, strings.Repeat("x", 105))
	testutils.RequireHasError(t, err, "didn't expect to build recipe ingredient")
	testutils.AssertErrorIs(t, domain.ErrRecipeIngredientDetailsTooBig, err, "invalid error type")
}
