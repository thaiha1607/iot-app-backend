package daos

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func GetPowerCapacity(app *pocketbase.PocketBase, deviceId string) int {
	var table string
	switch deviceId {
	case "1":
		table = "fan"
	case "2":
		table = "nebulizer"
	default:
		return 0
	}
	records, err := app.Dao().FindRecordsByExpr(table)
	if err != nil {
		log.Panicln(err)
	}
	for _, record := range records {
		return record.GetInt("power_capacity")
	}
	return 0
}
