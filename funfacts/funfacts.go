package funfacts

import (
	"fmt"

	"github.com/andrinecCF/funtemps/conv"
)

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

// SKal konverte verdiene med funskjoner

type FunFacts struct {
	Terra struct {
		Celsius     []float64
		Kelvin      []float64
		Farhrenheit []float64
	}
	Sun struct {
		Celsius     []float64
		Kelvin      []float64
		Farhrenheit []float64
	}
	Luna struct {
		Celsius     []float64
		Kelvin      []float64
		Farhrenheit []float64
	}
}

func GetFunFacts(fact string, tmp string) []float64 {

	var funFacts FunFacts
	funFacts.Sun.Kelvin = []float64{5778}
	funFacts.Sun.Celsius = []float64{55000000}
	funFacts.Sun.Farhrenheit = []float64{conv.CelsiusToFahrenheit(55000000)}

	funFacts.Luna.Celsius = []float64{-183, 106}
	funFacts.Luna.Kelvin = []float64{conv.CelsiusToKevin(-183), conv.CelsiusToKevin(106)}
	funFacts.Luna.Farhrenheit = []float64{conv.CelsiusToFahrenheit(-183), conv.CelsiusToFahrenheit(106)}

	funFacts.Terra.Farhrenheit = []float64{134, conv.CelsiusToFahrenheit(-89.4), conv.KelvinToFarhenheit(9329)}
	funFacts.Terra.Celsius = []float64{56.7, -89.4, conv.KevinToCelsius(9329)}
	funFacts.Terra.Kelvin = []float64{329.82, conv.CelsiusToKevin(-89.4), 9329}

	switch {
	case fact == "sun":
		if tmp == "C" {
			return funFacts.Sun.Celsius
		} else if tmp == "K" {
			return funFacts.Sun.Kelvin
		} else if tmp == "F" {
			return funFacts.Sun.Farhrenheit
		}
		break
	case fact == "luna":
		if tmp == "C" {
			return funFacts.Luna.Celsius
		} else if tmp == "K" {
			return funFacts.Luna.Kelvin
		} else if tmp == "F" {
			return funFacts.Luna.Farhrenheit
		}
		break
	case fact == "terra":
		if tmp == "C" {
			return funFacts.Terra.Celsius
		} else if tmp == "K" {
			return funFacts.Terra.Kelvin
		} else if tmp == "F" {
			return funFacts.Terra.Farhrenheit
		}
		break

	default:
		fmt.Println("Ikke gyldig verdi")
	}

	return []float64{}
}
