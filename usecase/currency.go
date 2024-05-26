package usecase

import (
	"fmt"

	"github.com/tfkhdyt/exchango/model"
)

func PrintCurrencies(currencies []model.Currency, unknown bool) {
	for _, currency := range currencies {
		if currency.Name == "" {
			if !unknown {
				continue
			}

			currency.Name = "<null>"
		}
		fmt.Printf("%s = %v\n", currency.Code, currency.Name)
	}
}
