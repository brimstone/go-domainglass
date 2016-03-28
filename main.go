package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var Mux *gin.Engine

func InitEngine() {
	Mux = gin.Default()
	Mux.GET("/api/v1", apiDomain)
	Mux.StaticFile("/", "root/index.html")
	Mux.NoRoute(static.ServeRoot("/", "root"))
}

func main() {

	InitEngine()

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
