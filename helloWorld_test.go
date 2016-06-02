// There are 3 unit tests in this file
// TestHi tests the response when regular json data sent
// TestHiWithEmptyFields tests the response when json data sent with empty fields
// TestHiWithNilBody tests the response when request body is nil

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// prototype for handler
type HandleTester func(
	method string,
	data []byte,
) *httptest.ResponseRecorder

func GenerateHandleTester(t *testing.T, handleFunc http.Handler) HandleTester {
	return func(
		method string,
		data []byte,
	) *httptest.ResponseRecorder {

		// initilization of a reader
		buf := bytes.NewReader(data)

		req, err := http.NewRequest(
			method,
			"",
			buf,
		)
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
		// w initializes response writer from the test request
		w := httptest.NewRecorder()
		// doing the request
		handleFunc.ServeHTTP(w, req)
		return w
	}
}

func TestHi(t *testing.T) {

	firstName := "Barack"
	lastName := "Obama"

	sampleUser := User{FirstName: firstName, LastName: lastName}
	fmt.Println("sampleUser", sampleUser)
	sampleUserJson, _ := json.Marshal(&sampleUser)
	fmt.Println("sampleUser (JSON, actually this pure []byte):")
	fmt.Println(sampleUserJson)
	fmt.Println("sampleUser (JSON, strinfii):")
	fmt.Println(string(sampleUserJson))
	//initialization of handle tester
	test := GenerateHandleTester(t, http.HandlerFunc(HiHandler))
	// w is http response
	w := test("POST", sampleUserJson)
	fmt.Println("Response is", w)
	fmt.Println("Response body is", w.Body)
	s := w.Body.String()

	if w.Code != http.StatusOK {
		t.Errorf("/hi didin't return %v", http.StatusOK)
	}

	greet := GreetUser{"Hi " + firstName + " " + lastName}
	greetJson, _ := json.Marshal(&greet)

	// strings.Compare(s, string(greetJson)) == -1 <== couldn't make this work
	// that's why used strings.Index

	found := strings.Index(s, string(greetJson))

	if found == -1 {
		t.Errorf("/hi didn't return the right JSON output")
	}
}

func TestHiWithEmptyFields(t *testing.T) {

	firstName := ""
	lastName := ""

	sampleUser := User{FirstName: firstName, LastName: lastName}

	sampleUserJson, _ := json.Marshal(&sampleUser)

	test := GenerateHandleTester(t, http.HandlerFunc(HiHandler))

	w := test("POST", sampleUserJson)
	s := w.Body.String()

	if w.Code != 400 {
		t.Errorf("/hi didin't return %v", http.StatusOK)
	}

	if strings.Index(s, "Please fill both fields") == -1 {
		t.Errorf("/hi didin't return the error message")
	}

}

func TestHiWithNilBody(t *testing.T) {

	test := GenerateHandleTester(t, http.HandlerFunc(HiHandler))
	w := test("POST", nil)
	fmt.Println("Response is", w)
	fmt.Println("Response body is", w.Body)
	if w.Code != 400 {
		t.Errorf("/hi didn't return 400")
	}
}
