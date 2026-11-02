package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	rec := httptest.NewRecorder()

	greetHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
	var got map[string]string
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	if got["message"] != "hello" {
		t.Fatalf("message = %q, want %q", got["message"], "hello")
	}
}
