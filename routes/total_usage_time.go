package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/thaiha1607/iot-app-backend/daos"
)

// GET /total-usage-time/:deviceId?period=1d|1w|1m|1y"
func RegisterTotalUsageTimeRoute(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/total-usage-time/:deviceId", func(c echo.Context) error {
			id := c.PathParam("deviceId")
			period := c.QueryParam("period")
			totalUsageTime := daos.GetTotalUsageTime(app, id, period)
			return c.JSON(200, map[string]any{
				"total_usage_time": totalUsageTime,
				// Calc power consumption in kWh with totalUsageTime in minutes and power_capacity in W
				"power_consumption": float64(totalUsageTime) * float64(daos.GetPowerCapacity(app, id)) / 1000 / 60,
			})
		}, apis.ActivityLogger(app))
		return nil
	})
}
