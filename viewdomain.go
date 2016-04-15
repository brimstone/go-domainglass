package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func viewdomain(c *gin.Context) {
	// if we're not in the beta, kick them to the front page
	if beta, _ := c.Cookie("beta"); beta != "true" {
		c.Redirect(302, "/")
		return
	}

	// if we don't have a suburl that looks like a domain
	matched, _ := regexp.MatchString("^/[a-z0-9._]*\\.[a-z]{2,}$", c.Request.RequestURI)
	if !matched {
		// try to parse it as a static file
		static.ServeRoot("/", "root")(c)
		return
	}

	// get the requested domain
	domain := c.Request.RequestURI[1:]

	// look for it in our database

	// if it's not found
	// - find the email address
	// - store the information in the database
	// - send the verificationcode
	// - let the user know who we emailed

	// if it is found
	// - get our last known information
	// - display it to the user
	c.HTML(http.StatusOK, "viewdomain.html", gin.H{
		"domain": domain,
	})
}
