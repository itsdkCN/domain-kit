package op

import (
	"context"
	"flag"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/itsdkCN/domain-kit/daddy"
	"log"
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
	} else if idcType == "cloudflare" {
		// Construct a new API object using a global API key
		api, err := cloudflare.NewWithAPIToken(secretKey)
		// alternatively, you can use a scoped API token
		// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}

		// Most API calls require a Context
		ctx := context.Background()

		// Fetch user details on the account
		u, err := api.UserDetails(ctx)
		if err != nil {
			log.Fatal(err)
		}
		// Print user details
		fmt.Println(u)

		// Fetch the zone ID
		id, err := api.ZoneIDByName(domain) // Assuming example.com exists in your Cloudflare account already
		if err != nil {
			log.Fatal(err)
		}

		// Fetch zone details
		zone, err := api.ZoneDetails(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		// Print zone details
		fmt.Println(zone)
	} else {
		fmt.Println("不支持的类型")
	}
	return
}
