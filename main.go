package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/thaiha1607/iot-app-backend/hooks"
	"github.com/thaiha1607/iot-app-backend/routes"
)

func main() {
	godotenv.Load()
	app := pocketbase.New()
	routes.WrapAllRoutes(app)
	hooks.HandleFetchedData(app)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
