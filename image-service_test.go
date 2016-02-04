package main

import (
	_"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ( 
	response *httptest.ResponseRecorder
)

func setup() {
	response = httptest.NewRecorder()
}

func teardown() {
	response.Flush()
}

func TestHealth(t *testing.T) {
	setup()
	defer teardown()
	request, _ := http.NewRequest("GET", "/", nil)

	health(response, request)

	if response.Code != http.StatusOK {
        t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
    }

    assert.Equal(t, "I am alive!", response.Body.String())
}

func TestImageReturnsJSON(t *testing.T) {
	setup()
	defer teardown()
	request, _ := http.NewRequest("GET", "/image?search=clouds", nil)

	image(response, request)

	ct := response.Header().Get("Content-Type")
	assert.Equal(t, ct, "application/json", "Content-Type should be json")
}

func TestImageSearch(t *testing.T) {
	setup()
	defer teardown()
	request, _ := http.NewRequest("GET", "/image?search=dogs", nil)

	image(response, request)

	assert.Contains(t, response.Body.String(), "link", "body should contain link:")
}

func TestCreateSeachURL(t *testing.T) {
	query := queryParams{"gopher"}
	imageSearch := query.createSeachURL()

	assert.Contains(t, imageSearch, "gopher", "URL should contain gopher")
}