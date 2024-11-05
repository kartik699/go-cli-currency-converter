package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

var (
    currencyFrom int
    currencyTo int
    amount string
)

var optionToCurrency = map[int]string {
    0: "USD",
    1: "EUR",
    2: "JPY",
    3: "INR",
}

func main() {
    var amountGroup *huh.Group = huh.NewGroup(
        huh.NewInput().
        Title("Enter the amount").
        Prompt("-> ").
        Validate(func(s string) error {
            amt, _ := convertToFloat(s)
            if amt == 0 {
                return errors.New("amount should be greater than zero")
            } else if amt < 0 {
                return errors.New("amount cannot be negative")
            }

            return nil
        }).
        Value(&amount),
    )
    
    var currencyGroup *huh.Group = huh.NewGroup(
        huh.NewSelect[int]().
        Title("Choose the currency your amount is in").
        Options(
            huh.NewOption("USD", 0),
            huh.NewOption("EUR", 1),
            huh.NewOption("JPY", 2),
            huh.NewOption("INR", 3),
        ).
        Value(&currencyFrom),

        huh.NewSelect[int]().
        Title("Choose the currency you want to convert your amount to").
        Options(
            huh.NewOption("USD", 0),
            huh.NewOption("EUR", 1),
            huh.NewOption("JPY", 2),
            huh.NewOption("INR", 3),
        ).
        Value(&currencyTo),
    )

    form := huh.NewForm(amountGroup, currencyGroup).WithTheme(huh.ThemeCharm())

    err := form.Run()

    if err != nil {
        fmt.Printf("An error occurred - %s", err)
    }

    amountInFloat, _ := convertToFloat(amount)

    title := fmt.Sprintf("Converting to %s...", optionToCurrency[currencyTo])

    err = spinner.New().Title(title).
    Action(func() {
        GetCurrencyRate(optionToCurrency[currencyFrom], optionToCurrency[currencyTo], amountInFloat)
    }).
    Run()

    if err != nil {
        log.Fatal("Something went wrong!")
    }
}

func convertToFloat(amount string) (float64, error) {
    res, err := strconv.ParseFloat(amount, 32)

    if err != nil {
        return 0, err
    }

    return res, nil
}
