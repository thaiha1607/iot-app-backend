package daos

import (
	"log"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

func GetTotalUsageTime(app *pocketbase.PocketBase, deviceId string, period string) int {
	records, err := app.Dao().FindRecordsByExpr("on_off_time", dbx.HashExp{"output_id": deviceId})
	if err != nil {
		log.Panicln(err)
	}
	var timeUsage []int
	var upperBound, lowerBound time.Time
	// For testing purpose
	// today, _ := time.Parse("2006-Jan-02", "2023-Mar-01")
	today := time.Now()
	switch period {
	case "1d":
		lowerBound = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
		upperBound = time.Date(today.Year(), today.Month(), today.Day(), 23, 59, 59, 0, today.Location())
	case "1w":
		lowerBound = time.Date(today.Year(), today.Month(), today.Day()-int(today.Weekday()), 0, 0, 0, 0, today.Location())
		upperBound = time.Date(today.Year(), today.Month(), today.Day()+int(6-today.Weekday()), 23, 59, 59, 0, today.Location())
	case "1m":
		lowerBound = time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location())
		upperBound = time.Date(today.Year(), today.Month()+1, 0, 23, 59, 59, 0, today.Location())
	case "1y":
		lowerBound = time.Date(today.Year(), 1, 1, 0, 0, 0, 0, today.Location())
		upperBound = time.Date(today.Year(), 12, 31, 23, 59, 59, 0, today.Location())
	default:
		return 0
	}
	for _, record := range records {
		onTime := record.GetDateTime("on_time").Time()
		offTime := record.GetDateTime("off_time").Time()
		if onTime.After(lowerBound) && onTime.Before(upperBound) {
			if offTime.IsZero() {
				continue
			}
			timeUsage = append(timeUsage, int(offTime.Sub(onTime).Seconds()))
		}
	}
	var totalUsageTime int
	for _, time := range timeUsage {
		totalUsageTime += time
	}
	return int(totalUsageTime / 60)
}
