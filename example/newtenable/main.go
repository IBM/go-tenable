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
	sco_url := os.Getenv("SC05_URL")
	if sco_url == "" {
		fmt.Println("environment variable SCO5_URL required")
		return
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{Transport: transport}

	c, err := tenable.NewClient(&httpClient, sco_url)
	if err != nil {
		fmt.Println(err.Error())
	}
	user, resp, err := c.Repository.Get("", "id,name")
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = resp
	for _, v := range user {
		fmt.Println(v.Name, v.ID)
	}

}
