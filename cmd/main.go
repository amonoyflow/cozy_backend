package main

import (
	"fmt"

	"github.com/amonoyflow/cozy_backend/cmd/cozy/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load application configuration
	if err := config.LoadConfiguration(); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Loogle middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	r.Use(gin.Recovery())
}
