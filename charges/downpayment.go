package charges

type DownPayment struct {
	Amount               float64
	ChargeInPercentage   float32
	Frequency            int
	LSPCollected         bool
	PaidWithFirstEMI     bool
	DeductFromLoanAmount bool
}

func (downPayment DownPayment) Calculate() float64 {
	if downPayment.DeductFromLoanAmount {
		return downPayment.Amount
	}
	return 0
}

func (downPayment DownPayment) GetAmountToBeDeductedFromLoanAmount() float64 {
	if downPayment.DeductFromLoanAmount {
		return downPayment.Amount
	}
	return 0
}
