package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Mux is our main gin engine
var Mux *gin.Engine

// InitEngine sets up gin
func InitEngine() error {

	// first, use the default logging and reporting
	Mux = gin.Default()

	// next, setup our middlewares
	Mux.Use(Analytics)

	// setup any api routes
	Mux.GET("/api/v1", apiDomain)

	// setup the static index file
	Mux.StaticFile("/", "root/index.html")

	// catch everything else with a static server
	Mux.NoRoute(static.ServeRoot("/", "root"))

	return nil
}

func GetBind() string {
	bind := "0.0.0.0:8080"
	if os.Getenv("OPENSHIFT_GO_IP") != "" {
		bind = fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	}
	fmt.Printf("listening on %s...\n", bind)
	return bind
}

func main() {

	err := InitEngine()
	if err != nil {
		panic(err)
	}

	err = InitDatabase()
	if err != nil {
		panic(err)
	}

	err = Mux.Run(GetBind())
	if err != nil {
		panic(err)
	}
}
