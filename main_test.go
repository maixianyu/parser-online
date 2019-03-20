package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
)

func TestCljHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(cljHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("cljHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Welcome to Clojure.`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("cljHandler return unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/clojure")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status want %v got %v", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Welcome to Clojure."
	if expected != respString {
		t.Errorf("body want %v got %v", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL + "/clojure", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status want %v got %v", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""
	if expected != respString {
		t.Errorf("body want %v got %v", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %v got %v", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("content type want %s got %s", expectedContentType, contentType)
	}
}



