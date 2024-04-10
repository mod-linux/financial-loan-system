package loan

import (
	"linus/lms/finance"
)

type LoanDetails struct {
	LoanAmount           float64
	ApprovedAmount       float64
	Tenure               float64
	InterestType         int
	PaymentFrequency     int
	InterestRate         float64
	monthlyInterestRate  float64
	holdAmount           float64
	totalInterest        float64
	subventionAmount     float64
	ChargesConfiguration ChargesConfiguration
	AmountHold           AmountHold
}

type ChargesConfiguration struct {
	Charges []Charges
}

func (loanDetail LoanDetails) GetNetLoanAmount() float64 {
	// Initialize net amount with the loan amount
	netAmount := loanDetail.LoanAmount

	// Iterate over each charge and subtract its amount from the net amount
	for _, charge := range loanDetail.ChargesConfiguration.Charges {
		netAmount -= charge.Init(loanDetail).(Charges).DeductedFromDisbursement()
	}

	// Return the remaining net amount after deducting charges
	return netAmount
}

func (loanDetail LoanDetails) Init() LoanDetails {
	loanDetail.monthlyInterestRate = finance.ReducingInterestRate(loanDetail.InterestRate, loanDetail.GetNetLoanAmount(), loanDetail.Tenure, loanDetail.InterestType, loanDetail.PaymentFrequency)
	loanDetail.holdAmount = loanDetail.GetHoldAmount()
	return loanDetail
}

func (loanDetail LoanDetails) GetEMI() float64 {
	return finance.CalculatePMT(loanDetail.LoanAmount, loanDetail.monthlyInterestRate, loanDetail.Tenure)
}

func (loanDetail LoanDetails) GetTotalInterest() float64 {
	return finance.GetTotalInterestCharged(loanDetail.InterestRate, loanDetail.GetNetLoanAmount())
}

func (loanDetail LoanDetails) GetReducingInterestRate() float64 {
	return loanDetail.monthlyInterestRate
}

func (loanDetail LoanDetails) GetLoanSchedule() []finance.Schedule {
	return finance.GetLoanSchedule(loanDetail.Tenure, loanDetail.GetEMI(), loanDetail.ApprovedAmount, loanDetail.monthlyInterestRate)
}
