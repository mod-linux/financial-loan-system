package loan

type AmountHold struct {
	Amount                 float64
	ChargeInPercentage     bool
	DeductFromDisbursement bool
	Percentage             float64
}

func (holdAmount AmountHold) Init(loan LoanDetails) interface{} {
	holdAmount.Amount = holdAmount.Percentage * (loan.ApprovedAmount - loan.holdAmount) / 100
	return holdAmount
}

func (holdAmount AmountHold) DeductedFromDisbursement() float64 {
	if holdAmount.DeductFromDisbursement {
		return holdAmount.Amount
	}
	return 0
}

func (holdAmount AmountHold) DeductedFromHoldAmount() float64 {
	return 0
}

func (holdAmount AmountHold) GetHoldAmount() float64 {
	return holdAmount.Amount
}

func (loanDetail LoanDetails) GetHoldAmount() float64 {
	// Initialize amount with the loan amount
	amount := 0.0

	// Iterate over each charge and subtract its amount from the total amount
	for _, charge := range loanDetail.ChargesConfiguration.Charges {
		amount += charge.Init(loanDetail).(Charges).DeductedFromHoldAmount()
	}

	// Return the remaining amount after deducting charges
	return amount
}
