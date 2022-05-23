package domain

import (
	"errors"
	"strings"
)

type IngredientCode struct {
	Slug
}

func NewIngredientCode(s string) (IngredientCode, error) {
	slug, err := NewSlug(s)
	if err != nil {
		return IngredientCode{}, err
	}

	return IngredientCode{slug}, nil
}

type Ingredient struct {
	code IngredientCode
	name string
}

func NewIngredient(code IngredientCode, name string) (*Ingredient, error) {
	if code.IsZero() {
		return nil, errors.New("code must not be empty")
	}

	name, err := validateIngredientNameLength(name)
	if err != nil {
		return nil, err
	}

	return &Ingredient{code: code, name: name}, nil
}

func (i *Ingredient) Code() IngredientCode {
	return i.code
}

func (i *Ingredient) Name() string {
	return i.name
}

func (i *Ingredient) String() string {
	return i.code.String()
}

func validateIngredientNameLength(name string) (string, error) {
	name = strings.TrimSpace(name)

	if len(name) < 1 {
		return name, errors.New("name must be greater than 1 character")
	}

	if len(name) > 100 {
		return name, errors.New("name must be less than 100 characters")
	}

	return name, nil
}
