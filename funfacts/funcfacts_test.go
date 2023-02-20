package funfacts

import (
	"math"
	"testing"
)

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input []string  // her må du skrive riktig type for input
		want  []float64 // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{input: []string{"sun", "K"}, want: []float64{5778}},
		{input: []string{"sun", "C"}, want: []float64{55000000}},
		{input: []string{"sun", "F"}, want: []float64{99000032}},

		{input: []string{"terra", "F"}, want: []float64{134, -128.92, 16332.53}},
		{input: []string{"terra", "C"}, want: []float64{56.7, -89.4, 9055.85}},
		{input: []string{"terra", "K"}, want: []float64{329.82, 183.75, 9329}},

		{input: []string{"luna", "C"}, want: []float64{-183, 106}},
		{input: []string{"luna", "F"}, want: []float64{-297.4, 222.8}},
		{input: []string{"luna", "K"}, want: []float64{90.15, 379.15}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input[0], tc.input[1])
		for i := range tc.want {
			if !withinTolerance(tc.want[i], got[i], 1e-2) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		}

	}

}
