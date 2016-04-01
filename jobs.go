package main

import (
	"fmt"
	"log"

	"github.com/bamzi/jobrunner"
)

// InitJobs Setup scheduled jobs
func InitJobs() error {
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@midnight", AnalyticsEmails{})
	return nil
}

// AnalyticsEmails Job Specific Functions
type AnalyticsEmails struct {
	// filtered
}

// Run will get triggered automatically.
func (e AnalyticsEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Println("Nightly reminder emails")
	err := EmailAnalytics(Domain{"domain.glass", "matt@domain.glass"})
	if err != nil {
		log.Println("[ERROR] EmailAnalytics", err)
	}
}
