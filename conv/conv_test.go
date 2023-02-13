package conv

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

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	// Må legge til mer?
	tests := []test{
		{input: 134, want: 56.67},

		// Lagt til en ekstra test verdi
		{input: 150, want: 65.56},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	// Må legge til mer?
	tests := []test{
		{input: 100, want: 310.93},
		{input: 200, want: 366.48},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	// Må legge til mer?
	tests := []test{
		{input: 50, want: 122},
		{input: 66, want: 150.8},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKevin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	// Må legge til mer?
	tests := []test{
		{input: 34, want: 307.15},
		{input: 10, want: 283.15},
	}

	for _, tc := range tests {
		got := CelsiusToKevin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	// Må legge til mer?
	tests := []test{
		{input: 270, want: 26.33},
		{input: 290, want: 62.33},
	}

	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKevinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	// Må legge til mer?
	tests := []test{
		{input: 322, want: 48.85},
		{input: 300, want: 26.85},
	}

	for _, tc := range tests {
		got := KevinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
