package main

import (
	"flag"
)

type Domain struct {
	idcType   string
	keyId     string
	secretKey string
	host      string
	domain    string
	ip        string
	name      string
	opType    string
}

func main() {
	idcType := flag.String("idcType", "", "")
	keyID := flag.String("keyId", "", "")
	secretKey := flag.String("secretKey", "", "")
	host := flag.String("hostId", "", "")
	domain := flag.String("domain", "", "")
	ip := flag.String("ip", "", "")
	name := flag.String("name", *ip, "")
	opType := flag.String("type", "", "")
	flag.Parse()
	domainData := Domain{
		idcType:   *idcType,
		keyId:     *keyID,
		secretKey: *secretKey,
		host:      *host,
		domain:    *domain,
		ip:        *ip,
		name:      *name,
		opType:    *opType,
	}

	switch *opType {
	case "add":
		Add(domainData)
		break
	case "get":
		Get(domainData)
		break
	}
}
