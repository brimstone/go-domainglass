package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ClientRequest TODO
type ClientRequest struct {
	Timestamp    time.Time `xorm:"pk not null 'timestamp'"`
	IP           string    `xorm:"pk not null 'ip'"`
	URL          string    `xorm:"not null 'url'"`
	UserAgent    string    `xorm:"'user-agent'"`
	Referer      string    `xorm:"'referer'"`
	HTTPCode     int       `xorm:"'httpcode'"`
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
