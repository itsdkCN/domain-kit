package main

import (
	"flag"
	"github.com/itsdkCN/domain-kit/op"
)

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
	domainData := op.Domain{
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
		op.Add(domainData)
		break
	case "get":
		op.Get(domainData)
		break
	case "query":
		op.Query(domainData)
		break
	case "delete":
		op.Delete(domainData)
	}
}
