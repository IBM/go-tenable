/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

import (
	"context"
	"fmt"
)

// CurrentUserService handles users for the Tenable instance / API.
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/CurrentUser.htm
type CurrentUserService struct {
	client *Client
}

// CurrentUserResponse represents a Tenable user response
type CurrentUser struct {
	User
}

// CurrentUser represents a Tenable user.
type CurrentUserResponse struct {
	// "type": "regular",
	Type string
	// "response": {
	Response CurrentUser `json:"response"`
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
// Tenable API docs: https://docs.tenable.com/tenablesc/api/CurrentUser.md
func (s *CurrentUserService) GetWithContext(ctx context.Context) (*CurrentUser, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/currentUser")
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(CurrentUserResponse)
	//fmt.Println(req.URL)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, NewTenableError(resp, err)
	}
	return &user.Response, resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *CurrentUserService) Get() (*CurrentUser, *Response, error) {
	return s.GetWithContext(context.Background())
}
