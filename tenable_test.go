/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

const (
	testTenableInstanceURL = "https://issues.apache.org/tenable/"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testClient is the Tenable client being tested.
	testClient *Client

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

// setup sets up a test HTTP server along with a Tenable.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	// Tenable client configured to use test server
	var err error
	//fmt.Println(testServer.URL)
	testClient, err = NewClient(nil, testServer.URL)
	if err != nil {
		fmt.Println("new client fail", err)
	}
}

// teardown closes the test HTTP server.
func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testRequestURL(t *testing.T, r *http.Request, want string) {
	if got := r.URL.String(); !strings.HasPrefix(got, want) {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}

func testRequestParams(t *testing.T, r *http.Request, want map[string]string) {
	params := r.URL.Query()

	if len(params) != len(want) {
		t.Errorf("Request params: %d, want %d", len(params), len(want))
	}

	for key, val := range want {
		if got := params.Get(key); val != got {
			t.Errorf("Request params: %s, want %s", got, val)
		}

	}

}
func testSetEnv(t *testing.T) {
	os.Setenv("SC05_ACCESS_KEY", "foo")
	os.Setenv("SC05_SECRET_KEY", "bar")
}
func testTeardownEnv(t *testing.T) {
	os.Unsetenv("SC05_ACCESS_KEY")
	os.Unsetenv("SC05_SECRET_KEY")
}

func TestNewClient_WrongUrl(t *testing.T) {
	testSetEnv(t)
	c, err := NewClient(nil, "://issues.apache.org/Tenable/")

	if err == nil {
		t.Error("Expected an error. Got none")
	}
	if c != nil {
		t.Errorf("Expected no client. Got %+v", c)
	}
}

func TestNewClient_WithHttpClient(t *testing.T) {
	testSetEnv(t)
	httpClient := http.DefaultClient
	httpClient.Timeout = 10 * time.Minute

	c, err := NewClient(httpClient, testTenableInstanceURL)
	if err != nil {
		t.Errorf("Got an error: %s", err)
	}
	if c == nil {
		t.Error("Expected a client. Got none")
		return
	}
	if !reflect.DeepEqual(c.client, httpClient) {
		t.Errorf("HTTP clients are not equal. Injected %+v, got %+v", httpClient, c.client)
	}
}

func TestNewClient_WithServices(t *testing.T) {
	testSetEnv(t)
	c, err := NewClient(nil, testTenableInstanceURL)

	if err != nil {
		t.Errorf("Got an error: %s", err)
	}
	if c == nil {
		t.Error("Client should not be nil")
	}
	// if c.Authentication == nil {
	// 	t.Error("No AuthenticationService provided")
	// }
}

func TestCheckResponse(t *testing.T) {
	codes := []int{
		http.StatusOK, http.StatusPartialContent, 299,
	}

	for _, c := range codes {
		r := &http.Response{
			StatusCode: c,
		}
		if err := CheckResponse(r); err != nil {
			t.Errorf("CheckResponse throws an error: %s", err)
		}
	}
}
