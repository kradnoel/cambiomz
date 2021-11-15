package types

import (
	"fmt"
	"net/http"

	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/kradnoel/cambiomz/pkg/util"
)

type entityValue struct {
	Entity   string
	Currency currencyValue
}

type currencyValue struct {
	Country  string
	Currency string
	Buy      string
	Sell     string
}

type ExchangeRates struct{}

func NewCurrency() currencyValue {
	return currencyValue{}
}

func NewEntityValue() entityValue {
	return entityValue{}
}

func NewCurrencyValue() currencyValue {
	return currencyValue{}
}

func (e *ExchangeRates) LoadMilleniumBimRates() ([]currencyValue, error) {
	cambio := []currencyValue{}
	bimUrl := os.Getenv("BIM_URL")

	if bimUrl == "" {
		log.Fatal("$BIM_URL must be set")
	}

	res, err := http.Get(bimUrl)
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

	doc.Find(".rates-card-one .rates-values .values").Each(func(i int, s *goquery.Selection) {
		children := s.Find("span")
		tempCurrency := NewCurrency()

		for j := range children.Nodes {
			span := children.Eq(j).Text()

			// currency
			if j == 0 {
				tempCurrency.Currency = span
			}

			// buy
			if j == 1 {
				tempCurrency.Buy = util.FormattedCurrency(span, ",")
			}

			// sell
			if j == 2 {
				tempCurrency.Sell = util.FormattedCurrency(span, ",")
			}
		}
		cambio = append(cambio, tempCurrency)
	})
	return cambio, nil
}

func (e *ExchangeRates) LoadBCIRates() ([]currencyValue, error) {
	cambio := []currencyValue{}
	bciUrl := os.Getenv("BCI_URL")

	if bciUrl == "" {
		log.Fatal("$BCI_URL must be set")
	}

	res, err := http.Get(bciUrl)
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
		children := s.Find("tr")
		tempCurrency := NewCurrency()

		for j := range children.Nodes {

			children.Eq(j).Find("td").Each(func(k int, s1 *goquery.Selection) {

				td := s1.Text()

				// country
				if k == 0 {
					tempCurrency.Country = td

				}

				// currency
				if k == 1 {
					tempCurrency.Currency = td
				}

				// buy
				if k == 2 {
					tempCurrency.Buy = util.FormattedCurrency(td, ",")
				}

				// sell
				if k == 3 {
					tempCurrency.Sell = util.FormattedCurrency(td, ",")
				}
			})

			cambio = append(cambio, tempCurrency)
		}
	})

	return cambio, nil
}

func (e *ExchangeRates) LoadStandardBankRates() ([]currencyValue, error) {
	cambio := []currencyValue{}
	standardBankUrl := os.Getenv("STANDARD_URL")

	if standardBankUrl == "" {
		log.Fatal("$STANDARD_URL must be set")
	}

	res, err := http.Get(standardBankUrl)
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

	doc.Find("div.class-currency-exchange tbody").Each(func(i int, s *goquery.Selection) {
		children := s.Find("tr")
		tempCurrency := NewCurrency()

		for j := range children.Nodes {
			children.Eq(j).Find("td").Each(func(k int, s1 *goquery.Selection) {

				td := s1.Text()

				if td != "" {
					// currency
					if k == 0 {
						tempCurrency.Currency = td
						fmt.Println(td)
					}

					// buy
					if k == 1 {
						tempCurrency.Buy = util.FormattedCurrency(td, ",")
						fmt.Println(td)
					}

					// sell
					if k == 2 {
						tempCurrency.Sell = util.FormattedCurrency(td, ",")
						fmt.Println(td)
					}
				}
			})

			cambio = append(cambio, tempCurrency)
		}
	})

	return cambio, nil
}

func (e *ExchangeRates) LoadCurrencies() ([]currencyValue, error) {
	countries := make(map[int]string)
	currencies := make(map[int]string)
	buy := make(map[int]string)
	sell := make(map[int]string)
	cambio := []currencyValue{}

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
		cambio = append(cambio, currencyValue{Country: countries[y], Currency: currencies[y], Buy: buy[y], Sell: sell[y]})
	}

	return cambio, nil
}

func (e *ExchangeRates) LoadFakeCurrencies() ([]currencyValue, error) {
	cambio := []currencyValue{
		{Country: "USA", Currency: "USD", Buy: "61.04", Sell: "61.02"},
		{Country: "South Africa", Currency: "ZAR", Buy: "4.25", Sell: "4.35"},
	}
	return cambio, nil
}
