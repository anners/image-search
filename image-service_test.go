package main

import (
	_"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	health(response, request)

	if response.Code != http.StatusOK {
        t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
    }

    assert.Equal(t, "I am alive!", response.Body.String())
}

func TestImageReturnsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/image?search=clouds", nil)
	response := httptest.NewRecorder()

	image(response, request)

	ct := response.HeaderMap["Content-Type"][0]
	if !strings.EqualFold(ct, "application/json") {
		t.Fatalf("Content-Type does not equal 'application/json'")
	}
}

func TestImageSearch(t *testing.T) {
	request, _ := http.NewRequest("GET", "/image?search=dogs", nil)
	response := httptest.NewRecorder()

	image(response, request)

	assert.Contains(t, response.Body.String(), "link")
}

func TestCreateSeachURL(t *testing.T) {
	query := queryParams{"gopher"}
	imageSearch := query.createSeachURL()

	assert.Contains(t, imageSearch, "gopher", "URL should contain gopher")
}