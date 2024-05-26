package repo

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-zoox/fetch"
)

func GetRate(from, to string) (float64, error) {
	if err := VerifyCurrencyAvailability(from); err != nil {
		return 0, fmt.Errorf("%w (%s)", err, "base")
	}

	if err := VerifyCurrencyAvailability(to); err != nil {
		return 0, fmt.Errorf("%w (%s)", err, "target")
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return 0, fmt.Errorf("failed to read user cache dir: %w", err)
	}

	cacheFile := path.Join(path.Join(cacheDir, "exchango"), fmt.Sprintf("%s-%v.json", from, time.Now().Format(time.DateOnly)))

	var rate map[string]any
	if _, err := os.ReadFile(cacheFile); os.IsNotExist(err) {
		resp, err := fetch.Get(fmt.Sprintf("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.min.json", strings.ToLower(from)))
		if err != nil {
			return 0, fmt.Errorf("failed to fetch list of currencies: %w", err)
		}

		var baseCurrencies map[string]any
		data, errJson := resp.JSON()
		if errJson != nil {
			return 0, fmt.Errorf("failed to get JSON from API response body: %w", err)
		}

		if err := json.Unmarshal([]byte(data), &baseCurrencies); err != nil {
			return 0, fmt.Errorf("failed to unmarshal API response body: %w", err)
		}

		var ok bool
		rate, ok = baseCurrencies[strings.ToLower(from)].(map[string]any)
		if !ok {
			return 0, fmt.Errorf("API error: invalid response body")
		}

		newCache, err := json.Marshal(rate)
		if err != nil {
			return 0, fmt.Errorf("failed to marshal rate cache: %w", err)
		}

		if err := os.WriteFile(cacheFile, newCache, 0644); err != nil {
			return 0, fmt.Errorf("failed to write rate cache: %w", err)
		}
	} else {
		data, err := os.ReadFile(cacheFile)
		if err != nil {
			return 0, fmt.Errorf("failed to read rate from cache: %w", err)
		}

		if err := json.Unmarshal(data, &rate); err != nil {
			return 0, fmt.Errorf("failed to unmarshal API response body: %w", err)
		}
	}

	for code, rateVal := range rate {
		rateVal_, ok := rateVal.(float64)
		if !ok {
			return 0, fmt.Errorf("API error: invalid currency rate")
		}
		if code == strings.ToLower(to) {
			return rateVal_, nil
		}
	}

	return 0, fmt.Errorf("target currency is not found")
}
