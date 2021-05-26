package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Clients struct {
	session *session.Session
	configs map[string]*aws.Config
}

func (c Clients) Session() *session.Session {
	if c.session != nil {
		return c.session
	}
	sess := session.Must(session.NewSession())
	c.session = sess
	return sess
}

func (c Clients) Config(roleArn *string, r string) *aws.Config {
	// return no config for nil inputs
	if roleArn == nil || *roleArn == "" {
		return nil
	}

	region := r
	// include region in cache key otherwise concurrency errors
	key := fmt.Sprintf("%v::%v", region, roleArn)
	// check for cached config
	if c.configs != nil && c.configs[key] != nil {
		return c.configs[key]
	}
	// new creds
	creds := stscreds.NewCredentials(c.Session(), *roleArn)
	// new config
	config := aws.NewConfig().WithCredentials(creds).WithRegion(region).WithMaxRetries(10)
	if c.configs == nil {
		c.configs = map[string]*aws.Config{}
	}
	c.configs[key] = config
	return config
}
