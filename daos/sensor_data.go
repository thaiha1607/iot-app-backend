package daos

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func RetrieveLatestSensorData(d *daos.Dao, sensorType string) any {
	var sensorData any
	rawQuery := d.DB().NewQuery("SELECT * FROM {:table} ORDER BY created_at DESC LIMIT 1")
	var completeQuery *dbx.Query
	switch sensorType {
	case "temperature":
		completeQuery = rawQuery.Bind(dbx.Params{"table": "temp_sensor"})
	case "humidity":
		completeQuery = rawQuery.Bind(dbx.Params{"table": "humid_sensor"})
	case "infrared":
		completeQuery = rawQuery.Bind(dbx.Params{"table": "ir_sensor"})
	default:
		log.Panicln("Unknown sensor type: ", sensorType)
	}
	if err := completeQuery.One(&sensorData); err != nil {
		log.Panicln(err)
	}
	return sensorData
}

func InsertDataToSensorTable(d *daos.Dao, sensorType string, sensorValue int) {
	rawQuery := d.DB().NewQuery("INSERT INTO {:table} (value) VALUES ({:value})")
	var completeQuery *dbx.Query
	switch sensorType {
	case "temperature":
		completeQuery = rawQuery.Bind(dbx.Params{"table": "temp_sensor", "value": sensorValue})
	case "humidity":
		completeQuery = rawQuery.Bind(dbx.Params{"table": "humid_sensor", "value": sensorValue})
	case "infrared":
		completeQuery = rawQuery.Bind(dbx.Params{"table": "ir_sensor", "value": sensorValue})
	default:
		log.Panicln("Unknown sensor type: ", sensorType)
	}
	if _, err := completeQuery.Execute(); err != nil {
		log.Panicln(err)
	}
}
