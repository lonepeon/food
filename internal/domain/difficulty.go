package domain

type Difficulty struct {
	code string
}

var (
	DifficultyEasy    = Difficulty{"easy"}
	DifficultyAverage = Difficulty{"average"}
	DifficultyHard    = Difficulty{"hard"}
)

func (d Difficulty) IsZero() bool {
	return d.code == ""
}

func (d Difficulty) String() string {
	return d.code
}
