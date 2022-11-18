package panel

import (
	service "gold-panel/internal/service/v1"
)

func Launch() {
	service := service.NewService()
	RunBot(service)
}
