package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/iot-app-backend/utils"
)

func HandleFetchedData(app *pocketbase.PocketBase) {
	defer utils.RecoverAfterPanic(HandleFetchedData)
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			return nil
		})
		return nil
	})
}
