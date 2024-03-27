package loan

import (
	"fmt"
	Charges "linus/lms/charges"
	"linus/lms/constants"
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

func (loanDetail LoanDetails) deductFromAmount() float64 {
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
	monthlyInterestRate := finance.ConvertInterestRateBasedOnFrequency(loanDetail.InterestRate, loanDetail.PaymentFrequency)

	if loanDetail.InterestType == constants.Flat {
		// Only calculate the reducing interest rate if the interest type is reducing
		monthlyInterestRate = finance.ToReducing(loanDetail.InterestRate, loanDetail.Tenure, loanDetail.PaymentFrequency, loanDetail.deductFromAmount())
	}

	loanDetail.monthlyInterestRate = monthlyInterestRate
	fmt.Println(loanDetail.monthlyInterestRate)

	return loanDetail
}

func (loanDetail LoanDetails) GetEMI() float64 {
	return finance.CalculatePMT(loanDetail.LoanAmount, loanDetail.monthlyInterestRate, loanDetail.Tenure)
}
