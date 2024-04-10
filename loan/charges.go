package loan

import "reflect"

func (loanDetail LoanDetails) SubventionFeeCharges() SubventionFeeCharges {
	return loanDetail.GetCharges(SubventionFeeCharges{}).(SubventionFeeCharges)
}

func (loanDetail LoanDetails) ProcessingFeeCharges() ProcessingFeeCharges {
	return loanDetail.GetCharges(ProcessingFeeCharges{}).(ProcessingFeeCharges)
}

func (loanDetail LoanDetails) DownPaymentCharges() DownPayment {
	return loanDetail.GetCharges(DownPayment{}).(DownPayment)
}

func (loanDetail LoanDetails) HoldAmountCharge() AmountHold {
	return loanDetail.GetCharges(AmountHold{}).(AmountHold)
}

func (loanDetail LoanDetails) GetCharges(_interface interface{}) interface{} {
	// Get the charge's configuration from the loan details
	charges := loanDetail.ChargesConfiguration.Charges

	for _, charge := range charges {
		// Check if the charge is of the same type as the provided interface
		if reflect.TypeOf(charge) == reflect.TypeOf(_interface) {
			return charge.Init(loanDetail)
		}
	}

	// Return the provided interface if no matching charge type is found
	return _interface
}
