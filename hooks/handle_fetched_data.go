package hooks

import (
	"log"
	"strconv"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/thaiha1607/iot-app-backend/config"
	"github.com/thaiha1607/iot-app-backend/daos"
	"github.com/thaiha1607/iot-app-backend/services"
	"github.com/thaiha1607/iot-app-backend/utils"
)

func fetchAndInsertTemperatureAndHumidityData(app *pocketbase.PocketBase) {
	var tempObj services.FeedDataType
	var humObj services.FeedDataType
	services.FetchDataFromFeed(config.TemperatureFeedName, &tempObj)
	services.FetchDataFromFeed(config.HumidityFeedName, &humObj)
	tempVal, _ := strconv.ParseFloat(tempObj[0].Value, 32)
	humVal, _ := strconv.ParseFloat(humObj[0].Value, 32)
	createdAtVal, _ := time.Parse(time.RFC3339, tempObj[0].CreatedAt)
	err := daos.InsertTemperatureAndHumidityData(app, int(humVal), int(tempVal), createdAtVal)
	if err != nil {
		log.Panicln(err)
	}
}

func fetchAndInsertInfraredData(app *pocketbase.PocketBase) {
	var infraredObj services.FeedDataType
	services.FetchDataFromFeed(config.InfraredFeedName, &infraredObj)
	infraredVal, _ := strconv.ParseBool(infraredObj[0].Value)
	createdAtVal, _ := time.Parse(time.RFC3339, infraredObj[0].CreatedAt)
	err := daos.InsertInfraredData(app, infraredVal, createdAtVal)
	if err != nil {
		log.Panicln(err)
	}
}

func fetchAndInsertFanHistory(app *pocketbase.PocketBase) {
	var fanObj services.FeedDataType
	services.FetchDataFromFeed(config.FanFeedName, &fanObj)
	fanVal, _ := strconv.ParseBool(fanObj[0].Value)
	createdAtVal, _ := time.Parse(time.RFC3339, fanObj[0].CreatedAt)
	err := daos.InsertDeviceHistory(app, "fan", fanVal, createdAtVal)
	if err != nil {
		log.Panicln(err)
	}
}

func fetchAndInsertNebulizerHistory(app *pocketbase.PocketBase) {
	var nebulizerObj services.FeedDataType
	services.FetchDataFromFeed(config.NebulizerFeedName, &nebulizerObj)
	nebulizerVal, _ := strconv.ParseBool(nebulizerObj[0].Value)
	createdAtVal, _ := time.Parse(time.RFC3339, nebulizerObj[0].CreatedAt)
	err := daos.InsertDeviceHistory(app, "nebulizer", nebulizerVal, createdAtVal)
	if err != nil {
		log.Panicln(err)
	}
}

func fetchAndInsertTempThresholdData(app *pocketbase.PocketBase) {
	var tempThresholdObj services.FeedDataType
	services.FetchDataFromFeed(config.TemperatureThresholdFeedName, &tempThresholdObj)
	tempThresholdVal, _ := strconv.Atoi(tempThresholdObj[0].Value)
	createdAtVal, _ := time.Parse(time.RFC3339, tempThresholdObj[0].CreatedAt)
	err := daos.UpdateThresholdData(app, "fan", tempThresholdVal, createdAtVal)
	if err != nil {
		log.Panicln(err)
	}
}

func fetchAndInsertHumidThresholdData(app *pocketbase.PocketBase) {
	var humidThresholdData services.FeedDataType
	services.FetchDataFromFeed(config.HumidityThresholdFeedName, &humidThresholdData)
	humidThresholdVal, _ := strconv.Atoi(humidThresholdData[0].Value)
	createdAtVal, _ := time.Parse(time.RFC3339, humidThresholdData[0].CreatedAt)
	err := daos.UpdateThresholdData(app, "nebulizer", humidThresholdVal, createdAtVal)
	if err != nil {
		log.Panicln(err)
	}
}

func HandleFetchedData(app *pocketbase.PocketBase) {
	defer utils.RecoverAfterPanic(HandleFetchedData)
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		utils.DoRepetitiveTask(app, fetchAndInsertTemperatureAndHumidityData, 30)
		utils.DoRepetitiveTask(app, fetchAndInsertInfraredData, 10)
		utils.DoRepetitiveTask(app, fetchAndInsertFanHistory, 10)
		utils.DoRepetitiveTask(app, fetchAndInsertNebulizerHistory, 10)
		utils.DoRepetitiveTask(app, fetchAndInsertTempThresholdData, 30)
		utils.DoRepetitiveTask(app, fetchAndInsertHumidThresholdData, 30)
		return nil
	})
}
