package funfacts

import "fmt"

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
type FunFacts struct {
	Terra struct {
		Celsius     string
		Kelvin      string
		Farhrenheit string
	}
	Sun struct {
		Celsius string
		Kelvin  string
	}
	Luna struct {
		Celsius string
	}
}

func GetFunFacts(fact string, tmp string) string {

	// pesude code
	// switch
	// case for alle forskjell verdiene
	// i case return ønsket verdi

	var funFacts FunFacts
	funFacts.Sun.Kelvin = "Temperatur i Solens kjerne er 5778K."
	funFacts.Sun.Celsius = "Temperatur i Solens kjerne 15000000°C"

	funFacts.Luna.Celsius = "Temperatur på Månens overflate om natten -183°C.\nTemperatur på Månens overflate om dagen 106°C."

	funFacts.Terra.Farhrenheit = "Høyeste temperatur målt på Jordens overflate 134°F"
	funFacts.Terra.Celsius = "Høyeste temperatur målt på Jordens overflate 56.7°C\nLaveste temperatur målt på Jordens overflate -89.4°C"
	funFacts.Terra.Kelvin = "Temperatur i Jordens indre kjerne 9392K\nHøyeste temperatur målt på Jordens overflate 329.82K"

	switch {
	case fact == "Sun":
		if tmp == "C" {
			return funFacts.Sun.Celsius
		} else if tmp == "K" {
			return funFacts.Sun.Kelvin
		}
		break
	case fact == "Luna":
		if tmp == "C" {
			return funFacts.Luna.Celsius
		}
		break
	case fact == "Terra":
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

	return ""
}
