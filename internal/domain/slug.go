package domain

import (
	"fmt"
	"strings"
)

var (
	ErrSlugTooSmall              = EInvalid("must be greater than 1 character")
	ErrSlugTooBig                = EInvalid("must be less than 63 characters")
	ErrSlugUnusableCharacter     = EInvalid("must only contain lower case alphanumeric characters, and - or .")
	ErrSlugInvalidFirstCharacter = EInvalid("must start with an alphanumeric character")
	ErrSlugInvalidLastCharacter  = EInvalid("must end with an alphanumeric character")
)

// Slug represents a 63 characters maximum string containing alpha-numeric
// characters and dash or dot
type Slug struct {
	slug string
}

func NewSlug(s string) (Slug, error) {
	if err := validateSlugLength(s); err != nil {
		return Slug{""}, err
	}

	if err := validateSlugCharacterSet(s); err != nil {
		return Slug{""}, err
	}

	if err := validateSlugCharacterPositions(s); err != nil {
		return Slug{""}, err
	}

	return Slug{s}, nil
}

func (s Slug) IsZero() bool {
	return string(s.slug) == ""
}

func (s Slug) String() string {
	return s.slug
}

func validateSlugLength(s string) error {
	if len(s) < 1 {
		return ErrSlugTooSmall
	}

	if len(s) > 63 {
		return ErrSlugTooBig
	}

	return nil
}

func validateSlugCharacterSet(s string) error {
	for i, c := range s {
		if err := validateSlugCharacterInCharacterSet(c); err != nil {
			return fmt.Errorf("invalid character \"%c\" at index %d: %w", c, i, err)
		}
	}

	return nil
}

func validateSlugCharacterInCharacterSet(c rune) error {
	if c >= '0' && c <= '9' {
		return nil
	}

	if c >= 'a' && c <= 'z' {
		return nil
	}

	if c == '.' || c == '-' {
		return nil
	}

	return ErrSlugUnusableCharacter
}

func validateSlugCharacterPositions(s string) error {
	if strings.Contains(s, "--") {
		return fmt.Errorf("invalid sequence \"--\": %w", ErrSlugUnusableCharacter)
	}

	if strings.Contains(s, "..") {
		return fmt.Errorf("invalid sequence \"..\": %w", ErrSlugUnusableCharacter)
	}

	if strings.HasPrefix(s, ".") || strings.HasPrefix(s, "-") {
		return ErrSlugInvalidFirstCharacter
	}

	if strings.HasSuffix(s, ".") || strings.HasSuffix(s, "-") {
		return ErrSlugInvalidLastCharacter
	}

	return nil
}
