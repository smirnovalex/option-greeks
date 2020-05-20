package greeks

import "math"

// Standard normal cumulative distribution function.  The probability is estimated
// by expanding the CDF into a series using the first 100 terms. Returns the probability that a
// standard normal random variable will be less than or equal to number (The upper bound to integrate over.
// This is P{Z <= x} where Z is a standard normal random variable.)
// See http://en.wikipedia.org/wiki/Normal_distribution#Cumulative_distribution_function|Wikipedia
func GetStandardNormalCumulativeDistribution(number float64) float64 {
	var result = float64(0)
	// avoid divergence in the series which happens around +/-8 when summing the
	// first 100 terms
	if number >= 8 {
		result = 1
	} else if number <= -8 {
		result = 0
	} else {
		for i := 0; i < 100; i++ {
			result = result + math.Pow(number, float64(2*i+1))/DoubleFactorial(float64(2*i+1))
		}
		result *= math.Pow(math.E, -0.5*math.Pow(number, 2))
		result /= math.Sqrt(2 * math.Pi)
		result += 0.5
	}

	return result
}

// Double factorial http://en.wikipedia.org/wiki/Double_factorial|Wikipedia page
func DoubleFactorial(number float64) float64 {
	var result = float64(1)
	for i := number; i > 1; i -= 2 {
		result = result * i
	}

	return result
}

// Standard normal density function.
func GetStandardNormalDensity(number float64) float64 {
	return math.Pow(math.E, -1*math.Pow(number, 2)/2) / math.Sqrt(2*math.Pi)
}

// IsNaN reports whether number is an IEEE 754 "not-a-number" value.
func IsNaN(number float64) bool {
	// IEEE 754 says that only NaNs satisfy number != number.
	return number != number
}

// IsFinite reports whether number is neither NaN nor an infinity.
func IsFinite(number float64) bool {
	return !IsNaN(number - number)
}
