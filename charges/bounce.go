package charges

type BounceCharges struct {
	Amount             float32
	ChargeInPercentage float32
	Frequency          int
	TAXES              GSTCharges
}

type Result struct {
	Days   int
	Amount float32
}

func (bounceCharges BounceCharges) Calculate() float32 {
	return bounceCharges.Amount
}
