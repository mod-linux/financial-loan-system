package charges

type SubventionFeeCharges struct {
	Amount                     float64
	ChargeOnNetAmount          bool
	ChargeOnPrincipalRemaining bool
	ChargeInPercentage         bool
	LSPCollected               bool
	GSTApplicable              bool
	TAXES                      GSTCharges
}

func (subventionFeeCharges SubventionFeeCharges) Calculate() float64 {
	return subventionFeeCharges.Amount
}

func (subventionFeeCharges SubventionFeeCharges) GetAmountToBeDeductedFromLoanAmount() float64 {
	return 0
}
