package main

import (
	"io/ioutil"
	"net"
	"sort"
	"strings"
	"time"
)

type WhoisServer struct {
	ID      int64  `orm:"pk;auto;column(id)"`
	TLD     string `orm:"column(tld)"`
	Server  string
	Updated time.Time `orm:"auto_now_add"`
}

type WhoisInfo struct {
	ID          int64 `orm:"pk;auto;column(id)"`
	DomainName  string
	Raw         string
	Emails      []string
	Nameservers []string
}

func GetWhoisRaw(server string, domain string) (string, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	res, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func GetWhoisMap(whoisdata string) (map[string][]string, error) {
	out := make(map[string][]string)
	lines := strings.Split(whoisdata, "\n")
	for _, line := range lines {
		// ignore empty lines
		if len(line) == 0 {
			continue
		}
		// ignore headers
		if line[0] == '%' {
			continue
		}
		kv := strings.SplitN(line, ":", 2)
		// ignore lines that don't split in exactly two
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		out[key] = append(out[key], value)
	}
	return out, nil
}

func GetWhoisInfo(domain string) (WhoisInfo, error) {

	domainSlice := strings.Split(domain, ".")
	tld := domainSlice[len(domainSlice)-1]
	sld := domainSlice[len(domainSlice)-2]

	whoisserver := WhoisServer{TLD: tld}
	// I don't think I care about created and id
	_, _, err := orm.ReadOrCreate(&whoisserver, "TLD")
	// If the whois server for the domain doesn't exist in the database
	if err != nil {
		return WhoisInfo{}, err
	}
	if whoisserver.Server == "" {
		// ask whois.iana.org about it

		ianaRaw, err := GetWhoisRaw("whois.iana.org", tld)
		if err != nil {
			return WhoisInfo{}, err
		}
		// parse the result for the whois line
		ianaMap, err := GetWhoisMap(ianaRaw)
		if err != nil {
			return WhoisInfo{}, err
		}
		// save the server in the database
		whoisserver.Server = ianaMap["whois"][0]
		whoisserver.Updated = time.Now()
		orm.Update(&whoisserver)

	}
	// TODO check for empty whois
	domainRaw, err := GetWhoisRaw(whoisserver.Server, sld+"."+tld)
	if err != nil {
		return WhoisInfo{}, err
	}
	domainMap, err := GetWhoisMap(domainRaw)
	if err != nil {
		return WhoisInfo{}, err
	}
	whois := WhoisInfo{DomainName: domain}
	whois.Raw = domainRaw
	// let a map do the deduping for us
	emails := make(map[string]bool)
	for _, v := range domainMap {
		if strings.Index(v[0], "@") != -1 {
			emails[v[0]] = true
		}
	}
	// only return unique emails
	for k := range emails {
		whois.Emails = append(whois.Emails, k)
	}
	// get nameservers
	whois.Nameservers = domainMap["Name Server"]
	sort.Strings(whois.Nameservers)
	return whois, nil
}
