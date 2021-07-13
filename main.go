package main

import (
	"fmt"
	"github.com/vamegh/ssl-reverse-proxy/pkg/aws"
	"github.com/vamegh/ssl-reverse-proxy/pkg/cmdHandler"
	"github.com/vamegh/ssl-reverse-proxy/pkg/configHandler"
	"github.com/vamegh/ssl-reverse-proxy/pkg/proxy"
	"log"
	"net/http"
	"sigs.k8s.io/yaml"
)

func main() {
	args := cmdHandler.ArgParser()
	data := configHandler.GetData(args.ConfigFile)
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
		revProxy := proxy.RevProxy(ProxyDetails)

		http.HandleFunc(ProxyDetails.Destination, func(w http.ResponseWriter, r *http.Request) {
			revProxy.ServeHTTP(w, r)
		})
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
