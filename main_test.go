package main

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestUUIDHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/uuid", nil)
	w := httptest.NewRecorder()
	uuidHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if v := resp.StatusCode; v != 200 {
		t.Error("expected 200 got", v)
	}

	if _, err := uuid.ParseBytes(body); err != nil {
		t.Error("could not parse body into uuid value:", err)
	}

}

func TestSHAHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/sha", bytes.NewBufferString("hello, world"))
	w := httptest.NewRecorder()
	shaHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if v := resp.StatusCode; v != 200 {
		t.Error("expected 200 got", v)
	}

	// echo -n "hello, world" | shasum -a 256
	if string(body) != "09ca7e4eaa6e8ae9c7d261167129184883644d07dfba7cbfbc4c8a2e08360d5b" {
		t.Error("got unexpected body hash", string(body))
	}

}
