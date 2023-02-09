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
		Celsius     []string
		Kelvin      []string
		Farhrenheit []string
	}
	Sun struct {
		Celsius []string
		Kelvin  []string
	}
	Luna struct {
		Celsius []string
	}
}

func GetFunFacts(fact string, tmp string) string {

	// pesude code
	// switch
	// case for alle forskjell verdiene
	// i case return ønsket verdi

	var funFacts FunFacts
	funFacts.Sun.Kelvin = []string{"Temperatur på ytre lag av Solen 5506.85K.", "Temperatur i Solens kjerne er 15 000 000°C."}
	funFacts.Sun.Celsius = []string{"Temperatur i Solens kjerne"}

	funFacts.Luna.Celsius = []string{"Temperatur på Månens overflate om natten -183°C.", "Temperatur på Månens overflate om dagen 106°C."}

	funFacts.Terra.Farhrenheit = []string{"Høyeste temperatur målt på Jordens overflate 134°F"}
	funFacts.Terra.Celsius = []string{, "Høyeste temperatur målt på Jordens overflate 56.7°C", "Laveste temperatur målt på Jordens overflate -89.4°C"}
	funFacts.Terra.Kelvin = []string{"Temperatur i Jordens indre kjerne 9392K", "Høyeste temperatur målt på Jordens overflate 329.82K"}

	return funFacts.Sun.Kelvin[0] + "\n" + funFacts.Sun.Kelvin[1]
}
