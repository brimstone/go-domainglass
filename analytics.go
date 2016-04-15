package main

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ClientRequest TODO
type ClientRequest struct {
	Timestamp    time.Time `orm:"pk;column(timestamp);type(datetime)"`
	IP           string    `orm:"pk;column(ip)"`
	URL          string    `orm:"column(url)"`
	UserAgent    string    `orm:"column(user-agent)"`
	Referer      string
	HTTPCode     int           `orm:"column(httpcode)"`
	ResponseTime time.Duration `orm:"column(reponse-time)"`
}

/*
// TableUnique returns columns that should be unique
func (c *ClientRequest) TableUnique() [][]string {
	return [][]string{
		[]string{"ip"},
	}
}
*/

// Analytics TODO
func Analytics(c *gin.Context) {
	// Get the current time
	then := time.Now()

	// Let the rest of the request happen
	c.Next()

	clientip := c.Request.Header.Get("X-Forwarded-For")
	if clientip == "" {
		clientip = c.Request.RemoteAddr
	}
	clientipparts := strings.Split(clientip, ":")
	clientip = clientipparts[0]

	// Build our client request record
	cr := ClientRequest{
		Timestamp:    time.Now(),
		UserAgent:    c.Request.UserAgent(),
		IP:           clientip,
		URL:          c.Request.RequestURI,
		Referer:      c.Request.Referer(),
		HTTPCode:     c.Writer.Status(),
		ResponseTime: time.Since(then),
	}
	// log this to the database
	orm.Insert(cr)
}
