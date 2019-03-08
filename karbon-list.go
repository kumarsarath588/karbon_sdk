package main

import (
	"fmt"
    "encoding/json"
	"github.com/kumarsarath588/karbon/client"
	mainv3 "github.com/terraform-providers/terraform-provider-nutanix/client/v3"
	"github.com/kumarsarath588/karbon/client/v3"
)

//Config
type Config struct {
	Endpoint string
	Username string
	Password string
	Port     string
	Insecure bool
}

// Client represents the nutanix API client
type Client struct {
	API *v3.Client
}

// Client ...
func (c *Config) Client() (*Client, error) {

	configCreds := client.Credentials{
		URL:      fmt.Sprintf("%s:%s", c.Endpoint, c.Port),
		Endpoint: c.Endpoint,
		Username: c.Username,
		Password: c.Password,
		Port:     c.Port,
		Insecure: c.Insecure,
	}
	v3, err := v3.NewV3Client(configCreds)
	if err != nil {
		return nil, err
	}
	client := &Client{
		API: v3,
	}

	return client, nil
}

func main() {

	config := &Config{
		Endpoint: "10.46.4.2",
		Username: "admin",
		Password: "Nutanix.123",
		Port:	"7050",
		Insecure: true,
	}

	client, err := config.Client()
	if err != nil {
		fmt.Println(err)
    }
	meta := &mainv3.DSMetadata{}
	resp, err := (*client).API.V3.ListKarbonClusters(meta)
	if err != nil {
		fmt.Println(err)
	}
	out, err := json.Marshal(resp)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Last:", string(out))
}
