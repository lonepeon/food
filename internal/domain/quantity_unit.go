package domain

type QuantityUnit struct {
	code string
}

var (
	QuantityUnitNone = QuantityUnit{""}

	QuantityUnitMililiter  = QuantityUnit{"ml"}
	QuantityUnitCentiliter = QuantityUnit{"cl"}
	QuantityUnitLiter      = QuantityUnit{"l"}

	QuantityUnitMiligram = QuantityUnit{"mg"}
	QuantityUnitGram     = QuantityUnit{"g"}
	QuantityUnitKilogram = QuantityUnit{"kg"}

	QuantityUnitTeaSpoon   = QuantityUnit{"cc"}
	QuantityUnitTableSpoon = QuantityUnit{"cs"}
)

func (u QuantityUnit) String() string {
	return u.code
}

func (u QuantityUnit) IsZero() bool {
	return u.code == ""
}
