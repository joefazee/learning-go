package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {

	now := time.Now()

	fmt.Printf("Current time (default Go format): %s\n", now)

	fmt.Printf("Formatted as YYYY-MM-DD: %s\n", now.Format("2006-01-02"))

	fmt.Printf("Formatted as MM/DD/YYYY hh:mm:ss PM: %s\n", now.Format("01/02/2006 03:04:05 PM"))

	fmt.Printf("Formatted as Day, Month Date, Year: %s\n", now.Format("Mon, Jan 2, 2006"))

	fmt.Printf("Formatted as RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Formatted as ANSIC: %s\n", now.Format(time.ANSIC))

	dateStr1 := "2025-07-15"
	layout1 := "2006-01-02"

	parsedTime1, err := time.Parse(layout1, dateStr1)
	if err != nil {
		fmt.Printf("Error parsing '%s' with layout '%s': %v\n", dateStr1, layout1, err)
	} else {
		fmt.Printf("Parsed '%s' -> Year: %d, Month: %s, Day: %d\n",
			dateStr1, parsedTime1.Year(), parsedTime1.Month(), parsedTime1.Day())
	}

	rfc3339Str := "2025-12-25T10:00:00+01:00"
	parsedTimeRFC, err := time.Parse(time.RFC3339, rfc3339Str)
	if err != nil {
		fmt.Printf("Error parsing RFC3339 string '%s': %v\n", rfc3339Str, err)
	} else {
		fmt.Printf("Parsed RFC3339: %s (Location: %s, Offset: %d seconds)\n",
			parsedTimeRFC, parsedTimeRFC.Location(), getOffsetInSeconds(parsedTimeRFC))
		fmt.Printf("Parsed RFC3339 (in UTC): %s\n", parsedTimeRFC.UTC())
	}

	fetch(http.Get, "google")

	fetch(func(url string) (*http.Response, error) {
		return http.ReadResponse(bufio.NewReader(strings.NewReader(url)), nil)
	}, "google.com")

}

type getFunc = func(url string) (*http.Response, error)

func fetch(getFunc getFunc, url string) {

	_, _ = getFunc(url)

}

func getOffsetInSeconds(t time.Time) int {
	_, offset := t.Zone()
	return offset
}
