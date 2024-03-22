package usecase

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"time"
)

type InfoUseCase struct {
	cbrf InfoReq
}

var _ InfoContract = (*InfoUseCase)(nil)

func NewInfoUseCase(cbrf InfoReq) *InfoUseCase {
	return &InfoUseCase{cbrf: cbrf}
}

func (i InfoUseCase) GetCurrencyRate(currency string, date string) (string, error) {
	currency = strings.ToUpper(currency)

	correct := checkCurrencyCorrect(currency)
	if !correct {
		return "", fmt.Errorf("incorrect currency code")
	}

	dateFormatted, err := parseAndFormatDate(date)
	if err != nil {
		return "", err
	}

	req, err := i.cbrf.InitRequest(dateFormatted)
	if err != nil {
		return "", err
	}

	resp, err := i.cbrf.SendRequest(req)
	if err != nil {
		return "", err
	}

	rates, err := i.cbrf.DecodeResponse(resp)
	if err != nil {
		return "", err
	}

	currencyRate, err := i.cbrf.FindCurrencyRate(currency, rates)
	if err != nil {
		return "", err
	}

	return currencyRate, nil
}

func checkCurrencyCorrect(currency string) bool {
	correctCodes := []string{"USD", "AUD", "AZN", "GBP", "AMD", "BYN", "BGN", "BRL", "HUF", "VND", "HKD", "GEL", "DKK", "AED", "EUR", "EGP", "INR", "IDR", "KZT", "CAD", "QAR", "KGS", "CNY", "MDL", "NZD", "NOK", "PLN", "RON", "XDR", "SGD", "TJS", "THB", "TRY", "TMT", "UZS", "UAH", "CZK", "SEK", "CHF", "RSD", "ZAR", "KRW", "JPY"}
	if !slices.Contains(correctCodes, currency) {
		log.Printf("Currency code is incorrect")
		return false
	}
	return true
}

func parseAndFormatDate(date string) (string, error) {
	var dateFormatted string

	if date == "" {
		dateFormatted = time.Now().Format("02/01/2006")
	} else {
		dateParsed, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Printf("Error in parsing date")
			return "", err
		}

		if dateParsed.After(time.Now()) {
			return "", fmt.Errorf("incorrect time")
		}

		dateFormatted = dateParsed.Format("02/01/2006")
	}

	return dateFormatted, nil
}
