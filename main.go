package main

import (
	"fmt"
	"log"
	"os"

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

	// setup beta redirect
	Mux.Use(func(c *gin.Context) {
		if c.Request.RequestURI == "/beta" {
			c.SetCookie("beta", "true", 86700, "/", "", false, true)
			c.Redirect(302, "/")
		}
	})

	// setup the static index file
	// beta: Mux.StaticFile("/", "root/index.html")
	Mux.GET("/", func(c *gin.Context) {
		if beta, _ := c.Cookie("beta"); beta == "true" {
			c.File("root/betaok.html")
		} else {
			c.File("root/index.html")
		}
	})

	// setup any api routes
	v1 := Mux.Group("/api/v1")
	{
		v1.POST("/new", addDomain)
		v1.GET("/:domain", getDomain)
	}

	Mux.LoadHTMLGlob("tmpls/view*html")
	// catch everything else with a static server
	Mux.NoRoute(viewdomain)

	return nil
}

// GetBind Determine the interface on which to bind the webserver
func GetBind() string {
	bind := "0.0.0.0:8081"
	if os.Getenv("OPENSHIFT_GO_IP") != "" {
		bind = fmt.Sprintf("%s:%s",
			os.Getenv("OPENSHIFT_GO_IP"),
			os.Getenv("OPENSHIFT_GO_PORT"),
		)
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

	err = InitEmail()
	if err != nil {
		log.Println(err)
	}

	err = InitJobs()
	if err != nil {
		log.Println(err)
	}

	foo := new(AnalyticsEmails)
	foo.Run()

	err = Mux.Run(GetBind())
	if err != nil {
		panic(err)
	}
}
