package main

import (
	"github.com/zsandibe/eff_mobile_task/internal/app"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
