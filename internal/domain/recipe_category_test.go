package domain_test

import (
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestRecipeCategoryString(t *testing.T) {
	tcs := map[string]struct {
		category domain.RecipeCategory
		value    string
	}{
		"cold-meal": {
			category: domain.RecipeCategoryColdMeal,
			value:    "cold-meal",
		},
		"hot-meal": {
			category: domain.RecipeCategoryHotMeal,
			value:    "hot-meal",
		},
		"egg-dish": {
			category: domain.RecipeCategoryEggDish,
			value:    "egg-dish",
		},
		"meat-dish": {
			category: domain.RecipeCategoryMeatDish,
			value:    "meat-dish",
		},
		"fish-dish": {
			category: domain.RecipeCategoryFishDish,
			value:    "fish-dish",
		},
		"side-dish": {
			category: domain.RecipeCategorySideDish,
			value:    "side-dish",
		},
		"biscuit": {
			category: domain.RecipeCategoryBiscuit,
			value:    "biscuit",
		},
		"dessert": {
			category: domain.RecipeCategoryDessert,
			value:    "dessert",
		},

		"starter-dish": {
			category: domain.RecipeCategoryStarterDish,
			value:    "starter-dish",
		},
		"sauce": {
			category: domain.RecipeCategorySauce,
			value:    "sauce",
		},
		"spice": {
			category: domain.RecipeCategorySpice,
			value:    "spice",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualString(t, tc.value, tc.category.String(), "invalid string implementation")
		})
	}
}

func TestRecipeCategoryZeroTrue(t *testing.T) {
	var category domain.RecipeCategory
	testutils.AssertEqualBool(t, true, category.IsZero(), "expecting zero value")
}

func TestRecipeCategoryZeroFalse(t *testing.T) {
	tcs := map[string]domain.RecipeCategory{
		"cold-meal":    domain.RecipeCategoryColdMeal,
		"hot-meal":     domain.RecipeCategoryHotMeal,
		"egg-dish":     domain.RecipeCategoryEggDish,
		"meat-dish":    domain.RecipeCategoryMeatDish,
		"fish-dish":    domain.RecipeCategoryFishDish,
		"side-dish":    domain.RecipeCategorySideDish,
		"biscuit":      domain.RecipeCategoryBiscuit,
		"dessert":      domain.RecipeCategoryDessert,
		"starter-dish": domain.RecipeCategoryStarterDish,
		"sauce":        domain.RecipeCategorySauce,
		"spice":        domain.RecipeCategorySpice,
	}

	for name, unit := range tcs {
		t.Run(name, func(t *testing.T) {
			testutils.AssertEqualBool(t, false, unit.IsZero(), "expecting non zero value")
		})
	}
}
