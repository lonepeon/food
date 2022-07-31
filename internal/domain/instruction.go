package domain

import (
	"fmt"
	"strings"
)

const (
	InstructionTitleMaxLength     = 255
	InstructionHighlightMaxLength = 255
)

var (
	ErrInstructionTitleInvalidCharacterSet     = EInvalid("title must not start or end with spaces")
	ErrInstructionTitleTooSmall                = EInvalid("title must be greater than 1 character")
	ErrInstructionTitleTooBig                  = EInvalid(fmt.Sprintf("title must be less than %d characters", InstructionTitleMaxLength))
	ErrInstructionHighlightInvalidCharacterSet = EInvalid("highlight must not start or end with spaces")
	ErrInstructionHighlightTooSmall            = EInvalid("highlight must be greater than 1 character")
	ErrInstructionHighlightTooBig              = EInvalid(fmt.Sprintf("highlight must be less than %d characters", InstructionTitleMaxLength))
)

type Instruction struct {
	title       string
	recipeCode  RecipeCode
	highlighted string
}

func NewInstruction(title string) (Instruction, error) {
	if err := validateInstructionTitleLength(title); err != nil {
		return Instruction{}, err
	}

	return Instruction{title: title}, nil
}

func (i Instruction) HighlightedText() (string, bool) {
	return i.highlighted, len(i.highlighted) > 0
}

func (i Instruction) Title() string {
	return i.title
}

func (i Instruction) RecipeCode() (RecipeCode, bool) {
	return i.recipeCode, !i.recipeCode.IsZero()
}

func (i Instruction) LinkRecipe(code RecipeCode) (Instruction, error) {
	if code.IsZero() {
		return Instruction{}, ErrRecipeCodeNotInitialized
	}

	i.recipeCode = code
	return i, nil
}

func (i Instruction) HighlightText(txt string) (Instruction, error) {
	if err := validateInstructionHighlightLength(txt); err != nil {
		return Instruction{}, err
	}

	i.highlighted = txt
	return i, nil
}

func validateInstructionTitleLength(title string) error {
	if strings.TrimSpace(title) != title {
		return ErrInstructionTitleInvalidCharacterSet
	}

	if len(title) < 1 {
		return ErrInstructionTitleTooSmall
	}

	if len(title) > InstructionTitleMaxLength {
		return ErrInstructionTitleTooBig
	}

	return nil
}

func validateInstructionHighlightLength(highlight string) error {
	if strings.TrimSpace(highlight) != highlight {
		return ErrInstructionHighlightInvalidCharacterSet
	}

	if len(highlight) < 1 {
		return ErrInstructionHighlightTooSmall
	}

	if len(highlight) > InstructionHighlightMaxLength {
		return ErrInstructionHighlightTooBig
	}

	return nil
}
