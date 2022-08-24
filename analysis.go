/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

import (
	"context"
	"fmt"
)

// AnalysisService handles users for the Tenable instance / API.
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/Analysis.htm
type AnalysisService struct {
	client *Client
}

type Severity struct {
	ID          interface{} `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
}

type VPRContext struct {
	ID    interface{} `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Value string      `json:"value,omitempty"`
	Type  string      `json:"type,omitempty"`
}

type Family struct {
	ID   interface{} `json:"id,omitempty"`
	Name string      `json:"name,omitempty"`
	Type string      `json:"type,omitempty"`
}

type Analysis struct {
	PluginID       string      `json:"pluginID"`
	Severity       Severity    `json:"severity,omitempty"`
	VPRScore       string      `json:"vprScore,omitempty"`
	VPRContext     interface{} `json:"vprContext,omitempty"`
	IP             string      `json:"ip,omitempty"`
	UUID           string      `json:"uuid,omitempty"`
	Port           string      `json:"port,omitempty"`
	Protocol       string      `json:"protocol,omitempty"`
	Name           string      `json:"name,omitempty"`
	DNSName        string      `json:"dnsName,omitempty"`
	MACAddress     string      `json:"macAddress,omitempty"`
	NetBiosName    string      `json:"netBiosName,omitempty"`
	Uniqueness     string      `json:"uniqueness,omitempty"`
	HostUniqueness string      `json:"hostUniqueness,omitempty"`
	Family         Family      `json:"family,omitempty"`
	Repository     Repository  `json:"repository,omitempty"`
	PluginInfo     string      `json:"pluginInfo,omitempty"`
}

type AnalysisResultSet struct {
	TotalRecords             string     `json:"totalRecords,omitempty"`
	ReturnedRecords          int64      `json:"returnedRecords,omitempty"`
	StartOffset              string     `json:"startOffset,omitempty"`
	EndOffset                string     `json:"endOffset,omitempty"`
	MatchingDataElementCount string     `json:"MatchingDataElementCount,omitempty"`
	Results                  []Analysis `json:"results,omitempty"`
}

// Analysis represents a Tenable user.
type AnalysisResponse struct {
	Type      string            `json:",type,omitempty"`
	Response  AnalysisResultSet `json:"response,omitempty"`
	ErrorCode int               `json:"error_code,omitempty"`
	ErrorMsg  string            `json:"error_msg,omitempty"`
	Warnings  []string          `json:"warnings,omitempty"`
	Timestamp int               `json:"timestamp,omitempty"`
}

// PostWithContext gets user info from Tenable using its Account Id
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/Analysis.htm
func (s *AnalysisService) PostWithContext(ctx context.Context, body interface{}) (*AnalysisResponse, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/analysis")
	fmt.Printf("url: %s\n", apiEndpoint)
	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, body)
	if err != nil {
		return nil, nil, err
	}

	repoResp := new(AnalysisResponse)
	resp, err := s.client.Do(req, repoResp)
	if err != nil {
		return nil, resp, NewTenableError(resp, err)
	}
	return repoResp, resp, nil
}

// Get wraps PostWithContext using the background context.
func (s *AnalysisService) Post(body interface{}) (*AnalysisResponse, *Response, error) {
	return s.PostWithContext(context.Background(), body)
}

type AnalysisFilter struct {
	ID           string      `json:"id"`
	FilterName   string      `json:"filterName"`
	Operator     string      `json:"operator"`
	Type         string      `json:"type"`
	IsPredefined bool        `json:"isPredefined"`
	Value        interface{} `json:"value"`
}

type AnalysisQuery struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Context      string   `json:"context"`
	Status       int64    `json:"status"`
	CreatedTime  int64    `json:"createdTime"`
	ModifiedTime int64    `json:"modifiedTime"`
	Groups       []string `json:"groups"`
	Type         string   `json:"type"`
	Tool         string   `json:"tool"`
	SourceType   string   `json:"sourceType"`
	StartOffset  int64    `json:"startOffset"`
	EndOffset    int64    `json:"endOffset"`

	Filters  []AnalysisFilter `json:"filters"`
	VulnTool string           `json:"vulnTool"`
}

type AnalysisBody struct {
	Query      AnalysisQuery `json:"query"`
	SourceType string        `json:"sourceType"`
	Columns    interface{}   `json:"columns"`
	Type       string        `json:"type"`
}
