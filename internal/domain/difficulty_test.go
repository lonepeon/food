package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestDifficultyString(t *testing.T) {
	tcs := map[string]struct {
		unit  domain.Difficulty
		value string
	}{
		"easy":    {unit: domain.DifficultyEasy, value: "easy"},
		"average": {unit: domain.DifficultyAverage, value: "average"},
		"hard":    {unit: domain.DifficultyHard, value: "hard"},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualString(t, tc.value, tc.unit.String(), "invalid string implementation")
		})
	}
}

func TestDifficultyZeroTrue(t *testing.T) {
	var difficulty domain.Difficulty
	testutils.AssertEqualBool(t, true, difficulty.IsZero(), "expecting zero value")
}

func TestDifficultyZeroFalse(t *testing.T) {
	tcs := map[string]domain.Difficulty{
		"easy":    domain.DifficultyEasy,
		"average": domain.DifficultyAverage,
		"hard":    domain.DifficultyHard,
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, false, unit.IsZero(), "expecting non zero value")
		})
	}
}
