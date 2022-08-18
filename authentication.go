/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

const (
	// HTTP Basic Authentication
	authTypeAPIKey = 1
)

// AuthenticationService handles users for the Tenable instance / API.
//
// Tenable API docs: https://docs.tenable.com/tenablesc/api/index.htm
type AuthenticationService struct {
	client *Client

	// Authentication type
	authType int

	// apikey auth
	apiKey string

	apiSecret string
}

// SetAPIKeyAuth sets api_key and api_secret for the APIKey auth against the Jira instance.
//
// Deprecated: Use APIKeyAuthTransport instead
func (s *AuthenticationService) SetAPIKeyAuth(api_key, api_secret string) {
	s.apiKey = api_key
	s.apiSecret = api_secret
	s.authType = authTypeAPIKey
}

// Authenticated reports if the current Client has authentication details for Jira
func (s *AuthenticationService) Authenticated() bool {
	if s != nil {

		if s.authType == authTypeAPIKey {
			return (s.apiKey != "" && s.apiSecret != "")
		}
		return false

	}
	return false
}
