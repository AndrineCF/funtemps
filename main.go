package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/AndrineCF/funtemps/conv"
	"github.com/AndrineCF/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfact string
var funfactsTemp string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader Kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	flag.StringVar(&funfactsTemp, "t", "", "Hvilken grader funfacts skal vises i")
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfact)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	// konverte verdier fra farhenheit
	if out == "C" && isFlagPassed("F") {
		formatTemp(fahr, conv.FarhenheitToCelsius(fahr), "°F", "°C")
	} else if out == "K" && isFlagPassed("F") {
		formatTemp(fahr, conv.FarhenheitToKelvin(fahr), "°F", "K")
	}

	// Konverte verdier fra celsius
	if out == "F" && isFlagPassed("C") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		formatTemp(cels, conv.CelsiusToFahrenheit(cels), "°C", "°F")
	} else if out == "K" && isFlagPassed("C") {
		formatTemp(cels, conv.CelsiusToKevin(cels), "°C", "K")
	}

	//konverte verdier fra kelvin
	if out == "F" && isFlagPassed("K") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		formatTemp(kelv, conv.KelvinToFarhenheit(kelv), "K", "°F")
	} else if out == "C" && isFlagPassed("K") {
		formatTemp(kelv, conv.KevinToCelsius(kelv), "K", "°C")
	}

	if funfact == "luna" {
		fmt.Println(funfacts.GetFunFacts(funfact, funfactsTemp))
	}

	if funfact == "terra" {
		fmt.Println(funfacts.GetFunFacts(funfact, funfactsTemp))
	}

	if funfact == "sun" {
		fmt.Println(funfacts.GetFunFacts(funfact, funfactsTemp))
	}

}

func formatTemp(temp float64, konvTemp float64, inputF string, outputF string) {
	if temp == math.Trunc(temp) {
		fmt.Printf("%d%s er %.2f%s", int(temp), inputF, konvTemp, outputF)
	} else {
		fmt.Printf("%.3f%s er %.2f%s", temp, inputF, konvTemp, outputF)
	}
}

func formattingBignumber(temp float64) string {
	//Tar bort desimalene fra float64 om det er heltall
	s := ""
	if temp == math.Trunc(temp) {
		s = fmt.Sprintf("%.2f", temp)
	}

	//Konverte string til rune, altså symboler for string for å enkelt sett den i array
	sArray := []rune(s)
	tempS := ""
	result := ""

	// starter bakfram array og dele opp
	for i := len(sArray) - 1; i >= 0; i-- {
		tempS += string(sArray[i])

		//sett mellom i string, samt sjekker om den starter på desimaler
		if (i+1)%3 == 0 && string(sArray[i]) != "." {
			tempS += " "
		}
	}

	//reverser tilbake array som blir skrevet ut riktig
	for _, el := range tempS {
		result = string(el) + result
	}

	return result
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
