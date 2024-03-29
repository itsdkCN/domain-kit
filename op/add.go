package op

import (
	"context"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/cloudflare/cloudflare-go"
	"github.com/itsdkCN/domain-kit/daddy"
	"log"
)

func Add(domainObj Domain) {
	idcType := domainObj.IdcType
	keyID := domainObj.KeyId
	secretKey := domainObj.SecretKey
	host := domainObj.Host
	domain := domainObj.Domain
	ip := domainObj.Ip
	name := domainObj.Name
	flag.Parse()
	if idcType == "aws" {
		awsKeyID := keyID
		awsSecretKey := secretKey
		creds := credentials.NewStaticCredentials(awsKeyID, awsSecretKey, "")
		awsConfig := aws.Config{Region: aws.String("us-west-2"), Credentials: creds}
		newSession, err := session.NewSession(&awsConfig)
		if err != nil {
			return
		}
		r53Svc := route53.New(newSession)
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
	} else if idcType == "cloudflare" {
		api, err := cloudflare.New(secretKey, keyID)
		// alternatively, you can use a scoped API token
		// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}

		// Most API calls require a Context
		ctx := context.Background()
		zoneId, err := api.ZoneIDByName(domain)
		if err != nil {
			return
		}
		_, err = api.CreateDNSRecord(ctx, zoneId, cloudflare.DNSRecord{Content: ip, Name: name, TTL: 3600, Type: "A"})
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		fmt.Printf("response is %#v\n", name+"."+domain)

	} else {
		fmt.Println("不支持的类型")
	}
}
