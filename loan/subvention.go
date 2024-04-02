package loan

type SubventionFeeCharges struct {
	Amount                 float64
	ChargeOnAppliedAmount  bool
	ChargeOnApprovedAmount bool
	ChargeInPercentage     bool
	Percentage             float64
	LSPCollected           bool
	GSTApplicable          bool
	TAXES                  GSTCharges
}

func (subventionFeeCharges SubventionFeeCharges) Init(loan LoanDetails) interface{} {
	if subventionFeeCharges.Amount > 0 {
		return subventionFeeCharges
	}
	if subventionFeeCharges.ChargeInPercentage {
		if subventionFeeCharges.ChargeOnAppliedAmount {
			subventionFeeCharges.Amount = loan.LoanAmount * subventionFeeCharges.Percentage
		} else if subventionFeeCharges.ChargeOnApprovedAmount {
			subventionFeeCharges.Amount = loan.ApprovedAmount * subventionFeeCharges.Percentage
		}
	}
	return subventionFeeCharges
}

func (subventionFeeCharges SubventionFeeCharges) GetAmountToBeDeductedFromLoanAmount(loan LoanDetails) float64 {
	if !subventionFeeCharges.LSPCollected {
		return subventionFeeCharges.Init(loan).(SubventionFeeCharges).Amount
	}
	return 0
}

func (subventionFeeCharges SubventionFeeCharges) GetSubventionAmount() float64 {
	return subventionFeeCharges.Amount
}
