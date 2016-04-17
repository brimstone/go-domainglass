package main

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ClientRequest TODO
type ClientRequest struct {
	ID           int64     `orm:"pk;auto;column(id)"`
	Timestamp    time.Time `orm:"auto_now_add;column(timestamp);type(datetime)"`
	IP           string    `orm:"column(ip)"`
	URL          string    `orm:"column(url)"`
	UserAgent    string
	Referer      string
	HTTPCode     int `orm:"column(httpcode)"`
	ResponseTime time.Duration
}

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
	cr := new(ClientRequest)
	cr.UserAgent = c.Request.UserAgent()
	cr.IP = clientip
	cr.URL = c.Request.RequestURI
	cr.Referer = c.Request.Referer()
	cr.HTTPCode = c.Writer.Status()
	cr.ResponseTime = time.Since(then)
	// log this to the database
	orm.Insert(cr)
}
