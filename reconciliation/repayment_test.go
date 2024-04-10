package reconciliation

import (
	"linus/lms/constants"
	"linus/lms/loan"
	"testing"
	"time"
)

func TestReconcile(t *testing.T) {
	LoanDetails := loan.LoanDetails{
		Tenure:           12.00,                    // Duration of the loan in month
		LoanAmount:       50000,                    // Initial loan amount
		ApprovedAmount:   50000,                    // Approved loan amount
		InterestType:     constants.Reducing,       // Type of interest calculation
		InterestRate:     13,                       // Annual interest rate
		PaymentFrequency: constants.PaymentMonthly, // Frequency of payments
		ChargesConfiguration: loan.ChargesConfiguration{ // Charges configuration
			Charges: []loan.Charges{ // Slice of charge configurations
				// Processing fee charge configuration
				loan.ProcessingFeeCharges{
					Frequency:              0,     // Frequency of processing fee charge
					Amount:                 0,     // Amount of processing fee
					LSPCollected:           false, // Whether processing fee collected by LSP
					DeductFromDisbursement: false, // Whether processing fee collected by LSP
					PaidWithFirstEMI:       false, // Whether processing fee paid with first EMI
				},
				// Subvention fee charge configuration
				loan.SubventionFeeCharges{
					//Amount:                 4000,  // Amount of subvention fee
					Percentage:             12.90,
					ChargeInPercentage:     true,
					LSPCollected:           true, // Whether subvention fee collected by LSP
					DeductFromDisbursement: true,
					DeductFromHoldAmount:   true,
					ChargeOnApprovedAmount: true, // Whether subvention fee charged on net amount
					GSTApplicable:          true,
					TAXES: loan.GSTCharges{
						//CGST: 9,
						//SGST: 9,
						IGST: 18,
					},
				},
				// Down payment charge configuration
				loan.DownPayment{
					Amount:               0,     // Amount of down payment
					LSPCollected:         false, // Whether down payment collected by LSP
					DeductFromLoanAmount: true,  // Whether down payment deducted from loan amount
				},

				loan.AmountHold{
					ChargeInPercentage:     true, // Whether hold amount collected by LSP
					DeductFromDisbursement: true, // Whether hold amount deducted From Disbursement
					Percentage:             5,    // Percentage charged
				},
			},
		},
	}

	payment := Payment{
		Amount:      4500,
		Foreclosure: false,
		PaymentDate: time.Now(),
		LoanRepaymentObject: LoanRepaymentObject{
			Schedule: LoanDetails.Init().GetLoanSchedule(),
		},
	}

	payment.Reconcile()
}
