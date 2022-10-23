package main

import (
	configs "final_project/database"
	"final_project/routers"
)

func main() {
	configs.StartDB()

	routers.Start().Run("localhost:8001")
}
