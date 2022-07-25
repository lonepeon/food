package domain

import (
	"fmt"
	"strings"
)

var (
	ErrRecipeIngredientCodeNotInitialized         = EInternal("code must not be zero")
	ErrRecipeIngredientQuantityNotInitialized     = EInternal("quantity must not be zero")
	ErrRecipeIngredientDetailsTooSmall            = EInvalid("details must be greater than 0")
	ErrRecipeIngredientDetailsTooBig              = EInvalid("details must be less than 100")
	ErrRecipeIngredientDetailsInvalidCharacterSet = EInvalid("details must not start or end with spaces")
)

type RecipeIngredient struct {
	code     IngredientCode
	quantity Quantity
	details  string
}

func NewRecipeIngredient(code IngredientCode, quantity Quantity) (RecipeIngredient, error) {
	if code.IsZero() {
		return RecipeIngredient{}, ErrRecipeIngredientCodeNotInitialized
	}

	if quantity.IsZero() {
		return RecipeIngredient{}, ErrRecipeIngredientQuantityNotInitialized
	}

	return RecipeIngredient{code: code, quantity: quantity}, nil
}

func NewRecipeIngredientWithDetails(code IngredientCode, quantity Quantity, details string) (RecipeIngredient, error) {
	recipeIngredient, err := NewRecipeIngredient(code, quantity)
	if err != nil {
		return RecipeIngredient{}, err
	}

	if details != strings.TrimSpace(details) {
		return RecipeIngredient{}, ErrRecipeIngredientDetailsInvalidCharacterSet
	}

	if len(details) == 0 {
		return RecipeIngredient{}, ErrRecipeIngredientDetailsTooSmall
	}

	if len(details) > 100 {
		return RecipeIngredient{}, ErrRecipeIngredientDetailsTooBig
	}

	recipeIngredient.details = details

	return recipeIngredient, nil
}

func (r RecipeIngredient) Code() IngredientCode {
	return r.code
}

func (r RecipeIngredient) Quantity() Quantity {
	return r.quantity
}

func (r RecipeIngredient) Details() (string, bool) {
	return r.details, r.details != ""
}

func (r RecipeIngredient) IsZero() bool {
	return r.code.IsZero() || r.quantity.IsZero()
}

func (r RecipeIngredient) String() string {
	return fmt.Sprintf("%s(%s)", r.code, r.quantity)
}
