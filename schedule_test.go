package schedule

import (
	Charges "linus/lms/charges"
	"linus/lms/constants"
	"linus/lms/loan"
	"testing"
)

func TestGetRepayment(t *testing.T) {

	// Initialize LoanDetails with specific values
	LoanDetails := loan.LoanDetails{
		Tenure:           12.00,                    // Duration of the loan in years
		LoanAmount:       50000,                    // Initial loan amount
		InterestType:     constants.Reducing,       // Type of interest calculation
		InterestRate:     13.00,                    // Annual interest rate
		PaymentFrequency: constants.PaymentMonthly, // Frequency of payments
		ChargesConfiguration: loan.ChargesConfiguration{ // Charges configuration
			Charges: []Charges.Charges{ // Slice of charge configurations
				// Processing fee charge configuration
				Charges.ProcessingFeeCharges{
					Frequency:        0,     // Frequency of processing fee charge
					Amount:           999,   // Amount of processing fee
					LSPCollected:     false, // Whether processing fee collected by LSP
					PaidWithFirstEMI: false, // Whether processing fee paid with first EMI
				},
				// Subvention fee charge configuration
				Charges.SubventionFeeCharges{
					Amount:            4000, // Amount of subvention fee
					LSPCollected:      true, // Whether subvention fee collected by LSP
					ChargeOnNetAmount: true, // Whether subvention fee charged on net amount
				},
				// Down payment charge configuration
				Charges.DownPayment{
					Amount:               0,     // Amount of down payment
					LSPCollected:         false, // Whether down payment collected by LSP
					DeductFromLoanAmount: true,  // Whether down payment deducted from loan amount
				},
			},
		},
	}
}
