package main

import (
	logger "github.com/JotaCodePro/logs/logger"
)

func main() {

	logger := logger.NewLogger(2, "DarkGray", "Cyan", "")
	logger.Info("Prueba INFO")
	logger.Warning("Prueba Warning")
	logger.Error("Prueba error")
}
