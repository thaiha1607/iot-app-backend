package daos

import (
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
)

func UpdateThresholdData(app *pocketbase.PocketBase, deviceType string, val int, createdAt time.Time) error {
	var prefixStr string
	if deviceType == "fan" {
		prefixStr = "temp_"
	} else {
		prefixStr = "humi_"
	}
	record, err := app.Dao().FindFirstRecordByData(deviceType, prefixStr+"id", 1)
	if err != nil {
		return err
	}
	form := forms.NewRecordUpsert(app, record)
	form.LoadData(map[string]any{
		prefixStr + "threshold":  val,
		"threshold_setting_time": createdAt,
	})
	if err := form.Submit(); err != nil {
		return err
	}
	return nil
}
