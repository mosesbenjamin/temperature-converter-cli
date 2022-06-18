package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("invalid arguments")
var errReadingInput = errors.New("error reading input")

func main() {
	// validate args
	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}
	// ensure consistency
	originUnit = strings.ToUpper(os.Args[1])
	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")

		// read current temperature
		_, err = fmt.Scanln(&originValue)
		if err != nil {
			printError(errReadingInput)
		}

		// convert the temperature
		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")

		// prompt to convert again
		_, err := fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			printError(errReadingInput)
		}

		// pass prompt answer
		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
