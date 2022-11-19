package main

import (
	"gold-panel/config"
	"gold-panel/internal/panel"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Panel running...")

	config.LoadConfigFrom("./config")
	config.LoadEnvFrom(".")

	panel.Launch()
}
