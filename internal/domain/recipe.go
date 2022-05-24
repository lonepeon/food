package domain

type RecipeCode struct {
	Slug
}

func NewRecipeCode(s string) (RecipeCode, error) {
	slug, err := NewSlug(s)
	if err != nil {
		return RecipeCode{}, err
	}

	return RecipeCode{slug}, nil
}
