package greeks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCalculateDelta(t *testing.T) {
	assert.Equal(t, 0.5076040742445566, GetDelta(100, 100, .086, .1, .0015, OptionCall))
	assert.Equal(t, -0.49239592575544344, GetDelta(100, 100, .086, .1, .0015, OptionPut))
}

func TestShouldCalculateTheta(t *testing.T) {
	// Theta - the default scale is 365 (days per year)
	assert.Equal(t, -0.03877971361524501, GetTheta(206.35, 206, .086, .1, .0015, OptionCall, 365))
	assert.Equal(t, -0.0379332474739548, GetTheta(206.35, 206, .086, .1, .0015, OptionPut, 365))

	// or you can set the scale to a value like 252 (trading days per year)
	assert.Equal(t, -0.05616902964112869, GetTheta(206.35, 206, .086, .1, .0015, OptionCall, 252))
	assert.Equal(t, -0.054942997333307556, GetTheta(206.35, 206, .086, .1, .0015, OptionPut, 252))
}

func TestShouldCalculateGamma(t *testing.T) {
	// Gamma - call and put gammas are equal at a given strike
	assert.Equal(t, 0.06573105549942765, GetGamma(206.35, 206, .086, .1, .0015))
}

func TestShouldCalculateVega(t *testing.T) {
	// Vega - call and put vegas are equal at a given strike
	// Note: vega is calculated per 1 percentage point change in volatility
	assert.Equal(t, 0.24070106056306834, GetVega(206.35, 206, .086, .1, .0015))
}

func TestShouldCalculateRho(t *testing.T) {
	// Rho - the default scale is 100 (rho per 1%, or 100BP, change in the risk-free interest rate)
	assert.Equal(t, 0.09193271711465777, GetRho(206.35, 206, .086, .1, .0015, OptionCall, 100))
	assert.Equal(t, -0.08520443071933861, GetRho(206.35, 206, .086, .1, .0015, OptionPut, 100))

	// or you can set the scale to a value like 10000 (rho per .01%, or 1BP, change in the risk-free interest rate)
	assert.Equal(t, 0.0009193271711465777, GetRho(206.35, 206, .086, .1, .0015, OptionCall, 10000))
	assert.Equal(t, -0.0008520443071933862, GetRho(206.35, 206, .086, .1, .0015, OptionPut, 10000))
}
