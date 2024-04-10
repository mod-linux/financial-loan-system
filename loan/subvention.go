package loan

import (
	"linus/lms/finance"
)

type SubventionFeeCharges struct {
	Amount                 float64
	GSTAmount              float64
	ChargeOnAppliedAmount  bool
	ChargeOnApprovedAmount bool
	ChargeInPercentage     bool
	Percentage             float64
	LSPCollected           bool
	DeductFromDisbursement bool
	DeductFromHoldAmount   bool
	GSTApplicable          bool
	TAXES                  GSTCharges
}

func (subventionFeeCharges SubventionFeeCharges) Init(loan LoanDetails) interface{} {
	// If the amount is already set, return immediately
	if subventionFeeCharges.Amount > 0 {
		return subventionFeeCharges
	}

	// Calculate the subvention fee based on the appropriate amount
	var chargeAmount float64
	switch {
	case subventionFeeCharges.ChargeInPercentage && subventionFeeCharges.ChargeOnAppliedAmount:
		chargeAmount = loan.LoanAmount * subventionFeeCharges.Percentage / 100
	case subventionFeeCharges.ChargeInPercentage && subventionFeeCharges.ChargeOnApprovedAmount:
		chargeAmount = loan.ApprovedAmount * subventionFeeCharges.Percentage / 100
	default:
		// Return original charges if no calculations are needed
		return subventionFeeCharges
	}

	// Calculate GST charges
	subventionFeeCharges.Amount = chargeAmount
	subventionFeeCharges.GSTAmount = calculateGSTCharges(chargeAmount, subventionFeeCharges.TAXES)

	return subventionFeeCharges
}

func calculateGSTCharges(amount float64, charges GSTCharges) float64 {
	return finance.CalculateGST(amount, charges.CGST, charges.SGST, charges.IGST)
}

func (subventionFeeCharges SubventionFeeCharges) DeductedFromDisbursement() float64 {
	amount := 0.0
	if subventionFeeCharges.DeductFromDisbursement {
		if subventionFeeCharges.GSTApplicable {
			amount = subventionFeeCharges.GSTAmount
		}
		amount += subventionFeeCharges.Amount
	}
	return amount
}

func (subventionFeeCharges SubventionFeeCharges) DeductedFromHoldAmount() float64 {
	amount := 0.0
	if subventionFeeCharges.DeductFromHoldAmount {
		if subventionFeeCharges.GSTApplicable {
			amount = subventionFeeCharges.GSTAmount
		}
		amount += subventionFeeCharges.Amount
	}
	return amount
}

func (subventionFeeCharges SubventionFeeCharges) GetSubventionAmount() float64 {
	return subventionFeeCharges.Amount
}

func (subventionFeeCharges SubventionFeeCharges) GetSubventionAmountWithGST() float64 {
	return subventionFeeCharges.Amount + subventionFeeCharges.GSTAmount
}

func (subventionFeeCharges SubventionFeeCharges) GetSubventionGST() float64 {
	return subventionFeeCharges.GSTAmount
}
