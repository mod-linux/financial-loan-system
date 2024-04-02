package loan

import (
	"linus/lms/finance"
	"reflect"
)

type LoanDetails struct {
	LoanAmount           float64
	ApprovedAmount       float64
	Tenure               float64
	InterestType         int
	PaymentFrequency     int
	InterestRate         float64
	monthlyInterestRate  float64
	totalInterest        float64
	subventionAmount     float64
	ChargesConfiguration ChargesConfiguration
}

type ChargesConfiguration struct {
	Charges []Charges
}

func (loanDetail LoanDetails) GetNetLoanAmount() float64 {
	// Initialize amount with the loan amount
	amount := loanDetail.LoanAmount

	// Get the charge's configuration from the loan details
	charges := loanDetail.ChargesConfiguration.Charges

	// Iterate over each charge and subtract its amount from the total amount
	for _, charge := range charges {

		amount -= charge.GetAmountToBeDeductedFromLoanAmount(loanDetail)
	}

	// Return the remaining amount after deducting charges
	return amount
}

func (loanDetail LoanDetails) Init() LoanDetails {
	loanDetail.monthlyInterestRate = finance.ReducingInterestRate(loanDetail.InterestRate, loanDetail.GetNetLoanAmount(), loanDetail.Tenure, loanDetail.InterestType, loanDetail.PaymentFrequency)
	return loanDetail
}

func (loanDetail LoanDetails) GetEMI() float64 {
	return finance.CalculatePMT(loanDetail.LoanAmount, loanDetail.monthlyInterestRate, loanDetail.Tenure)
}

func (loanDetail LoanDetails) SubventionFeeCharges() SubventionFeeCharges {
	return loanDetail.GetCharges(SubventionFeeCharges{}).(SubventionFeeCharges)
}

func (loanDetail LoanDetails) ProcessingFeeCharges() ProcessingFeeCharges {
	return loanDetail.GetCharges(ProcessingFeeCharges{}).(ProcessingFeeCharges)
}

func (loanDetail LoanDetails) DownPaymentCharges() DownPayment {
	return loanDetail.GetCharges(DownPayment{}).(DownPayment)
}

func (loanDetail LoanDetails) GetTotalInterest() float64 {
	return finance.GetTotalInterestCharged(loanDetail.InterestRate, loanDetail.GetNetLoanAmount())
}

func (loanDetail LoanDetails) GetReducingInterestRate() float64 {
	return loanDetail.monthlyInterestRate
}

func (loanDetail LoanDetails) GetCharges(_interface interface{}) interface{} {
	// Get the charge's configuration from the loan details
	_charges := loanDetail.ChargesConfiguration.Charges

	interfaceType := reflect.TypeOf(_interface)
	for _, charge := range _charges {
		shapeType := reflect.TypeOf(charge)
		if shapeType == interfaceType {
			return charge.Init(loanDetail)
		}
	}
	return _interface
}
