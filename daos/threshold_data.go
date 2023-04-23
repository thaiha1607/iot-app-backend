package daos

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func UpdateThresholdData(d *daos.Dao, deviceType string, val int) {
	rawQuery := d.DB().NewQuery("UPDATE output_devices SET env_limit = {:value} WHERE type = {:type}")
	var completeQuery *dbx.Query
	switch deviceType {
	case "fan":
		completeQuery = rawQuery.Bind(dbx.Params{"type": "fan", "value": val})
	case "nebulizer":
		completeQuery = rawQuery.Bind(dbx.Params{"type": "nebulizer", "value": val})
	default:
		log.Panicln("Unknown sensor type: ", deviceType)
	}
	if _, err := completeQuery.Execute(); err != nil {
		log.Panicln(err)
	}
}
