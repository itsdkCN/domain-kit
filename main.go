package main

import (
	"flag"
)

type Domain struct {
	IdcType   string
	KeyId     string
	SecretKey string
	Host      string
	Domain    string
	Ip        string
	Name      string
	OpType    string
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
		IdcType:   *idcType,
		KeyId:     *keyID,
		SecretKey: *secretKey,
		Host:      *host,
		Domain:    *domain,
		Ip:        *ip,
		Name:      *name,
		OpType:    *opType,
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
