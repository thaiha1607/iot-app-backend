package daos

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func InsertDeviceHistory(d *daos.Dao, deviceType string, deviceValue bool) {
	rawQuery := d.DB().NewQuery("INSERT INTO {:table} (is_turned_on) VALUES ({:value})")
	if deviceValue {
		rawQuery.Bind(dbx.Params{"value": 1})
	} else {
		rawQuery.Bind(dbx.Params{"value": 0})
	}
	var completeQuery *dbx.Query
	switch deviceType {
	case "fan":
		completeQuery = rawQuery.Bind(dbx.Params{"type": "fan"})
	case "nebulizer":
		completeQuery = rawQuery.Bind(dbx.Params{"type": "nebulizer"})
	default:
		log.Panicln("Unknown sensor type: ", deviceType)
	}
	if _, err := completeQuery.Execute(); err != nil {
		log.Panicln(err)
	}
}
