/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestCheckRepositoryGet200(t *testing.T) {
	testSetEnv(t)
	defer testTeardownEnv(t)
	setup()
	defer teardown()
	if testClient == nil {
		t.Fatal(fmt.Errorf("testclient nil"))
	}
	tmp := os.Getenv("SC05_ACCESS_KEY")
	if tmp == "" {
		t.Fatal(errors.New("env error"))
	}

	raw, err := ioutil.ReadFile("./mocks/repository_get_all_no_filter.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc("/rest/repository", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/repository")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(raw))

	})
	if testClient == nil {
		t.Fatal(fmt.Errorf("testclient nil"))
	}
	repos, resp, err := testClient.Repository.Get("All", "")
	if err != nil {
		t.Fatal(err)
	}
	if repos == nil {
		t.Fatal(errors.New("repos is nil"))
	}
	if resp == nil {
		t.Fatal(errors.New("resp is nil"))
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf("Status code should be %d", http.StatusOK))
	}
}
