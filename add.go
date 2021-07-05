package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/itsdkCN/domain-kit/daddy"
)

func add(domainObj Domain) {
	idcType := domainObj.idcType
	keyID := domainObj.keyId
	secretKey := domainObj.secretKey
	host := domainObj.host
	domain := domainObj.domain
	ip := domainObj.ip
	name := domainObj.name
	flag.Parse()
	if idcType == "aws" {
		awsKeyID := keyID
		awsSecretKey := secretKey
		creds := credentials.NewStaticCredentials(awsKeyID, awsSecretKey, "")
		awsConfig := aws.Config{Region: aws.String("us-west-2"), Credentials: creds}
		r53Svc := route53.New(session.New(&awsConfig))
		fmt.Printf("domain name：%s\ninput ip：%s\n", domain, ip)
		input := &route53.ChangeResourceRecordSetsInput{
			HostedZoneId: aws.String(host),
			ChangeBatch: &route53.ChangeBatch{
				Comment: aws.String("Add Domain Success"),
				Changes: []*route53.Change{
					{
						Action: aws.String(route53.ChangeActionUpsert),
						ResourceRecordSet: &route53.ResourceRecordSet{
							Name: aws.String(name + "." + domain),
							Type: aws.String(route53.RRTypeA),
							TTL:  aws.Int64(300),
							ResourceRecords: []*route53.ResourceRecord{
								{
									Value: aws.String(ip),
								},
							},
						},
					},
				},
			},
		}

		result, err := r53Svc.ChangeResourceRecordSets(input)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	} else if idcType == "ali" {
		client, err := alidns.NewClientWithAccessKey(host, keyID, secretKey)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		request := alidns.CreateAddDomainRecordRequest()
		request.Scheme = "https"
		fmt.Printf("input：%s\n", ip)
		request.DomainName = domain
		request.RR = name
		request.Type = "A"
		request.Value = ip
		response, err := client.AddDomainRecord(request)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Printf("response is %#v\n", response)
	} else if idcType == "godaddy" {
		client, _ := daddy.NewClient(keyID, secretKey, false)
		records := []daddy.DNSRecord{
			{Data: ip, Name: name, TTL: 3600, Type: "A"},
		}
		err := client.Domains.AddRecords(domain, records)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Printf("response is %#v\n", name+"."+domain)
	} else {
		fmt.Println("不支持的类型")
	}
}
