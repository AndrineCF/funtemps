package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input []string // her må du skrive riktig type for input
		want  string   // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{input: []string{"Sun", "K"}, want: "Temperatur i Solens kjerne er 5778K."},
		{input: []string{"Sun", "C"}, want: "Temperatur i Solens kjerne 15000000°C"},
		{input: []string{"Terra", "C"}, want: "Høyeste temperatur målt på Jordens overflate 56.7°C\nLaveste temperatur målt på Jordens overflate -89.4°C"},
		{input: []string{"Terra", "F"}, want: "Høyeste temperatur målt på Jordens overflate 134°F"},
		{input: []string{"Terra", "K"}, want: "Temperatur i Jordens indre kjerne 9392K\nHøyeste temperatur målt på Jordens overflate 329.82K"},
		{input: []string{"Luna", "C"}, want: "Temperatur på Månens overflate om natten -183°C.\nTemperatur på Månens overflate om dagen 106°C."},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input[0], tc.input[1])
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
