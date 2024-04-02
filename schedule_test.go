package schedule

import (
	"fmt"
	"linus/lms/constants"
	"linus/lms/loan"
	"testing"
)

func TestGetRepayment(t *testing.T) {

	// Initialize LoanDetails with specific values
	LoanDetails := loan.LoanDetails{
		Tenure:           12.00,                    // Duration of the loan in month
		LoanAmount:       50000,                    // Initial loan amount
		ApprovedAmount:   50000,                    // Apporved loan amount
		InterestType:     constants.Reducing,       // Type of interest calculation
		InterestRate:     13,                       // Annual interest rate
		PaymentFrequency: constants.PaymentMonthly, // Frequency of payments
		ChargesConfiguration: loan.ChargesConfiguration{ // Charges configuration
			Charges: []loan.Charges{ // Slice of charge configurations
				// Processing fee charge configuration
				loan.ProcessingFeeCharges{
					Frequency:        0,     // Frequency of processing fee charge
					Amount:           999,   // Amount of processing fee
					LSPCollected:     false, // Whether processing fee collected by LSP
					PaidWithFirstEMI: false, // Whether processing fee paid with first EMI
				},
				// Subvention fee charge configuration
				loan.SubventionFeeCharges{
					Amount:                 4000,  // Amount of subvention fee
					LSPCollected:           false, // Whether subvention fee collected by LSP
					ChargeOnApprovedAmount: true,  // Whether subvention fee charged on net amount
				},
				// Down payment charge configuration
				loan.DownPayment{
					Amount:               0,     // Amount of down payment
					LSPCollected:         false, // Whether down payment collected by LSP
					DeductFromLoanAmount: true,  // Whether down payment deducted from loan amount
				},
			},
		},
	}

	loanDetails := LoanDetails.Init()

	emi := loanDetails.GetEMI()
	subvention := loanDetails.SubventionFeeCharges().GetSubventionAmount()
	processing := loanDetails.ProcessingFeeCharges().GetProcessingFeeAmount()
	downPayment := loanDetails.DownPaymentCharges().GetDownPayment()
	netAmount := loanDetails.GetNetLoanAmount()

	fmt.Println(emi)
	fmt.Println(subvention)
	fmt.Println(processing)
	fmt.Println(downPayment)
	fmt.Println(netAmount)
}
