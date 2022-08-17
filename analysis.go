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
	PluginID       string     `json:"pluginID"`
	Severity       Severity   `json:"severity,omitempty"`
	VPRScore       string     `json:"vprScore,omitempty"`
	VPRContext     VPRContext `json:"vprContext,omitempty"`
	IP             string     `json:"ip,omitempty"`
	UUID           string     `json:"uuid,omitempty"`
	Port           string     `json:"port,omitempty"`
	Protocol       string     `json:"protocol,omitempty"`
	Name           string     `json:"name,omitempty"`
	DNSName        string     `json:"dnsName,omitempty"`
	MACAddress     string     `json:"macAddress,omitempty"`
	NetBiosName    string     `json:"netBiosName,omitempty"`
	Uniqueness     string     `json:"uniqueness,omitempty"`
	HostUniqueness string     `json:"hostUniqueness,omitempty"`
	Family         Family     `json:"family,omitempty"`
	Repository     Repository `json:"repository,omitempty"`
	PluginInfo     string     `json:"pluginInfo,omitempty"`
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
	Type      string              `json:",type,omitempty"`
	Response  []AnalysisResultSet `json:"response"`
	ErrorCode int                 `json:"error_code"`
	ErrorMsg  string              `json:"error_msg"`
	Warnings  []string            `json:"warnings,omitempty"`
	Timestamp int                 `json:"timestamp,omitempty"`
}

// GetWithContext gets user info from Tenable using its Account Id
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/Analysis.htm
func (s *AnalysisService) GetWithContext(ctx context.Context, requestType, fields string) (*AnalysisResponse, *Response, error) {
	if requestType == "" {
		requestType = "All"
	}
	apiEndpoint := fmt.Sprintf("/rest/analysis?type=%s", requestType)
	if len(fields) > 0 {
		apiEndpoint = apiEndpoint + "&fields=" + fields
	}
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
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

// Get wraps GetWithContext using the background context.
func (s *AnalysisService) Get(requestType, fields string) (*AnalysisResponse, *Response, error) {
	return s.GetWithContext(context.Background(), requestType, fields)
}
