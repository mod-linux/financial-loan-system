package charges

type ProcessingFeeCharges struct {
	Amount                     float64
	ChargeOnNetAmount          float32
	ChargeOnPrincipalRemaining bool
	ChargeInPercentage         float32
	Frequency                  int
	LSPCollected               bool
	PaidWithFirstEMI           bool
	GSTApplicable              bool
	TAXES                      GSTCharges
}

func (processingFeeCharges ProcessingFeeCharges) Calculate() float64 {
	return processingFeeCharges.Amount
}

func (processingFeeCharges ProcessingFeeCharges) GetAmountToBeDeductedFromLoanAmount() float64 {
	return 0
}
