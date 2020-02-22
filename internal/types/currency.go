package types

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type currency struct {
	Country  string
	Currency string
	Buy      string
	Sell     string
}

func NewCurrency() currency {
	return currency{}
}

func CurrencyWithValues(_country string, _currency string, _buy string, _sell string) currency {
	return currency{
		Country:  _country,
		Currency: _currency,
		Buy:      _buy,
		Sell:     _sell,
	}
}

func LoadCurrencies() ([]currency, error) {
	countries := make(map[int]string)
	currencies := make(map[int]string)
	buy := make(map[int]string)
	sell := make(map[int]string)
	cambio := []currency{}

	res, err := http.Get("https://www.bci.co.mz/cambio/")
	if err != nil {
		return cambio, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return cambio, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return cambio, err
	}

	doc.Find("table#cambio-table .cambio_total").Each(func(i int, s *goquery.Selection) {
		var a int = 0
		var b int = 1
		var c int = 2
		var d int = 3

		size := doc.Find("table#cambio-table .cambio_total tr").Size()
		sel := doc.Find("table#cambio-table .cambio_total tr td")

		for i := range sel.Nodes {
			td := sel.Eq(i).Text()

			if i == a {
				countries[13-size] = td
				a = a + 4
			}

			if i == b {
				currencies[12-size] = td
				b = b + 4
			}

			if i == c {
				buy[12-size] = td
				c = c + 4
			}

			if i == d {
				sell[12-size] = td
				d = d + 4
			}

			if i%4 == 0 {
				size = size - 1
			}
		}
	})

	for y := 0; y < 12; y++ {
		cambio = append(cambio, currency{Country: countries[y], Currency: currencies[y], Buy: buy[y], Sell: sell[y]})
	}

	return cambio, nil
}

func LoadFakeCurrencies() ([]currency, error) {
	cambio := []currency{
		currency{Country: "USA", Currency: "USD", Buy: "61.04", Sell: "61.02"},
		currency{Country: "South Africa", Currency: "ZAR", Buy: "4.25", Sell: "4.35"},
	}
	return cambio, nil
}
