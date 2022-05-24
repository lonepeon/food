package domain

import "time"

type RecipeDuration struct {
	preparation time.Duration
	cooking     time.Duration
	resting     time.Duration
}

func NewRecipeDuration(prep, cooking, resting time.Duration) RecipeDuration {
	return RecipeDuration{preparation: prep, cooking: cooking, resting: resting}
}

func (d RecipeDuration) Preparation() time.Duration {
	return d.preparation
}

func (d RecipeDuration) Cooking() time.Duration {
	return d.cooking
}

func (d RecipeDuration) Resting() time.Duration {
	return d.resting
}

func (d RecipeDuration) TotalDuration() time.Duration {
	return d.preparation + d.cooking + d.resting
}

func (d RecipeDuration) String() string {
	return d.TotalDuration().String()
}
