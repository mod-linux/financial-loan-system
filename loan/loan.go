package loan

import (
	Charges "linus/lms/charges"
	"linus/lms/finance"
)

type LoanDetails struct {
	LoanAmount           float64
	Tenure               float64
	InterestType         int
	PaymentFrequency     int
	InterestRate         float64
	monthlyInterestRate  float64
	totalInterest        float64
	ChargesConfiguration ChargesConfiguration
}

type ChargesConfiguration struct {
	Charges []Charges.Charges
}

func (loanDetail LoanDetails) getNetLoanAmount() float64 {
	// Initialize amount with the loan amount
	amount := loanDetail.LoanAmount

	// Get the charge's configuration from the loan details
	charges := loanDetail.ChargesConfiguration.Charges

	// Iterate over each charge and subtract its amount from the total amount
	for _, charge := range charges {
		amount -= charge.GetAmountToBeDeductedFromLoanAmount()
	}

	// Return the remaining amount after deducting charges
	return amount
}

func (loanDetail LoanDetails) Init() LoanDetails {
	loanDetail.monthlyInterestRate = finance.ReducingInterestRate(loanDetail.InterestRate, loanDetail.getNetLoanAmount(), loanDetail.Tenure, loanDetail.InterestType, loanDetail.PaymentFrequency)
	return loanDetail
}

func (loanDetail LoanDetails) GetEMI() float64 {
	return finance.CalculatePMT(loanDetail.LoanAmount, loanDetail.monthlyInterestRate, loanDetail.Tenure)
}

func (loanDetail LoanDetails) GetTotalInterest() float64 {
	return finance.GetTotalInterestCharged(loanDetail.InterestRate, loanDetail.getNetLoanAmount())
}

func (loanDetail LoanDetails) GetReducingInterestRate() float64 {
	return loanDetail.monthlyInterestRate
}
