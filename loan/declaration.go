package loan

type Charges interface {
	Init(LoanDetails) interface{}
	GetAmountToBeDeductedFromLoanAmount(LoanDetails) float64
}
type GSTCharges struct {
	IGST int
	CGST int
	SGST int
}
