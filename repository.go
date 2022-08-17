/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

import (
	"context"
	"fmt"
)

// RepositoryService handles users for the Tenable instance / API.
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/Repository.htm
type RepositoryService struct {
	client *Client
}

type Repository struct {
	ID          interface{} `json:"id"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	DataFormat  string      `json:"dataFormat,omitempty"`
	UUID        string      `json:"uuid,omitempty"`
}

// Repository represents a Tenable user.
type RepositoryResponse struct {
	// "type": "regular",
	Type string
	// "response": {
	Response []Repository `json:"response"`
	// "error_code": 0,
	ErrorCode int `json:"error_code"`
	// "error_msg": "",
	ErrorMsg string `json:"error_msg"`
	// "warnings": [],
	Warnings []string
	// "timestamp": 1657818772
	Timestamp int
}

// GetWithContext gets user info from Tenable using its Account Id
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/Repository.htm
func (s *RepositoryService) GetWithContext(ctx context.Context, requestType, fields string) ([]Repository, *Response, error) {
	if requestType == "" {
		requestType = "All"
	}
	apiEndpoint := fmt.Sprintf("/rest/repository?type=%s", requestType)
	if len(fields) > 0 {
		apiEndpoint = apiEndpoint + "&fields=" + fields
	}
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	repoResp := new(RepositoryResponse)
	resp, err := s.client.Do(req, repoResp)
	if err != nil {
		return nil, resp, NewTenableError(resp, err)
	}
	return repoResp.Response, resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *RepositoryService) Get(requestType, fields string) ([]Repository, *Response, error) {
	return s.GetWithContext(context.Background(), requestType, fields)
}
