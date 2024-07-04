package main

import (
	"github.com/zsandibe/eff_mobile_task/internal/app"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

// @title Effective mobile API
// @version 1.0
// @description This is basic server for a user searching service
// @host localhost:8888
// @BasePath /api/v1
func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
