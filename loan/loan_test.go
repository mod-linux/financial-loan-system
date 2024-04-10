package loan

import (
	"linus/lms/constants"
	"testing"
)

func TestGetLoanInstance(t *testing.T) {

	// Initialize LoanDetails with specific values
	LoanDetails := LoanDetails{
		Tenure:           12.00,                    // Duration of the loan in month
		LoanAmount:       50000,                    // Initial loan amount
		ApprovedAmount:   50000,                    // Approved loan amount
		InterestType:     constants.Reducing,       // Type of interest calculation
		InterestRate:     13,                       // Annual interest rate
		PaymentFrequency: constants.PaymentMonthly, // Frequency of payments
		ChargesConfiguration: ChargesConfiguration{ // Charges configuration
			Charges: []Charges{ // Slice of charge configurations
				// Processing fee charge configuration
				ProcessingFeeCharges{
					Frequency:              0,     // Frequency of processing fee charge
					Amount:                 0,     // Amount of processing fee
					LSPCollected:           false, // Whether processing fee collected by LSP
					DeductFromDisbursement: false, // Whether processing fee collected by LSP
					PaidWithFirstEMI:       false, // Whether processing fee paid with first EMI
				},
				// Subvention fee charge configuration
				SubventionFeeCharges{
					//Amount:                 4000,  // Amount of subvention fee
					Percentage:             12.90,
					ChargeInPercentage:     true,
					LSPCollected:           true, // Whether subvention fee collected by LSP
					DeductFromDisbursement: true,
					DeductFromHoldAmount:   true,
					ChargeOnApprovedAmount: true, // Whether subvention fee charged on net amount
					GSTApplicable:          true,
					TAXES: GSTCharges{
						//CGST: 9,
						//SGST: 9,
						IGST: 18,
					},
				},
				// Down payment charge configuration
				DownPayment{
					Amount:               0,     // Amount of down payment
					LSPCollected:         false, // Whether down payment collected by LSP
					DeductFromLoanAmount: true,  // Whether down payment deducted from loan amount
				},

				AmountHold{
					ChargeInPercentage:     true, // Whether hold amount collected by LSP
					DeductFromDisbursement: true, // Whether hold amount deducted From Disbursement
					Percentage:             5,    // Percentage charged
				},
			},
		},
	}

	loanDetails := LoanDetails.Init()

	//emi := loanDetails.GetEMI()
	//subventionCharges := loanDetails.SubventionFeeCharges()
	//subvention := subventionCharges.GetSubventionAmount()
	//subventionGST := subventionCharges.GetSubventionGST()
	//subventionWithGST := subventionCharges.GetSubventionAmountWithGST()
	//processing := loanDetails.ProcessingFeeCharges().GetProcessingFeeAmount()
	//downPayment := loanDetails.DownPaymentCharges().GetDownPayment()
	//netAmount := loanDetails.GetNetLoanAmount()
	//holdAmountCharge := loanDetails.HoldAmountCharge()
	//holdAmount := holdAmountCharge.GetHoldAmount()
	loanDetails.GetLoanSchedule()
	//
	//fmt.Println("EMI:", emi)
	//fmt.Println("Subvention:", subvention)
	//fmt.Println("SubventionGST:", subventionGST)
	//fmt.Println("SubventionWithGST:", subventionWithGST)
	//fmt.Println("Processing:", processing)
	//fmt.Println("DownPayment:", downPayment)
	//fmt.Println("NetAmount:", netAmount)
	//fmt.Println("HoldAmount:", holdAmount)
}
