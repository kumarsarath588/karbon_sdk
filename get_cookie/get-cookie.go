package main

import (
	"fmt"
	"crypto/tls"
	"net/http"
	"net/url"
	"encoding/base64"
	"github.com/terraform-providers/terraform-provider-nutanix/client"
	karbonclient "github.com/kumarsarath588/karbon/client"
	"bytes"
)

const (
	libraryVersion = "v3"
	defaultBaseURL = "https://%s/"
	absolutePath   = "api/nutanix/" + libraryVersion
	karbonabsolutePath   = "acs/k8s"
	userAgent      = "nutanix/" + libraryVersion
	mediaType      = "application/json"
)

func main() {
	configCreds := &client.Credentials{
		URL:      fmt.Sprintf("%s:%s", "10.46.6.2", "9440"),
		Endpoint: "10.46.6.2",
		Username: "admin",
		Password: "Nutanix.123",
		Port:     "9440",
		Insecure: true,
	}
	c, err := client.NewClient(configCreds)

	urlStr := "/clusters/list"
	rel, errp := url.Parse(absolutePath + urlStr)
	if errp != nil {
		fmt.Println(errp)
	}
	u := c.BaseURL.ResolveReference(rel)
	var body = []byte(`{}`)
	req, err := http.NewRequest(http.MethodPost,u.String(), bytes.NewBuffer(body))
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Authorization", "Basic "+
		base64.StdEncoding.EncodeToString([]byte("admin"+":"+"Nutanix.123")))

	client := http.DefaultClient
	transCfg := &http.Transport{
		// nolint:gas
		TLSClientConfig: &tls.Config{InsecureSkipVerify: configCreds.Insecure}, // ignore expired SSL certificates
	}
	client.Transport = transCfg
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Printf("%+v \n",cookie)
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	fmt.Println("")
//carbon client create
	kconfigCreds := &karbonclient.Credentials{
		URL:      fmt.Sprintf("%s:%s", "10.46.6.2", "7050"),
		Endpoint: "10.46.6.2",
		Username: "admin",
		Password: "Nutanix.123",
		Port:     "7050",
		Insecure: true,
	}

	kc, err := karbonclient.NewClient(kconfigCreds)
	kurlStr := "/cluster/list"

	krel, errp := url.Parse(karbonabsolutePath + kurlStr)
	if errp != nil {
		fmt.Println(errp)
	}
	ku := kc.BaseURL.ResolveReference(krel)
	kreq, err := http.NewRequest(http.MethodPost,ku.String(), bytes.NewBuffer(body))
	kreq.Header.Add("Content-Type", mediaType)
	kreq.Header.Add("Accept", mediaType)
	kreq.Header.Add("User-Agent", userAgent)
	kreq.Header.Add("Authorization", "Basic "+
		base64.StdEncoding.EncodeToString([]byte("admin"+":"+"Nutanix.123")))
	for _, cookie := range cookies {
		kreq.AddCookie(cookie)
	}
	kresp, err := client.Do(kreq)
	if err != nil{
		fmt.Println(err)
	}
	req_cookie, err := kreq.Cookie("NTNX_IGW_SESSION")
	fmt.Printf("%+v \n",req_cookie)
	fmt.Printf("%+v \n",*kreq)
	fmt.Printf("%+v \n",kresp)
}
