package main

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Domain holds information about a domain
type Domain struct {
	ID               int64  `xorm:"pk not null autoincr"`
	Name             string `xorm:"unique not null"`
	VerificationCode string `xorm:"not null"`
	OwnerEmail       string `xorm:"not null"`
	Payment          Payment
}

// Payment holds payment information
type Payment struct {
	Timestamp time.Time `xorm:"pk not null 'timestamp'"`
	Plan      string    `xorm:"not null"`
}

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
