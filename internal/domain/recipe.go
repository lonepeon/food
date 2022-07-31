package domain

var (
	ErrRecipeCodeNotInitialized = EInternal("code must not be zero")
)

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
