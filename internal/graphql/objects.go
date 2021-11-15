package graphql

import (
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/kradnoel/cambiomz/internal/types"
)

var exchangeRates = types.ExchangeRates{}

var currencyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Currency",
		Fields: graphql.Fields{
			"country":  &graphql.Field{Description: "Country origin", Type: graphql.String},
			"currency": &graphql.Field{Description: "Currency name", Type: graphql.String},
			"buy":      &graphql.Field{Description: "Buy price", Type: graphql.String},
			"sell":     &graphql.Field{Description: "Sell price", Type: graphql.String},
		},
	})

var Entity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Entity",
		Fields: graphql.Fields{
			"entity":   &graphql.Field{Description: "Entity names", Type: graphql.String},
			"currency": &graphql.Field{Description: "Entity names", Type: currencyType},
		},
	})

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"single": &graphql.Field{
				Type: Entity,
				Args: graphql.FieldConfigArgument{
					"entity": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"currency": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					entity := types.NewEntityValue()

					idEntity, isEntityOk := params.Args["entity"].(string)

					if isEntityOk {
						switch idEntity {
						case "bci":
							currencies, _ := exchangeRates.LoadBCIRates()
							currency := types.NewCurrencyValue()
							entity = types.NewEntityValue()
							entity.Entity = "bci"

							idCurrency, isCurrencyOk := params.Args["currency"].(string)

							if isCurrencyOk {
								for _, cur := range currencies {
									if cur.Currency == strings.ToUpper(idCurrency) {
										currency.Country = cur.Country
										currency.Currency = cur.Currency
										currency.Buy = cur.Buy
										currency.Sell = cur.Sell
										entity.Currency = currency
									}
								}
							}
						
						case "bim":
							currencies, _ := exchangeRates.LoadMilleniumBimRates()
							currency := types.NewCurrencyValue()
							entity = types.NewEntityValue()
							entity.Entity = "bim"

							idCurrency, isCurrencyOk := params.Args["currency"].(string)

							if isCurrencyOk {
								for _, cur := range currencies {
									if cur.Currency == strings.ToUpper(idCurrency) {
										currency.Country = cur.Country
										currency.Currency = cur.Currency
										currency.Buy = cur.Buy
										currency.Sell = cur.Sell
										entity.Currency = currency
									}
								}
							}
						
						case "bmo":
							currencies, _ := exchangeRates.LoadBMRates()
							currency := types.NewCurrencyValue()
							entity = types.NewEntityValue()
							entity.Entity = "bmo"

							idCurrency, isCurrencyOk := params.Args["currency"].(string)

							if isCurrencyOk {
								for _, cur := range currencies {
									if cur.Currency == strings.ToUpper(idCurrency) {
										currency.Country = cur.Country
										currency.Currency = cur.Currency
										currency.Buy = cur.Buy
										currency.Sell = cur.Sell
										entity.Currency = currency
									}
								}
							}

						case "sbk":
							currencies, _ := exchangeRates.LoadStandardBankRates()
							currency := types.NewCurrencyValue()
							entity = types.NewEntityValue()
							entity.Entity = "sbk"

							idCurrency, isCurrencyOk := params.Args["currency"].(string)

							if isCurrencyOk {
								for _, cur := range currencies {
									if cur.Currency == strings.ToUpper(idCurrency) {
										currency.Country = cur.Country
										currency.Currency = cur.Currency
										currency.Buy = cur.Buy
										currency.Sell = cur.Sell
										entity.Currency = currency
									}
								}
							}

						default:
							entity = types.NewEntityValue()
						}
					}

					return entity, nil
				},
			},
			"currencies": &graphql.Field{
				Type: graphql.NewList(currencyType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					currencies, _ := exchangeRates.LoadCurrencies()
					return currencies, nil
				},
			},
		},
	},
)
