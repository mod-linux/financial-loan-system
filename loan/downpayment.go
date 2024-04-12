package loan

type DownPayment struct {
	Amount                 float64
	ChargeInPercentage     float32
	Frequency              int
	LSPCollected           bool
	DeductFromDisbursement bool
	PaidWithFirstEMI       bool
	DeductFromLoanAmount   bool
}

func (downPayment DownPayment) Init(loan LoanDetails) interface{} {
	return downPayment
}

func (downPayment DownPayment) GetDownPayment() float64 {
	return downPayment.Amount
}

func (downPayment DownPayment) DeductedFromDisbursement() float64 {
	if downPayment.DeductFromLoanAmount && !downPayment.LSPCollected {
		return downPayment.Amount
	}
	return 0
}

func (downPayment DownPayment) DeductedFromHoldAmount() float64 {
	return 0
}
