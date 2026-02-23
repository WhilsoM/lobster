package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	rr := httptest.NewRecorder()

	status := http.StatusOK
	data := map[string]string{"message": "hello"}

	WriteJSON(rr, status, data)

	if rr.Code != status {
		t.Errorf("except status %d, get %d", status, rr.Code)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("except Content-Type application/json, get %s", contentType)
	}

	expectedBody := `{"message":"hello"}`
	actualBody := strings.TrimSpace(rr.Body.String())
	if actualBody != expectedBody {
		t.Errorf("except %s, get %s", expectedBody, actualBody)
	}
}

