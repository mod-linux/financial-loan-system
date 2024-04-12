package finance

import (
	"encoding/json"
	"fmt"
	"linus/lms/constants"
	"math"
)

func ReducingInterestRate(rate, netAmount, tenure float64, interestType, frequency int) float64 {
	if interestType == constants.Flat {
		// Only calculate the reducing interest rate if the interest type is reducing
		return toReducing(GetTotalInterestCharged(rate, netAmount), tenure, frequency, netAmount)
	} else {
		// Calculate the periodic interest rate by dividing the annual interest rate (r) by the frequency
		return toFrequency(rate, frequency)
	}
}

func toFrequency(rate float64, frequency int) float64 {
	return rate / float64(frequency) / 100
}

func toReducing(interestPaid, remainingTenure float64, frequency int, netAmount float64) float64 {
	// Initial guess for the rate
	upfrontFromCustomer := 0.0
	rate := calculate(remainingTenure, (netAmount+interestPaid)/remainingTenure, -netAmount+upfrontFromCustomer, 0, 0.1, false) * 12 * 100
	return toFrequency(rate, frequency)
}

func calculate(periods, payment, presentValue, futureValue, initialGuess float64, beginning bool) float64 {
	when := 0.0
	if beginning {
		when = 1.0
	}

	f := func(x []float64) float64 {
		return x[4] + x[3]*math.Pow(1+x[0], x[1]) + x[2]*(1+x[0]*x[5])/x[0]*((math.Pow(1+x[0], x[1]))-1)
	}

	arg := []float64{initialGuess, periods, payment, presentValue, futureValue, when}

	return solve(f, arg, 0, 1e-6, 0, 100)
}

// Function to solve
func solve(function func([]float64) float64, args []float64, target float64, tol float64, position int, iterations int) float64 {
	// Initialize
	args1 := make([]float64, len(args))
	copy(args1, args)
	guess := args[position]
	i := 0

	for {
		args1[position] = guess + tol // Load the initial guess into the arguments
		args[position] = guess        // Load the initial guess into the arguments
		y := function(args)
		yAtXPlusDelX := function(args1)
		slope := (yAtXPlusDelX - y) / tol
		delY := target - y
		if math.Abs(slope) < tol {
			return math.NaN()
		}
		guess = delY/slope + guess
		dif := math.Abs(delY)
		i++
		if dif <= tol || i >= iterations {
			break
		}
	}

	if i >= iterations {
		return math.NaN()
	}
	return guess
}

func CalculatePMT(amount, r, tenure float64) float64 {
	// Calculate the numerator of the PMT formula: amount multiplied by the interest rate
	numerator := amount * r
	// Calculate the denominator of the PMT formula: 1 - (1 + rate) raised to the power of -tenure
	denominator := 1 - math.Pow(1+r, -tenure)
	// Calculate the PMT (periodic payment) using the formula: numerator divided by denominator
	return numerator / denominator
}

func GetTotalInterestCharged(r, netAmount float64) float64 {
	return r * netAmount / 100
}

func CalculateGST(amount, cgst, sgst, igst float64) float64 {
	// Initialize GST amount
	gstAmount := 0.0

	// Calculate GST for each component
	gstAmount += amount * (cgst + sgst + igst) / 100

	return gstAmount
}

type Schedule struct {
	EMI             float64
	InterestAmount  float64
	PrincipalAmount float64
	Balance         float64
}

func GetLoanSchedule(tenure, emi, amount, monthlyInterest float64) []Schedule {

	var schedules []Schedule

	balance := amount
	for i := 0; i < int(tenure)-1; i++ {
		var schedule Schedule

		schedule.EMI = math.Round(emi)
		schedule.InterestAmount = balance * monthlyInterest
		schedule.PrincipalAmount = schedule.EMI - schedule.InterestAmount
		balance -= schedule.PrincipalAmount
		schedule.Balance = balance

		schedules = append(schedules, schedule)
	}

	lastSchedule := Schedule{}
	lastSchedule.InterestAmount = balance * monthlyInterest
	lastSchedule.PrincipalAmount = balance
	lastSchedule.Balance = 0
	lastSchedule.EMI = lastSchedule.InterestAmount + lastSchedule.PrincipalAmount

	schedules = append(schedules, lastSchedule)

	b, err := json.Marshal(schedules)
	if err != nil {
		fmt.Println(err)
		return schedules
	}
	fmt.Println(string(b))
	return schedules
}
