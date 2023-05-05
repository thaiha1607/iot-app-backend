package daos

import (
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/thaiha1607/iot-app-backend/config"
)

func InsertTemperatureAndHumidityData(app *pocketbase.PocketBase, humidVal int, tempVal int, createdAt time.Time) error {
	records, err := app.Dao().FindRecordsByExpr("sensor", dbx.HashExp{"time": createdAt.Format(config.DefaultDateLayout)})
	if err != nil {
		return err
	}
	if len(records) > 0 {
		return nil
	}
	collection, err := app.Dao().FindCollectionByNameOrId("sensor")
	if err != nil {
		return err
	}
	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)
	form.LoadData(map[string]any{
		"humi_value": humidVal,
		"temp_value": tempVal,
		"time":       createdAt,
	})
	if err := form.Submit(); err != nil {
		return err
	}
	return nil
}

func InsertInfraredData(app *pocketbase.PocketBase, infraredVal bool, createdAt time.Time) error {
	records, err := app.Dao().FindRecordsByExpr("infrared_sensor", dbx.HashExp{"time": createdAt.Format(config.DefaultDateLayout)})
	if err != nil {
		return err
	}
	if len(records) > 0 {
		return nil
	}
	collection, err := app.Dao().FindCollectionByNameOrId("infrared_sensor")
	if err != nil {
		return err
	}
	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)
	form.LoadData(map[string]any{
		"infrared": infraredVal,
		"time":     createdAt,
	})
	if err := form.Submit(); err != nil {
		return err
	}
	return nil
}
