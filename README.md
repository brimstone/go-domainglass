domainglass
===========
[![Build Status](https://travis-ci.org/brimstone/go-domainglass.svg?branch=master)](https://travis-ci.org/brimstone/go-domainglass)
[![Coverage Status](https://coveralls.io/repos/github/brimstone/go-domainglass/badge.svg?branch=master)](https://coveralls.io/github/brimstone/go-domainglass?branch=master)


Use case:
- Take domainname as user input
- Save domain and verification key to database
- Send admin email address verification email
  -  use postmaster@ if no admin or technical contact
- When user clicks link in verification email, validate domain

Worker:
- Periodically run each check on each valid domain in database

UI:
- `/` Shows page prompting for domain
- `/about` Explains more about how this service works
- `/{{domain}}` Shows report for last checks against domain or status that verification is pending
- `/{{domain}}/{{verificationHash}}` Shows details about checks, lets user cancel service
- `/{{domain}}/{{verificationHash}}/cancel` Terminates service

API:
- `/{{domain}}.json` Gets status of checks as json
- `/{{domain}}/{{verificationHash}}.json` Shows details about checks as json

TODO:
- [X] Setup Travis job to vet, lint, test, and push
- [ ] Find free transactional email service
- [ ] Find free domain email hosting service
- [X] Database with openshift
- [ ] Database with Go
  - xorm?
  - mysql if the env var is present, OPENSHIFT_MYSQL_DB_URL
  - Sqlite from memory if not
- [X] Static content with Go
- [ ] Simple UI
- [ ] Task runner/job queue
- [ ] Whois extract
- [ ] DNS record extract
- [ ] TCP service discovery
- [ ] Figure out how to take people's money

Checks:
- [ ] SSL extract
- [ ] HTTP quality check
- [ ] SMTP quality check
- [ ] IMAP quality check
