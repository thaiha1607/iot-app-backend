package utils

import (
	"time"

	"github.com/pocketbase/pocketbase"
)

// Run a function every n seconds. Returns a channel that can be used to stop the task.
func DoRepetitiveTask(app *pocketbase.PocketBase, f func(app *pocketbase.PocketBase), seconds int) chan struct{} {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				f(app)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return quit
}
