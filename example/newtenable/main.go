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
	//"14272,11219,22964"
	/*
		{
			"query":
			{
				"name":"",
				"description":"",
				"context":"",
				"status":-1,
				"createdTime":0,
				"modifiedTime":0,
				"groups":[],
				"type":"vuln",
				"tool":"listvuln",
				"sourceType":"cumulative",
				"startOffset":0,
				"endOffset":50,
				"filters":[
				{
					"id":"pluginID",
					"filterName":"pluginID",
					"operator":"=",
					"type":"vuln",
					"isPredefined":true,
					"value":"14272,11219,22964"
				}
				],
				"vulnTool":"listvuln"
			},
			"sourceType":"cumulative",
			"columns":[],
			"type":"vuln"
		}
	*/
	f := tenable.AnalysisFilter{}
	q := tenable.AnalysisQuery{}
	b := tenable.AnalysisBody{}
	q.Name = ""
	q.Description = ""
	q.Context = ""
	q.Status = -1
	q.CreatedTime = 0
	q.ModifiedTime = 0
	q.Groups = nil
	q.Type = "vuln"
	q.Tool = "listvuln"
	q.SourceType = "cumulative"
	q.StartOffset = 0
	q.EndOffset = 2

	f.ID = "pluginID"
	f.FilterName = "pluginID"
	f.Operator = "="
	f.Type = "vuln"
	f.IsPredefined = true
	f.Value = "14272,11219,22964"
	q.Filters = []tenable.AnalysisFilter{f}
	q.VulnTool = "listvuln"

	b.Query = q
	b.SourceType = "cumulative"
	b.Columns = nil
	b.Type = "vuln"

	items, resp, err := c.Analysis.Post(b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = resp
	for _, v := range items.Response.Results {
		fmt.Println(v)
	}

}
