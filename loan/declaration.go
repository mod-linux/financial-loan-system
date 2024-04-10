package loan

type Charges interface {
	Init(LoanDetails) interface{}
	DeductedFromDisbursement() float64
	DeductedFromHoldAmount() float64
}

type GSTCharges struct {
	IGST float64
	CGST float64
	SGST float64
}
