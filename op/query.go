package op

import (
	"flag"
	"fmt"
	"github.com/itsdkCN/domain-kit/daddy"
)

func Query(domainData Domain) {
	idcType := domainData.IdcType
	keyID := domainData.KeyId
	secretKey := domainData.SecretKey
	host := domainData.Host
	flag.Parse()
	if idcType == "aws" {
		fmt.Println(host)
	} else if idcType == "ali" {

	} else if idcType == "godaddy" {
		client, _ := daddy.NewClient(keyID, secretKey, false)
		resp, err := client.Domains.List([]string{""}, []string{""}, 100, "", []string{}, "")
		if err != nil {
			fmt.Printf("response is %#v\n", err)
			return
		}
		fmt.Printf("response is %#v\n", resp)
	} else {
		fmt.Println("不支持的类型")
	}
	return
}
