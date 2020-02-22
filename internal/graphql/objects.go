package graphql

import (
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/kradnoel/cambiomz/internal/types"
)

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

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"currency": &graphql.Field{
				Type: currencyType,
				Args: graphql.FieldConfigArgument{
					"currency": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := params.Args["currency"].(string)
					//currencies, _ := types.LoadFakeCurrencies()
					currencies, _ := types.LoadCurrencies()
					currency := types.NewCurrency()
					if isOK {
						// Search for el with id
						for _, curr := range currencies {
							if curr.Currency == strings.ToUpper(idQuery) {
								return curr, nil
							}

						}
					}

					return currency, nil
				},
			},
			"currencies": &graphql.Field{
				Type: graphql.NewList(currencyType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//currencies, _ := types.LoadFakeCurrencies()
					currencies, _ := types.LoadCurrencies()
					return currencies, nil
				},
			},
		},
	},
)
