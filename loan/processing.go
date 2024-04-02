package loan

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

func (processingFeeCharges ProcessingFeeCharges) Init(loan LoanDetails) interface{} {
	return processingFeeCharges
}

func (processingFeeCharges ProcessingFeeCharges) GetProcessingFeeAmount() float64 {
	return processingFeeCharges.Amount
}

func (processingFeeCharges ProcessingFeeCharges) GetAmountToBeDeductedFromLoanAmount(loan LoanDetails) float64 {
	if !processingFeeCharges.LSPCollected {
		return processingFeeCharges.Amount
	}
	return 0
}
