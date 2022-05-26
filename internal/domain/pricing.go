package domain

type Pricing struct {
	code string
}

var (
	PricingCheap      = Pricing{"cheap"}
	PricingAffordable = Pricing{"affordable"}
	PricingExpensive  = Pricing{"expensive"}
)

func (p Pricing) IsZero() bool {
	return p.code == ""
}

func (p Pricing) String() string {
	return p.code
}
