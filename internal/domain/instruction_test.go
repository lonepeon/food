package domain_test

import (
	"strings"
	"testing"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/food/internal/domain/domaintest"
	"github.com/lonepeon/golib/testutils"
)

func TestNewInstructionSuccess(t *testing.T) {
	tcs := map[string]string{
		"success": "Put ingredient to a bowl",
		"maxSize": strings.Repeat("x", domain.InstructionTitleMaxLength),
	}

	for name, title := range tcs {
		t.Run(name, func(t *testing.T) {
			instruction, err := domain.NewInstruction(title)
			testutils.RequireNoError(t, err, "failed to build instruction")
			testutils.AssertEqualString(t, title, instruction.Title(), "invalid title")
		})
	}
}

func TestNewInstructionInvalidTitle(t *testing.T) {
	tcs := map[string]struct {
		title string
		err   error
	}{
		"empty": {
			title: "",
			err:   domain.ErrInstructionTitleTooSmall,
		},
		"startWithSpaces": {
			title: "   Put ingredient to a bowl",
			err:   domain.ErrInstructionTitleInvalidCharacterSet,
		},
		"endsWithSpaces": {
			title: "Put ingredient to a bowl   ",
			err:   domain.ErrInstructionTitleInvalidCharacterSet,
		},
		"onlySpaces": {
			title: "   ",
			err:   domain.ErrInstructionTitleInvalidCharacterSet,
		},
		"tooBig": {
			title: strings.Repeat("x", domain.InstructionTitleMaxLength+1),
			err:   domain.ErrInstructionTitleTooBig,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewInstruction(tc.title)
			testutils.RequireHasError(t, err, "expecting failure on new instruction")
			testutils.AssertErrorIs(t, tc.err, err, "invalid new instruction failure")
		})
	}
}

func TestInstructionLinkedToRecipeSuccess(t *testing.T) {
	instruction, err := domain.NewInstruction("Get the pizza dough")
	testutils.RequireNoError(t, err, "failed to build new instruction")

	recipeCode, err := domain.NewRecipeCode("pizza-dough")
	testutils.RequireNoError(t, err, "failed to build recipe code")

	instruction, err = instruction.LinkRecipe(recipeCode)
	testutils.RequireNoError(t, err, "failed to link recipe to instruction")

	instructionRecipeCode, isSet := instruction.RecipeCode()
	testutils.RequireEqualBool(t, true, isSet, "invalid recipe code flag")
	domaintest.AssertRecipeCode(t, recipeCode, instructionRecipeCode, "invalid recipe code")
}

func TestInstructionLinkedToRecipeZeroCode(t *testing.T) {
	instruction, err := domain.NewInstruction("Get the pizza dough")
	testutils.RequireNoError(t, err, "failed to build new instruction")

	var recipeCode domain.RecipeCode

	_, err = instruction.LinkRecipe(recipeCode)
	testutils.RequireHasError(t, err, "expecting to fail to link recipe to instruction")

	testutils.AssertErrorIs(t, domain.ErrRecipeCodeNotInitialized, err, "invalid error")
}

func TestHighlightTextSuccess(t *testing.T) {
	instruction, err := domain.NewInstruction("Get the pizza dough")
	testutils.RequireNoError(t, err, "failed to build new instruction")

	text := "and let it rest for 30m"
	instruction, err = instruction.HighlightText(text)
	testutils.RequireNoError(t, err, "failed to add highlighted text")

	highlightedText, isSet := instruction.HighlightedText()
	testutils.RequireEqualBool(t, true, isSet, "expected to have highlighted text")
	testutils.AssertEqualString(t, text, highlightedText, "invalid highlighted text")
}

func TestHighlightTextError(t *testing.T) {
	tcs := map[string]struct {
		text string
		err  error
	}{
		"empty": {
			text: "",
			err:  domain.ErrInstructionHighlightTooSmall,
		},
		"startWithSpaces": {
			text: "   Put ingredient to a bowl",
			err:  domain.ErrInstructionHighlightInvalidCharacterSet,
		},
		"endsWithSpaces": {
			text: "Put ingredient to a bowl   ",
			err:  domain.ErrInstructionHighlightInvalidCharacterSet,
		},
		"onlySpaces": {
			text: "   ",
			err:  domain.ErrInstructionHighlightInvalidCharacterSet,
		},
		"tooBig": {
			text: strings.Repeat("x", domain.InstructionHighlightMaxLength+1),
			err:  domain.ErrInstructionHighlightTooBig,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			instruction, err := domain.NewInstruction("Get the pizza dough")
			testutils.RequireNoError(t, err, "failed to build new instruction")

			_, err = instruction.HighlightText(tc.text)
			testutils.RequireHasError(t, err, "expected to fail to add highlighted text")

			testutils.AssertErrorIs(t, tc.err, err, "invalid error")
		})
	}
}
