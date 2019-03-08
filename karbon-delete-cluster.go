package main

import (
	"fmt"
	"time"
    "encoding/json"
	"github.com/kumarsarath588/karbon/client"
	//mainv3 "github.com/terraform-providers/terraform-provider-nutanix/client/v3"
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
	resp, err := (*client).API.V3.DeleteKarbonCluster("e926e552-be50-4384-4791-aff6ff0dbd21")
	if err != nil {
		fmt.Println(err)
	}
	status := "RUNNING"
	var out []byte
	for status != "SUCCEEDED" {
		taskresp, err := (*client).API.V3.TaskStatus(*resp.TaskUUID)
		if err != nil {
			fmt.Println(err)
		}

		out, err = json.Marshal(taskresp)
		if err != nil {
			fmt.Println(err)
		}

		status = *taskresp.Status
		time.Sleep(10)
	}
	fmt.Println("Last:", string(out))
}
