package main

import (
	"errors"
	"fmt"
	"strings"
)

func readingSourceCurrency() (string, error) {
	// Получение исходной валюты от пользователя

	var sourceCurrency string

	fmt.Print("Исходная валюта (USD / EUR / RUB): ")
	fmt.Scan(&sourceCurrency)

	sourceCurrency = strings.ToUpper(strings.TrimSpace(sourceCurrency))

	if sourceCurrency != "USD" && sourceCurrency != "EUR" && sourceCurrency != "RUB" {
		return "", errors.New("не верно указана исходная валюта")
	}

	return sourceCurrency, nil
}

func readingQuantity(Currency string) (float64, error) {
	// Получение исходной валюты от пользователя

	var quantity float64

	fmt.Printf("Введите количество %s для конвертации: ", Currency)
	fmt.Scan(&quantity)

	if quantity <= 0 {
		return 0.0, fmt.Errorf("не верное количество %s", Currency)
	}

	return quantity, nil
}

func readingTargetCurrency(sourceCurrency string) (string, error) {
	var option1, option2 string

	switch sourceCurrency {
	case "USD":
		option1, option2 = "EUR", "RUB"
	case "EUR":
		option1, option2 = "USD", "RUB"
	case "RUB":
		option1, option2 = "USD", "EUR"
	default:
		return "", fmt.Errorf("неподдерживаемая валюта: %s", sourceCurrency)
	}

	fmt.Printf("Целевая валюта (%s / %s): ", option1, option2)

	var targetCurrency string
	fmt.Scan(&targetCurrency)

	targetCurrency = strings.ToUpper(strings.TrimSpace(targetCurrency))

	if targetCurrency != option1 && targetCurrency != option2 {
		return "", fmt.Errorf("неверно указана валюта: %s", targetCurrency)
	}

	return targetCurrency, nil
}

func convertCurrency(sourceCurrency string, quantity float64, targetCurrency string) (float64, error) {
	const UsdRub float64 = 80
	const UsdEur float64 = 0.9
	EurRub := (1 / UsdEur) * UsdRub

	switch {
	case sourceCurrency == "USD" && targetCurrency == "RUB":
		return quantity * UsdRub, nil
	case sourceCurrency == "RUB" && targetCurrency == "USD":
		return quantity / UsdRub, nil
	case sourceCurrency == "USD" && targetCurrency == "EUR":
		return quantity * UsdEur, nil
	case sourceCurrency == "EUR" && targetCurrency == "USD":
		return quantity / UsdEur, nil
	case sourceCurrency == "EUR" && targetCurrency == "RUB":
		return quantity * EurRub, nil
	case sourceCurrency == "RUB" && targetCurrency == "EUR":
		return quantity / EurRub, nil
	default:
		return 0, fmt.Errorf("неподдерживаемая валютная пара: %s -> %s", sourceCurrency, targetCurrency)
	}
}

func main() {
	fmt.Println("__ Конвертер валют __")

	for {

		// Получение исходной валюты от пользователя
		sourceCurrency, err := readingSourceCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}

		quantity, err := readingQuantity(sourceCurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}

		targetCurrency, err := readingTargetCurrency(sourceCurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := convertCurrency(sourceCurrency, quantity, targetCurrency)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Результат конвертации %.2f %s -> %.2f %s\n", quantity, sourceCurrency, result, targetCurrency)

	}

}
