package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"html/template"

	beegoorm "github.com/astaxie/beego/orm"
	"github.com/dineshappavoo/basex"
)

type CheckNewDomain struct {
	VerificationCode string
}

func (c CheckNewDomain) Run() {
	var msgText bytes.Buffer
	var msgHTML bytes.Buffer

	domain := Domain{OwnerEmail: ""}
	err := orm.Read(&domain, "OwnerEmail")
	// No results, no log
	if err == beegoorm.ErrNoRows {
		return
	}
	if err != nil {
		// TODO email me about this
		fmt.Println("Error while running CheckNewDomain(Read):", err)
		return
	}
	fmt.Println("Adding", domain.Name)
	whois, err := GetWhoisInfo(domain.Name)
	if err != nil {
		// TODO email me about this
		fmt.Println("Error while running CheckNewDomain(whois):", err)
		return
	}
	// Save the url
	domain.OwnerEmail = whois.Emails[0]
	// Generate a unique validation code

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	encoded, _ := basex.Encode(strconv.Itoa(r.Intn(916132832)))
	// TODO email me about this
	domain.VerificationCode = encoded
	_, err = orm.Update(&domain)
	if err != nil {
		fmt.Println("Error while running CheckNewDomain(Update):", err)
		return
	}
	c.VerificationCode = encoded

	TemplateText, err := template.
		New("verificationemail.txt").
		ParseFiles("tmpls/verificationemail.txt")
	if err != nil {
		log.Println("[ERROR] CheckNewDomain TemplateText Parse", err)
		return
	}

	TemplateHTML, err := template.
		New("verificationemail.html").
		ParseFiles("tmpls/verificationemail.html")
	if err != nil {
		log.Println("[ERROR] CheckNewDomain TemplateHTML Parse", err)
		return
	}

	err = TemplateText.Execute(&msgText, c)
	if err != nil {
		log.Println("[ERROR] CheckNewDomain TemplateText Execute", err)
		return
	}
	err = TemplateHTML.Execute(&msgHTML, c)
	if err != nil {
		log.Println("[ERROR] CheckNewDomain TemplateHTML Execute", err)
		return
	}

	err = SendEmail(domain.OwnerEmail,
		"domain.glass Verification Code",
		msgText,
		msgHTML,
	)
	if err != nil {
		log.Println("[ERROR] CheckNewDomain", err)
		return
	}
}
