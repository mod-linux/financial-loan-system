package charges

type Charges interface {
	Calculate() float64
	GetAmountToBeDeductedFromLoanAmount() float64
}

type LoanDetails interface {
	GetEMI() float64
}

type GSTCharges struct {
	IGST int
	CGST int
	SGST int
}
