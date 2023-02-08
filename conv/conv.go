package conv

import (
	"math"
)

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    FarhenheitToKelvin
    CelsiusToFahrenheit
    CelsiusToKevin
    KelvinToFarhenheit
    KevinToCelsius
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	celsius := (value - 32) * (5.0 / 9.0)

	return math.Round(celsius*100) / 100
}

// De andre konverteringsfunksjonene implementere her
// ...

func FarhenheitToKelvin(value float64) float64 {
	return 0
}

func CelsiusToFahrenheit(value float64) float64 {
	return 0
}

func CelsiusToKevin(value float64) float64 {
	return 0
}

func KelvinToFarhenheit(value float64) float64 {
	return 0
}

func KevinToCelsius(value float64) float64 {
	return 0
}
