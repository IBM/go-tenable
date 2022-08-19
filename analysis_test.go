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

func TestCheckAnalysisPost200(t *testing.T) {
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

	raw, err := ioutil.ReadFile("./mocks/analysis_get.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc("/rest/analysis", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/analysis")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(raw))

	})
	if testClient == nil {
		t.Fatal(fmt.Errorf("testclient nil"))
	}
	b := AnalysisBody{}
	user, resp, err := testClient.Analysis.Post(b)
	_ = resp
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal(errors.New("user is nil"))
	}
	if resp == nil {
		t.Fatal(errors.New("resp is nil"))
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf("Status code should be %d", http.StatusOK))
	}
}
