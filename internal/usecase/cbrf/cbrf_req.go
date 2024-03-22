package cbrf

import (
	"bytes"
	"devops_course_app/internal/entity"
	"devops_course_app/internal/usecase"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"log"
	"net/http"
)

type InfoCBRF struct{}

func NewInfoReq() *InfoCBRF {
	return &InfoCBRF{}
}

func (i InfoCBRF) InitRequest(dateFormatted string) (*http.Request, error) {
	url := "https://www.cbr.ru/scripts/XML_daily.asp?date_req=" + dateFormatted

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error in creating request")
		return nil, err
	}

	// Getting 403 Forbidden error without setting this header
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	return req, nil
}

func (i InfoCBRF) SendRequest(r *http.Request) (*http.Response, error) {
	c := http.Client{}

	resp, err := c.Do(r)
	if err != nil {
		log.Printf("Error in sending request")
		return nil, err
	}

	return resp, nil
}

func (i InfoCBRF) DecodeResponse(response *http.Response) (*entity.ValCurs, error) {
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error in reading response")
		return nil, err
	}

	reader := bytes.NewReader(responseData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	rates := new(entity.ValCurs)

	err = decoder.Decode(rates)
	if err != nil {
		log.Printf("Error in decoding response")
		return nil, err
	}

	return rates, nil
}

func (i InfoCBRF) FindCurrencyRate(currency string, currencyRates *entity.ValCurs) (string, error) {
	for _, v := range currencyRates.Valutes {
		if v.CharCode == currency {
			return v.Value, nil
		}
	}
	return "", fmt.Errorf("Currency or rate not found")
}

var _ usecase.InfoReq = (*InfoCBRF)(nil)
