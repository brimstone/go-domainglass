package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var Mux *gin.Engine

func InitEngine() {

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
}

func main() {

	InitEngine()

	InitDatabase()

	bind := "0.0.0.0:8080"
	if os.Getenv("OPENSHIFT_GO_IP") != "" {
		gin.SetMode(gin.ReleaseMode)
		bind = fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	}
	fmt.Printf("listening on %s...", bind)
	err := Mux.Run(bind)
	if err != nil {
		panic(err)
	}
}
