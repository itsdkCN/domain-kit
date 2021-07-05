package main

import (
	"flag"
	"fmt"
	"github.com/itsdkCN/domain-kit/daddy"
)

func get(domainData Domain) {
	idcType := domainData.idcType
	keyID := domainData.keyId
	secretKey := domainData.secretKey
	host := domainData.host
	domain := domainData.domain
	name := domainData.name
	flag.Parse()
	if idcType == "aws" {
		fmt.Println(host)
	} else if idcType == "ali" {

	} else if idcType == "godaddy" {
		client, _ := daddy.NewClient(keyID, secretKey, false)
		resp, _ := client.Domains.GetRecords(domain, "A", name, 0, 10)
		fmt.Printf("response is %#v\n", resp)
	} else {
		fmt.Println("不支持的类型")
	}
	return
}