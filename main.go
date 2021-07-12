package main

import (
	"fmt"
	"github.com/vamegh/ssl-reverse-proxy/pkg/aws"
	"github.com/vamegh/ssl-reverse-proxy/pkg/cmdHandler"
	"github.com/vamegh/ssl-reverse-proxy/pkg/configHandler"
	"github.com/vamegh/ssl-reverse-proxy/pkg/proxy"
	"log"
	"sigs.k8s.io/yaml"
)


func main() {
	fmt.Print("BAH HUMBUG 1 ")
	args := cmdHandler.ArgParser()
	data := configHandler.GetData(args.ConfigFile)
	fmt.Print("BAH HUMBUG 2")
	dataTest, _ := yaml.Marshal(data)
	//dataTestJson, _ := yaml.YAMLToJSON(dataTest)
	fmt.Println(string(dataTest))

	type ProxyInterface interface {

	}
	if data.Aws.Enabled == true {
		secretMap, err := utils.GetSecret(data.Aws.Region, data.Aws.SecretsManager.SecretName)
		if err != nil {
			log.Fatalf("Failed to Retrieve Secret, Error: %s", err)
		}

		// Just printing out the secrets for now - not sure what values needed from here yet.
		for key, value := range secretMap {
			fmt.Printf("Key: %s, Value: %s", key, value)
		}
	}
	for idx, ProxyDetails := range data.ProxyDetails {
		log.Printf("Reading through Proxy Config: %d, data: %s", idx, ProxyDetails)
		log.Printf("destination: %s", ProxyDetails.Destination)
		log.Printf("Source Host: %s", ProxyDetails.SrcHost)
		proxy.RevProxy(&ProxyDetails)
	}




}
