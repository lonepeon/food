package domain

type ServingType struct {
	code string
}

var (
	ServingTypeGuest = ServingType{"guest"}
	ServingTypeUnit  = ServingType{"unit"}
)

func (t ServingType) IsZero() bool {
	return t.code == ""
}

func (t ServingType) String() string {
	return t.code
}
