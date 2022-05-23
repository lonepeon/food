package domain_test

import (
	"strings"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/food/internal/domain/domaintest"
	"github.com/lonepeon/golib/testutils"
)

func TestNewIngredientCodeSuccess(t *testing.T) {
	code, err := domain.NewIngredientCode("tomato")
	testutils.RequireNoError(t, err, "expecting no code error")
	testutils.AssertEqualString(t, "tomato", code.String(), "expecting another code value")
}

func TestNewIngredientCodeError(t *testing.T) {
	codes := map[string]string{
		"invalidCharacters": "t o m a t o",
		"empty":             "",
		"tooLong":           strings.Repeat("x", 65),
	}

	for name, rawCode := range codes {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewIngredientCode(rawCode)
			testutils.AssertHasError(t, err, "expecting code error")
		})
	}
}

func TestNewIngredientSuccess(t *testing.T) {
	code, err := domain.NewIngredientCode("black-tomato")
	testutils.RequireNoError(t, err, "expecting valid code")

	ingredient, err := domain.NewIngredient(code, "Black tomato")
	testutils.RequireNoError(t, err, "expecting valid ingredient")

	domaintest.AssertIngredientCode(t, code, ingredient.Code(), "invalid ingredient code")
	testutils.AssertEqualString(t, "Black tomato", ingredient.Name(), "invalid ingredient name")
}

func TestNewIngredientSlugZeroError(t *testing.T) {
	_, err := domain.NewIngredient(domain.IngredientCode{}, "Black tomato")
	testutils.RequireHasError(t, err, "expecting error")

	testutils.AssertEqualString(t, "code must not be empty", err.Error(), "invalid error message")
}

func TestNewIngredientNameError(t *testing.T) {
	code, err := domain.NewIngredientCode("black-tomato")
	testutils.RequireNoError(t, err, "expecting valid code")

	names := map[string]struct {
		name          string
		expectedError string
	}{
		"empty": {
			name:          "",
			expectedError: "name must be greater than 1 character",
		},
		"tooLong": {
			name:          strings.Repeat("x", 105),
			expectedError: "name must be less than 100 characters",
		},
		"onlySpaces": {
			name:          strings.Repeat(" ", 20),
			expectedError: "name must be greater than 1 character",
		},
	}

	for name, tc := range names {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewIngredient(code, tc.name)
			testutils.RequireHasError(t, err, "expecting error")
			testutils.AssertEqualString(t, tc.expectedError, err.Error(), "invalid error message")
		})
	}
}

func TestIngredientString(t *testing.T) {
	code, err := domain.NewIngredientCode("black-tomato")
	testutils.RequireNoError(t, err, "expecting valid code")

	ingredient, err := domain.NewIngredient(code, "Black tomato")
	testutils.RequireNoError(t, err, "expecting valid ingredient")

	testutils.AssertEqualString(t, "black-tomato", ingredient.String(), "invalid stringer implementation")
}
