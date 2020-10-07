package main

import (
	"fmt"

	"github.com/amonoyflow/cozy_backend/cmd/config"
)

func main() {
	if err := config.LoadConfiguration(); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.Database.ConnectionURI)
	fmt.Println(config.Config.Server.Port)
}
