package routes

import "github.com/pocketbase/pocketbase"

func WrapAllRoutes(app *pocketbase.PocketBase) {
	RegisterTotalUsageTimeRoute(app)
}
