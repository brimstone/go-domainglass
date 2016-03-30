package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ClientRequest TODO
type ClientRequest struct {
	Timestamp    time.Time
	IP           string
	URL          string
	UserAgent    string
	Referer      string
	HTTPCode     int
	ResponseTime time.Duration
}

// Analytics TODO
func Analytics(c *gin.Context) {
	// Pass on to the next-in-chain

	then := time.Now()
	c.Next()

	cr := ClientRequest{
		Timestamp:    time.Now(),
		UserAgent:    c.Request.UserAgent(),
		IP:           c.Request.RemoteAddr,
		URL:          c.Request.RequestURI,
		Referer:      c.Request.Referer(),
		HTTPCode:     c.Writer.Status(),
		ResponseTime: time.Since(then),
	}
}
