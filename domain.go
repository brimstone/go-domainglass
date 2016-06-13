package main

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Domain holds information about a domain
type Domain struct {
	ID               int64 `orm:"pk;auto;column(id)"`
	Name             string
	VerificationCode string `orm:"unique"`
	OwnerEmail       string
	Payments         []*Payment `orm:"reverse(many)"`
}

// Payment holds payment information
type Payment struct {
	ID        int64 `orm:"pk;auto;column(id)"`
	Timestamp time.Time
	Plan      string
	Domain    *Domain `orm:"rel(fk)"`
}

func validDomain(c *gin.Context, domainName string) bool {

	matched, _ := regexp.MatchString("^[a-z0-9._]*\\.[a-z]{2,}$", domainName)
	if !matched {
		c.JSON(400, gin.H{
			"error": "Domain " + domainName + " not in a valid format",
		})
		return false
	}
	return true
}

// Domain Handle main route
func getDomain(c *gin.Context) {
	domainName := strings.ToLower(c.Param("domain"))
	if !validDomain(c, domainName) {
		return
	}
	// This should only read from the database
	domain := Domain{Name: domainName}
	err := orm.Read(&domain, "Name")
	// if the domain doesn't exist
	if err != nil {
		// Let the user know
		c.JSON(400, gin.H{
			"error": "Domain " + domainName + " does not exist",
		})
		return
	}
	/*
		whois, _ := GetWhoisInfo(domain)
	*/
	c.JSON(200, gin.H{
		"checking":     true,       // TODO make this mean something
		"cooldowntime": time.Now(), // TODO account for payments, checks, etc
		"email":        domain.OwnerEmail,
		"updated":      time.Now(), // TODO make this a max of the checks
	})
}

// addDomain
func addDomain(c *gin.Context) {
	domainName := strings.ToLower(c.PostForm("domain"))
	if !validDomain(c, domainName) {
		return
	}
	domain := Domain{Name: domainName}
	_, _, err := orm.ReadOrCreate(&domain, "Name")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Queued for checking",
	})

}

func viewdomain(c *gin.Context) {
	// if we don't have a suburl that looks like a domain
	matched, _ := regexp.MatchString("^/[a-z0-9._]*\\.[a-z]{2,}$", c.Request.RequestURI)
	if !matched {
		// try to parse it as a static file
		static.ServeRoot("/", "root")(c)
		return
	}

	// if we're not in the beta, kick them to the front page
	if beta, _ := c.Cookie("beta"); beta != "true" {
		c.Redirect(302, "/")
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
