package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/noc-tech/todo/datastore"
)

func Test_Resource_List(t *testing.T) {
	r := chi.NewRouter()
	ds := datastore.NewDatastore([]*datastore.Todo{{Id: "id1", Body: "to"}})
	r.Mount("/todos", todoResource{ds}.Routes())
	ts := httptest.NewServer(r)
	defer ts.Close()
	or, resp := testRequest(t, ts, "GET", "/todos", nil)
	if resp != `[{"id":"id1","text":"to"}]` && or.StatusCode == http.StatusOK {
		t.Fatalf(resp)
	}
}

func Test_Resource_Create(t *testing.T) {
	r := chi.NewRouter()
	ds := datastore.NewDatastore([]*datastore.Todo{{Id: "id1", Body: "to"}})
	r.Mount("/todos", todoResource{ds}.Routes())
	ts := httptest.NewServer(r)
	defer ts.Close()
	or, resp := testRequest(t, ts, "POST", "/todos", bytes.NewBuffer([]byte(`{"id":"id1","text":"to"`)))
	if resp != `{"error":"can't decode body"}` && or.StatusCode == http.StatusBadRequest {
		t.Fatalf(resp)
	}
	or, resp = testRequest(t, ts, "POST", "/todos", bytes.NewBuffer([]byte(`{"id":"id1","text":"to"}`)))
	if resp != `` && or.StatusCode == http.StatusCreated {
		t.Fatalf(resp)
	}
}

func Test_Resource_Delete(t *testing.T) {
	r := chi.NewRouter()
	ds := datastore.NewDatastore([]*datastore.Todo{{Id: "id1", Body: "to"}})
	r.Mount("/todos", todoResource{ds}.Routes())
	ts := httptest.NewServer(r)
	defer ts.Close()
	or, resp := testRequest(t, ts, http.MethodDelete, "/todos/id2", nil)
	if resp != `{"error":"can't delete todo"}` && or.StatusCode == http.StatusInternalServerError {
		t.Fatalf(resp)
	}
	or, resp = testRequest(t, ts, http.MethodDelete, "/todos/id1", nil)
	if resp != `` && or.StatusCode == http.StatusOK {
		t.Fatalf(resp)
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, strings.TrimRight(string(respBody), "\r\n")
}
