package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"github.com/bamzi/jobrunner"
)

// InitJobs Setup scheduled jobs
func InitJobs() error {
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@midnight", AnalyticsEmails{})
	return nil
}

// AnalyticURL Counts of urls
type AnalyticURL struct {
	Count    int
	URL      string
	HTTPCode int
}

// AnalyticIP Counts of IPs
type AnalyticIP struct {
	Count int
	IP    string
}

// AnalyticReferer Counts of Referers
type AnalyticReferer struct {
	Count   int
	Referer string
}

// AnalyticsEmails Job Specific Functions
type AnalyticsEmails struct {
	// filtered
	WeeklyNot200  []AnalyticURL
	WeeklyIP      []AnalyticIP
	WeeklyReferer []AnalyticReferer
}

// Run will get triggered automatically.
func (e AnalyticsEmails) Run() {
	var err error
	var msgText bytes.Buffer
	var msgHTML bytes.Buffer
	// Queries the DB
	// Sends some email
	fmt.Println("Nightly Analytic email")

	// Non 200s
	results, err := orm.Query(`select count(url) as count, url, httpcode
from client_request
where httpcode != 200
group by url
order by count desc;`)
	for _, r := range results {
		count, _ := strconv.Atoi(string(r["count"]))
		httpcode, _ := strconv.Atoi(string(r["httpcode"]))
		stat := &AnalyticURL{Count: count,
			URL:      string(r["url"]),
			HTTPCode: httpcode,
		}
		e.WeeklyNot200 = append(e.WeeklyNot200, *stat)
	}

	// IPs
	results, err = orm.Query(`select count(ip) as count, ip
from client_request
group by ip
order by count desc`)
	for _, r := range results {
		count, _ := strconv.Atoi(string(r["count"]))
		stat := &AnalyticIP{Count: count,
			IP: string(r["ip"]),
		}
		e.WeeklyIP = append(e.WeeklyIP, *stat)
	}

	// referers
	results, err = orm.Query(`select count(referer) as count, referer
from client_request
group by referer
order by count desc`)
	for _, r := range results {
		count, _ := strconv.Atoi(string(r["count"]))
		stat := &AnalyticReferer{Count: count,
			Referer: string(r["referer"]),
		}
		e.WeeklyReferer = append(e.WeeklyReferer, *stat)
	}

	if err != nil {
		log.Println("[ERROR] EmailAnalytics Query", err)
		return
	}

	TemplateText, err := template.
		New("emailanalytics.txt").
		ParseFiles("tmpls/emailanalytics.txt")
	if err != nil {
		log.Println("[ERROR] EmailAnalytics TemplateText Parse", err)
		return
	}

	TemplateHTML, err := template.
		New("emailanalytics.html").
		ParseFiles("tmpls/emailanalytics.html")
	if err != nil {
		log.Println("[ERROR] EmailAnalytics TemplateHTML Parse", err)
		return
	}

	err = TemplateText.Execute(&msgText, e)
	if err != nil {
		log.Println("[ERROR] EmailAnalytics TemplateText Execute", err)
		return
	}
	err = TemplateHTML.Execute(&msgHTML, e)
	if err != nil {
		log.Println("[ERROR] EmailAnalytics TemplateHTML Execute", err)
		return
	}

	/*
		Debug bits
				log.Println(msgHTML.String())
					return
	*/

	err = SendEmail("matt@domain.glass",
		"Nightly Analytics",
		msgText,
		msgHTML,
	)
	if err != nil {
		log.Println("[ERROR] EmailAnalytics", err)
		return
	}
}
