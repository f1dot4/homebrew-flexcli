package commands

import (
	"fmt"
	"net/url"
	"time"
)

const dateFlagFormat = "2006-01-02"

// validateDateFlag ensures a --start-date/--end-date flag value is either
// empty or a valid YYYY-MM-DD date.
func validateDateFlag(value string, flagName string) error {
	if value == "" {
		return nil
	}
	if _, err := time.Parse(dateFlagFormat, value); err != nil {
		return fmt.Errorf("invalid %s %q: expected format YYYY-MM-DD", flagName, value)
	}
	return nil
}

// buildDateRangeQuery builds a "start_date=...&end_date=..." query string
// fragment from optional start/end date flag values, omitting empty ones.
func buildDateRangeQuery(startDate string, endDate string) string {
	values := url.Values{}
	if startDate != "" {
		values.Set("start_date", startDate)
	}
	if endDate != "" {
		values.Set("end_date", endDate)
	}
	return values.Encode()
}
