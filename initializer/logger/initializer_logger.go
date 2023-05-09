package logger

import (
	"fmt"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
)

var (
	logger *config.Logger
)

func InitializerLogger() {
	fmt.Println("Inicializer Logger Application")
	logger = config.GetLogger("Initializer loggers")
}
