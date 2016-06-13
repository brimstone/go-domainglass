package main

import "github.com/bamzi/jobrunner"

// InitJobs Setup scheduled jobs
func InitJobs() error {
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@midnight", AnalyticsEmails{})
	jobrunner.Schedule("@every 5s", CheckNewDomain{})
	return nil
}
