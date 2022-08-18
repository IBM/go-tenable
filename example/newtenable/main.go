/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	tenable "github.com/IBM/go-tenable"
)

func main() {
	apiURL := os.Getenv("SC05_URL")
	apiKey := os.Getenv("SC05_ACCESS_KEY")
	apiSecret := os.Getenv("SC05_SECRET_KEY")
	if apiURL == "" || apiKey == "" || apiSecret == "" {
		fmt.Printf("Missing env vars. Requred env vars SC05_URL,SC05_ACCESS_KEY and SC05_SECRET_KEY\n")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	t := tenable.APIKeyAuthTransport{Transport: transport,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	c, err := tenable.NewClient(t.Client(), apiURL)
	if err != nil {
		fmt.Println(err.Error())
	}
	user, resp, err := c.Repository.Get("", "id,name")
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = resp
	for _, v := range user {
		fmt.Println(v.ID, v.Name)
	}

}
