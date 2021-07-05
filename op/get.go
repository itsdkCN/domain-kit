package op

import (
	"flag"
	"fmt"
	"github.com/itsdkCN/domain-kit/daddy"
)

func Get(domainData Domain) {
	idcType := domainData.IdcType
	keyID := domainData.KeyId
	secretKey := domainData.SecretKey
	host := domainData.Host
	domain := domainData.Domain
	name := domainData.Name
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
