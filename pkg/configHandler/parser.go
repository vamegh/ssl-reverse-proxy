package configHandler

import (
	"io/ioutil"
	"log"
	"sigs.k8s.io/yaml"
)

type ProxyDetails struct {
	Destination  string `json:"destination"`
	SrcHost      string `json:"src_host"`
	SrcPort      int    `json:"src_port"`
	SrcProto     string `json:"src_protocol"`
	StripPrefix  bool   `json:"strip_prefix"`
	IdleTimeout  int    `json:"idle_timeout"`
}

type Management struct {
	Health string `json:"health"`
	Info   string `json:"info"`
}

type Credentials struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type SecretsManager struct {
	SecretName string `json:"secret_name"`
}

type AWS struct {
	Region         string         `json:"region"`
	Enabled        bool           `json:"enabled"`
	SecretsManager SecretsManager `json:"secrets_manager"`
	Credentials    Credentials    `json:"credentials"`
}

type Conf struct {
	ProxyDetails []ProxyDetails `json:"proxy_details"`
	Management   Management     `json:"management"`
	Aws          AWS            `json:"aws"`
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


