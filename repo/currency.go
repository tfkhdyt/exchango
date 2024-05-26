package repo

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/go-zoox/fetch"

	"github.com/tfkhdyt/exchango/model"
)

func FindAllCurrencies() ([]model.Currency, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("failed to read user cache dir: %w", err)
	}

	cacheFile := path.Join(path.Join(cacheDir, "exchango"), fmt.Sprintf("currencies-%s.json", strings.ToLower(time.Now().Month().String())))

	currencies := make([]model.Currency, 0)

	if _, err := os.ReadFile(cacheFile); os.IsNotExist(err) {
		response, err := fetch.Get("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.min.json")
		if err != nil {
			fmt.Printf("failed to fetch list of currencies: %v\n", err)
			os.Exit(1)
		}

		data, errJson := response.JSON()
		if errJson != nil {
			fmt.Printf("failed to get JSON from API response body: %v\n", errJson)
			os.Exit(1)
		}

		var currenciesMap map[string]string
		if err := json.Unmarshal([]byte(data), &currenciesMap); err != nil {
			fmt.Printf("failed to unmarshal API response body: %v\n", err)
			os.Exit(1)
		}

		for code, name := range currenciesMap {
			currencies = append(currencies, model.Currency{
				Code: strings.ToUpper(code),
				Name: name,
			})
		}

		sort.Slice(currencies, func(i, j int) bool {
			return currencies[i].Code < currencies[j].Code
		})

		newCache, err := json.Marshal(currencies)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal currencies cache: %w", err)
		}

		if err := os.WriteFile(cacheFile, newCache, 0644); err != nil {
			return nil, fmt.Errorf("failed to write currencies cache: %w", err)
		}
	} else {
		data, err := os.ReadFile(cacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read data from cache: %w", err)
		}

		if err := json.Unmarshal(data, &currencies); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cache: %w", err)
		}
	}

	return currencies, nil
}
