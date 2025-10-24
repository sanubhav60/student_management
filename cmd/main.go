package main

import (
	"student_management/internal/config"
	"student_management/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)

}
