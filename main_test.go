package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestViewHandler(t *testing.T) {
	// Create a dummy page
	data := []byte("This is a test page")
	os.WriteFile("test_page.txt", data, 0644)
	defer os.Remove("test_page.txt")

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/view/test_page", nil)

	// Record the response
	rec := httptest.NewRecorder()
	viewHandler(rec, req)

	// Check the status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rec.Code)
	}

	// Check the response body
	expectedBody := "<h1>test_page</h1><div>This is a test page</div>"
	if !bytes.Equal(rec.Body.Bytes(), []byte(expectedBody)) {
		t.Errorf("Expected body %s, got %s", expectedBody, rec.Body.String())
	}
}

func TestLoadPage(t *testing.T) {
	// Test loading a valid page
	data := []byte("This is a test page")
	os.WriteFile("test_page.txt", data, 0644)
	defer os.Remove("test_page.txt")
	page, err := loadPage("test_page")
	if err != nil {
		t.Errorf("Error loading page: %v", err)
	}

	if page.Title != "test_page" {
		t.Errorf("Expected title 'test_page', got %s", page.Title)
	}

	if !bytes.Equal(page.Body, []byte("This is a test page")) {
		t.Errorf("Expected body 'This is a test page', got %s", page.Body)
	}

	// Test loading a non-existent page
	_, err = loadPage("not_found")
	if err == nil {
		t.Errorf("Expected error loading non-existent page")
	}
}
