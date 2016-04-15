package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

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

	// setup beta redirect
	Mux.Use(func(c *gin.Context) {
		if c.Request.RequestURI == "/beta" {
			c.SetCookie("beta", "true", 86700, "/", "", false, true)
			c.Redirect(302, "/")
		}
	})

	// setup any api routes
	Mux.GET("/api/v1", apiDomain)

	// setup the static index file
	// beta: Mux.StaticFile("/", "root/index.html")
	Mux.GET("/", func(c *gin.Context) {
		if beta, _ := c.Cookie("beta"); beta == "true" {
			c.File("root/betaok.html")
		} else {
			c.File("root/index.html")
		}
	})

	Mux.POST("/", func(c *gin.Context) {
		domain := c.PostForm("domain")

		matched, _ := regexp.MatchString("^[a-z0-9._]*\\.[a-z]{2,}$", domain)
		if matched {
			c.Redirect(302, "/"+domain)
		} else {
			c.Redirect(302, "/")
		}
	})

	Mux.LoadHTMLGlob("tmpls/view*html")
	// catch everything else with a static server
	Mux.NoRoute(func(c *gin.Context) {
		matched, _ := regexp.MatchString("^/[a-z0-9._]*\\.[a-z]{2,}$", c.Request.RequestURI)
		if matched {
			if beta, _ := c.Cookie("beta"); beta != "true" {
				c.Redirect(301, "/")
				return
			}
			c.HTML(http.StatusOK, "viewdomain.html", gin.H{
				"domain": c.Request.RequestURI,
			})
			return
		}
		static.ServeRoot("/", "root")(c)
	})

	return nil
}

// GetBind Determine the interface on which to bind the webserver
func GetBind() string {
	bind := "0.0.0.0:8080"
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
