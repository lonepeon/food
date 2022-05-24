package domain_test

import (
	"strings"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestNewRecipeCodeSuccess(t *testing.T) {
	code, err := domain.NewRecipeCode("pasta-pesto")
	testutils.RequireNoError(t, err, "expecting no code error")
	testutils.AssertEqualString(t, "pasta-pesto", code.String(), "expecting another code value")
}

func TestNewRecipeCodeError(t *testing.T) {
	codes := map[string]string{
		"invalidCharacters": "pasta pesto",
		"empty":             "",
		"tooLong":           strings.Repeat("x", 65),
	}

	for name, rawCode := range codes {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewRecipeCode(rawCode)
			testutils.AssertHasError(t, err, "expecting code error")
		})
	}
}
