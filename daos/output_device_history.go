package daos

import (
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/thaiha1607/iot-app-backend/config"
)

func InsertDeviceHistory(app *pocketbase.PocketBase, deviceType string, deviceValue bool, createdAt time.Time) error {
	var deviceId int
	if deviceType == "fan" {
		deviceId = 1
	} else {
		deviceId = 2
	}
	records, err := app.Dao().FindRecordsByExpr("on_off_time", dbx.HashExp{"output_id": deviceId, "on_time": createdAt.Format(config.DefaultDateLayout)})
	if err != nil {
		return err
	}
	if len(records) > 0 {
		return nil
	}
	if !deviceValue {
		updateExistingRecord(app, deviceType, createdAt)
	}
	collection, err := app.Dao().FindCollectionByNameOrId("on_off_time")
	if err != nil {
		return err
	}
	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)
	form.LoadData(map[string]any{
		"output_id": deviceId,
		"on_time":   createdAt,
	})
	if err := form.Submit(); err != nil {
		return err
	}
	return nil
}

func updateExistingRecord(app *pocketbase.PocketBase, deviceType string, created_at time.Time) error {
	var deviceId int
	if deviceType == "fan" {
		deviceId = 1
	} else {
		deviceId = 2
	}
	records, err := app.Dao().FindRecordsByExpr("on_off_time", dbx.HashExp{"output_id": deviceId, "off_time": nil})
	if err != nil {
		return err
	}
	for _, record := range records {
		form := forms.NewRecordUpsert(app, record)
		form.LoadData(map[string]any{
			"off_time": created_at,
		})
		if err := form.Submit(); err != nil {
			return err
		}
	}
	return nil
}
