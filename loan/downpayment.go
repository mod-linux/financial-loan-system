package loan

type DownPayment struct {
	Amount               float64
	ChargeInPercentage   float32
	Frequency            int
	LSPCollected         bool
	PaidWithFirstEMI     bool
	DeductFromLoanAmount bool
}

func (downPayment DownPayment) Init(loan LoanDetails) interface{} {
	return downPayment
}

func (downPayment DownPayment) GetDownPayment() float64 {
	return downPayment.Amount
}

func (downPayment DownPayment) GetAmountToBeDeductedFromLoanAmount(loan LoanDetails) float64 {
	if downPayment.DeductFromLoanAmount {
		return downPayment.Amount
	}
	return 0
}
