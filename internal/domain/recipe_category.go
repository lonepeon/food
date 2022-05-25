package domain

type RecipeCategory struct {
	code string
}

var (
	RecipeCategoryColdMeal    = RecipeCategory{"cold-meal"}
	RecipeCategoryHotMeal     = RecipeCategory{"hot-meal"}
	RecipeCategoryEggDish     = RecipeCategory{"egg-dish"}
	RecipeCategoryMeatDish    = RecipeCategory{"meat-dish"}
	RecipeCategoryFishDish    = RecipeCategory{"fish-dish"}
	RecipeCategorySideDish    = RecipeCategory{"side-dish"}
	RecipeCategoryBiscuit     = RecipeCategory{"biscuit"}
	RecipeCategoryDessert     = RecipeCategory{"dessert"}
	RecipeCategoryStarterDish = RecipeCategory{"starter-dish"}
	RecipeCategorySauce       = RecipeCategory{"sauce"}
	RecipeCategorySpice       = RecipeCategory{"spice"}
)

func (c RecipeCategory) String() string {
	return c.code
}

func (c RecipeCategory) IsZero() bool {
	return c.code == ""
}
