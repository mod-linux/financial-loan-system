package finance

import (
	"fmt"
)

func ReducingInterestRateTest(rate, netAmount, tenure float64, interestType, frequency int) {
	reducingRate := ReducingInterestRate(rate, netAmount, tenure, interestType, frequency)
	fmt.Printf("Calculate Reducing rate based on type of interest %.2f%%\n", reducingRate)
}

func calculateTest(periods, payment, presentValue, futureValue, initialGuess float64, beginning bool) {
	amount := calculate(periods, payment, presentValue, futureValue, initialGuess, beginning)
	fmt.Printf("Calculate Reducing rate based on type of interest %.2f%%\n", amount)
}

func CalculatePMTTest(amount, r, tenure float64) {
	pmt := CalculatePMT(amount, r, tenure)
	fmt.Printf("Calculate Reducing rate based on type of interest %.2f", pmt)
}

func GetLoanScheduleTest(tenure, emi, amount, monthlyInterest float64) {
	schedule := GetLoanSchedule(tenure, emi, amount, monthlyInterest)
	fmt.Printf("Loan Scheudle %+v", schedule)
}
