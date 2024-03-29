package charges

type PenalCharges struct {
	Amount             float32
	ChargeInPercentage float32
	Frequency          int
	Days               int
	GSTApplicable      bool
	TAXES              GSTCharges
}

func (penalCharges PenalCharges) Calculate() float32 {
	return penalCharges.Amount * float32(penalCharges.Days)
}
