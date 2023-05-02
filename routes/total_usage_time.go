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
			return c.JSON(200, map[string]interface{}{
				"total_usage_time": daos.GetTotalUsageTime(app, id, period),
			})
		}, apis.ActivityLogger(app))
		return nil
	})
}
