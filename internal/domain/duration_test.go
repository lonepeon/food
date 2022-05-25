package domain_test

import (
	"testing"
	"time"

	"github.com/lonepeon/food/internal/domain"
	"github.com/lonepeon/golib/testutils"
)

func TestRecipeDurationFields(t *testing.T) {
	tcs := map[string]struct {
		preparation time.Duration
		cooking     time.Duration
		resting     time.Duration
	}{
		"allSet": {
			preparation: 25 * time.Minute,
			cooking:     15 * time.Minute,
			resting:     40 * time.Minute,
		},
		"preparationNotSet": {
			cooking: 15 * time.Minute,
			resting: 40 * time.Minute,
		},
		"cookingNotSet": {
			preparation: 25 * time.Minute,
			resting:     40 * time.Minute,
		},
		"restingNotSet": {
			preparation: 25 * time.Minute,
			cooking:     15 * time.Minute,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			duration := domain.NewRecipeDuration(tc.preparation, tc.cooking, tc.resting)
			testutils.AssertEqualDuration(t, tc.preparation, duration.Preparation(), "invalid preparation")
			testutils.AssertEqualDuration(t, tc.cooking, duration.Cooking(), "invalid cooking")
			testutils.AssertEqualDuration(t, tc.resting, duration.Resting(), "invalid resting")
		})
	}
}

func TestRecipeDurationTotal(t *testing.T) {
	tcs := map[string]struct {
		preparation time.Duration
		cooking     time.Duration
		resting     time.Duration
		expected    time.Duration
	}{
		"allSet": {
			preparation: 25 * time.Minute,
			cooking:     15 * time.Minute,
			resting:     40 * time.Minute,
			expected:    1*time.Hour + 20*time.Minute,
		},
		"preparationNotSet": {
			cooking:  15 * time.Minute,
			resting:  40 * time.Minute,
			expected: 55 * time.Minute,
		},
		"cookingNotSet": {
			preparation: 25 * time.Minute,
			resting:     40 * time.Minute,
			expected:    time.Hour + 5*time.Minute,
		},
		"restingNotSet": {
			preparation: 25 * time.Minute,
			cooking:     15 * time.Minute,
			expected:    40 * time.Minute,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			duration := domain.NewRecipeDuration(tc.preparation, tc.cooking, tc.resting)
			testutils.AssertEqualDuration(t, tc.expected, duration.Total(), "invalid total duration")
		})
	}
}

func TestRecipeDurationString(t *testing.T) {
	tcs := map[string]struct {
		preparation time.Duration
		cooking     time.Duration
		resting     time.Duration
		expected    string
	}{
		"allSet": {
			preparation: 25 * time.Minute,
			cooking:     15 * time.Minute,
			resting:     40 * time.Minute,
			expected:    "1h20m0s",
		},
		"preparationNotSet": {
			cooking:  15 * time.Minute,
			resting:  40 * time.Minute,
			expected: "55m0s",
		},
		"cookingNotSet": {
			preparation: 25 * time.Minute,
			resting:     40 * time.Minute,
			expected:    "1h5m0s",
		},
		"restingNotSet": {
			preparation: 25 * time.Minute,
			cooking:     15 * time.Minute,
			expected:    "40m0s",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			duration := domain.NewRecipeDuration(tc.preparation, tc.cooking, tc.resting)
			testutils.AssertEqualString(t, tc.expected, duration.String(), "invalid string implementation")
		})
	}
}
