package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestViewHandler(t *testing.T) {
	data := []byte("This is a test page")
	os.WriteFile("test_page.txt", data, 0644)
	defer os.Remove("test_page.txt")

	req := httptest.NewRequest(http.MethodGet, "/view/test_page", nil)

	rec := httptest.NewRecorder()
	viewHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rec.Code)
	}

	expectedBody := "<h1>test_page</h1><div>This is a test page</div>"
	if !bytes.Equal(rec.Body.Bytes(), []byte(expectedBody)) {
		t.Errorf("Expected body %s, got %s", expectedBody, rec.Body.String())
	}
}

func TestLoadPage(t *testing.T) {
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

	var nil_page string
	_, err = loadPage(nil_page)
	if err == nil {
		t.Errorf("Expected error loading non-existent page")
	}
}
