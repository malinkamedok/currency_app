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

func (i InfoUseCase) GetCurrencyRate(currencyCode string, date string) (map[string]float64, error) {

	dateFormatted, err := parseAndFormatDate(date)
	if err != nil {
		return nil, err
	}

	req, err := i.cbrf.InitRequest(dateFormatted)
	if err != nil {
		return nil, err
	}

	resp, err := i.cbrf.SendRequest(req)
	if err != nil {
		return nil, err
	}

	rates, err := i.cbrf.DecodeResponse(resp)
	if err != nil {
		return nil, err
	}

	currencyResult := make(map[string]float64)

	if currencyCode == "" {
		for _, v := range rates.Valutes {
			currencyRate, err := i.cbrf.FindCurrencyRate(v.CharCode, rates)
			if err != nil {
				return nil, err
			}

			currencyResult[v.CharCode] = currencyRate
		}
	} else {
		currencyCode = strings.ToUpper(currencyCode)

		correct := checkCurrencyCorrect(currencyCode)
		if !correct {
			return nil, fmt.Errorf("incorrect currency code")
		}

		currencyRate, err := i.cbrf.FindCurrencyRate(currencyCode, rates)
		if err != nil {
			return nil, err
		}

		currencyResult[currencyCode] = currencyRate
	}

	return currencyResult, nil
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
