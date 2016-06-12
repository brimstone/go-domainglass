domainglass
===========
[![Build Status](https://travis-ci.org/brimstone/go-domainglass.svg?branch=master)](https://travis-ci.org/brimstone/go-domainglass)
[![Coverage Status](https://coveralls.io/repos/github/brimstone/go-domainglass/badge.svg?branch=master)](https://coveralls.io/github/brimstone/go-domainglass?branch=master)
[![Stories in Ready](https://badge.waffle.io/brimstone/go-domainglass.png?label=ready&title=Ready)](https://waffle.io/brimstone/go-domainglass)


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

API:
- POST `/` Sets up domain to be checked.
  - domain={{domain}}
- GET `/{{domain}}` Gets status of checks as json
  - `checking`: true or false.
    - true: The system is checking attributes of the domain
    - false: The system is not keeping an eye on the domain. Check on `cooldowntime`
  - `cooldowntime`: Time of when the system will start checking the domain again.
  - `email`: Email address alerts are sent
  - `updated`: Time of last check of any type.
- GET `/{{domain}}/{{verificationHash}}` Shows details about checks as json.
  - Same as the `/{{domain}}` checks above, including:
  - `checks`: Array of checks.
    - `updated`: Time when the check was last performed.
    - `name`: Name of check.
    - `value`: Value of the check the system is watching to change.
- POST `/{{domain}}/{{verificationHash}}/cancel` Terminates service.
  - confirm=true
