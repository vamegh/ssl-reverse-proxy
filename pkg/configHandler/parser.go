package configHandler

import (
	"io/ioutil"
	"log"
	"sigs.k8s.io/yaml"
)

type Conf struct {
	ProxyDetails []struct {
		Destination  string `json:"destination"`
		SrcHost      string `json:"src_host"`
		SrcPort      string `json:"src_port"`
		SrcProto     string `json:"src_protocol"`
		StripPrefix  string `json:"strip_prefix"`
	} `json:"proxy_details"`
	Management struct {
		Health string `json:"health"`
		Info   string `json:"info"`
	} `json:"management"`
	Aws struct {
		Region         string `json:"region"`
		Enabled        bool   `json:"enabled"`
		SecretsManager struct {
			SecretName string `json:"secret_name"`
		} `json:"secrets_manager"`
		Credentials struct {
			AccessKey string `json:"access_key"`
			SecretKey string `json:"secret_key"`
		} `json:"credentials"`
	} `json:"aws"`
}

func GetData(configFile string) Conf {
	var dataStruct Conf
	confData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error Reading: %s :: Message: %s", configFile, err)
	}
	err = yaml.Unmarshal([]byte(confData), &dataStruct)
	if err != nil {
		log.Fatalf("Failed to Parse File: %s, Error: %s", configFile, err)
	}
	return dataStruct
}


