package main

import (
	"marketplace/api"
	db_component "marketplace/internal/db-component"
)

func main() {
	ctx, rdb := db_component.ConnectToRedis()
	defer rdb.Close()

	api := api.NewAPI(ctx, rdb)

	api.Run()
}
