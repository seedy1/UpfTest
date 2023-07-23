package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seedy1/UpfTest/routes"
)

// test /analysis endpoint
func TestAnalysisGetRoute(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/analysis?duration=1s&dimension=likes", nil)
	w := httptest.NewRecorder()
	routes.HandleSSE(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Error("Expected 200 status code")
	}

}

func TestAnalysisPostRoute(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/analysis?duration=1s&dimension=likes", nil)
	w := httptest.NewRecorder()
	routes.HandleSSE(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 405 {
		t.Error("Expected 200 status code")
	}

}
