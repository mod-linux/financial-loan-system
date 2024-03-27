package finance

import "math"

func ConvertInterestRateBasedOnFrequency(r float64, frequency int) float64 {
	// Calculate the periodic interest rate by dividing the annual interest rate (r) by the frequency
	return r / float64(frequency) / 100
}

func ToReducing(r, tenure float64, frequency int, presentValue float64) float64 {
	// Calculate the interest rate per period
	rate := r / float64(frequency)

	if rate > 0 {
		// Calculate the numerator: present value multiplied by (1 + rate) raised to the power of tenure
		numerator := -presentValue * math.Pow(1+rate, tenure)
		// Calculate the denominator: 1 - (1 + rate)
		denominator := 1 - (1 + rate)
		// Calculate the reducing interest rate using the annuity formula
		return -numerator / denominator / ((math.Pow(1+rate, tenure) - 1) / rate)
	} else {
		// If the interest rate is 0 or negative, return the present value divided by the tenure
		return -presentValue / tenure
	}
}

func CalculatePMT(amount, r, tenure float64) float64 {
	// Calculate the numerator of the PMT formula: amount multiplied by the interest rate
	numerator := amount * r
	// Calculate the denominator of the PMT formula: 1 - (1 + rate) raised to the power of -tenure
	denominator := 1 - math.Pow(1+r, -tenure)
	// Calculate the PMT (periodic payment) using the formula: numerator divided by denominator
	return numerator / denominator
}
