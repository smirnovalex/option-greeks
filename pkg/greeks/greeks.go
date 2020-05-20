package greeks

import (
	"math"
)

const (
	OptionCall = 1
	OptionPut  = 2
)

// Calculates the delta of an option. interestRate Annual risk-free interest rate, optionType - The type of option - OptionCall or OptionPut
func GetDelta(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64, optionType int) float64 {
	if optionType == OptionCall {
		return callDelta(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
	}

	return putDelta(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
}

// Calculates the delta of a call option.
func callDelta(currentPrice float64, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)

	if !IsFinite(omega) {
		if currentPrice > strikePrice {
			return 1
		} else {
			return 0
		}
	}
	return GetStandardNormalCumulativeDistribution(omega)
}

// Calculate the delta of a pull option
func putDelta(currentPrice float64, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var delta = callDelta(currentPrice, strikePrice, timeToExpiration, volatility, interestRate) - 1

	if delta == -1 && strikePrice == currentPrice {
		return 0
	}
	return delta
}

// Calculate omega as defined in the Black-Scholes formula
func getOmega(currentPrice float64, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	return (interestRate*timeToExpiration + math.Pow(volatility, float64(2))*timeToExpiration/float64(2) - math.Log(strikePrice/currentPrice)) / (volatility * math.Sqrt(timeToExpiration))
}

// Calculates the theta of an option. scale - The number of days to scale theta by - usually 365 or 252
func GetTheta(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64, optionType int, scale float64) float64 {
	if optionType == OptionCall {
		return callTheta(currentPrice, strikePrice, timeToExpiration, volatility, interestRate) / scale
	}

	return putTheta(currentPrice, strikePrice, timeToExpiration, volatility, interestRate) / scale
}

// Calculates the theta of a call option
func callTheta(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
	if IsFinite(omega) {
		return -1*volatility*currentPrice*GetStandardNormalDensity(omega)/(2*math.Sqrt(timeToExpiration)) - strikePrice*interestRate*math.Pow(math.E, -1*interestRate*timeToExpiration)*GetStandardNormalCumulativeDistribution(omega-volatility*math.Sqrt(timeToExpiration))
	} else {
		return 0
	}
}

// Calculates the theta of a put option
func putTheta(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
	if IsFinite(omega) {
		return -1*volatility*currentPrice*GetStandardNormalDensity(omega)/(2*math.Sqrt(timeToExpiration)) + strikePrice*interestRate*math.Pow(math.E, -1*interestRate*timeToExpiration)*GetStandardNormalCumulativeDistribution(volatility*math.Sqrt(timeToExpiration)-omega)
	}
	return 0
}

// Calculates the gamma of a call and put option
func GetGamma(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)

	if IsFinite(omega) {
		return GetStandardNormalDensity(omega) / (currentPrice * volatility * math.Sqrt(timeToExpiration))
	}

	return 0
}

// Calculates the rho of an option. Scale - The value to scale rho by (100=100BPS=1%, 10000=1BPS=.01%)
func GetRho(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64, optionType int, scale float64) float64 {
	if optionType == OptionCall {
		return callRho(currentPrice, strikePrice, timeToExpiration, volatility, interestRate) / scale
	}
	return putRho(currentPrice, strikePrice, timeToExpiration, volatility, interestRate) / scale
}

// Calculates the rho of a call option
func callRho(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
	if !IsNaN(omega) {
		return strikePrice * timeToExpiration * math.Pow(math.E, -1*interestRate*timeToExpiration) * GetStandardNormalCumulativeDistribution(omega-volatility*math.Sqrt(timeToExpiration))
	}
	return 0
}

// Calculates the rho of a put option
func putRho(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
	if !IsNaN(omega) {
		return -1 * strikePrice * timeToExpiration * math.Pow(math.E, -1*interestRate*timeToExpiration) * GetStandardNormalCumulativeDistribution(volatility*math.Sqrt(timeToExpiration)-omega)
	}
	return 0
}

// Calculates the vega of a call and put option
func GetVega(currentPrice, strikePrice float64, timeToExpiration float64, volatility float64, interestRate float64) float64 {
	var omega = getOmega(currentPrice, strikePrice, timeToExpiration, volatility, interestRate)
	if IsFinite(omega) {
		return currentPrice * math.Sqrt(timeToExpiration) * GetStandardNormalDensity(omega) / 100
	}
	return 0
}
