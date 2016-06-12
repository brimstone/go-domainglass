package main

import "fmt"

type CheckNewDomain struct {
}

func (c CheckNewDomain) Run() {
	fmt.Println("Checking for New domains")
	domain := Domain{OwnerEmail: ""}
	err := orm.Read(&domain, "OwnerEmail")
	if err != nil {
		fmt.Println("Error while running CheckNewDomain(Read):", err)
		return
	}
	fmt.Println(domain)
	whois, err := GetWhoisInfo(domain.Name)
	if err != nil {
		fmt.Println("Error while running CheckNewDomain(whois):", err)
		return
	}
	domain.OwnerEmail = whois.Emails[0]
	_, err = orm.Update(&domain)
	if err != nil {
		fmt.Println("Error while running CheckNewDomain(Update):", err)
		return
	}
}
