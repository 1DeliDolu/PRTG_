package plugin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetHttpBody(t *testing.T) {
	t.Helper()

	// Ungültige URL
	invalidURL := " http://localhost"
	timeout := int64(10000)
	_, _, err := getHTTPBody(invalidURL, timeout)
	if err == nil {
		t.Errorf("Es sollte ein Fehler auftreten (bei NewRequest()) für URL %v", invalidURL)
	}

	// Server nicht erreichbar
	nonExistentServerURL := "http://localhost"
	_, _, err = getHTTPBody(nonExistentServerURL, timeout)
	if err == nil {
		t.Errorf("Es sollte ein Fehler auftreten (bei Send Request), wenn der Server nicht läuft: %v", err)
	}
}

func TestRespStatusCode(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(GetSensorDetailsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{}")
	})
	httpServer := setupTestServer(mux)
	defer httpServer.Close()

	serverURL, _ := url.Parse(httpServer.URL)

	path := "wrong/path"
	u := fmt.Sprintf("%v/%v", serverURL, path)
	timeout := int64(10000)
	_, _, err := getHTTPBody(u, timeout)
	if err == nil {
		t.Errorf("Erwarteter Fehler für falschen Pfad: %v", err)
	}
}

func TestUnauthorizedAccess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(GetSensorDetailsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("LoginAgain", "true")
		w.WriteHeader(http.StatusUnauthorized)
	})
	httpServer := setupTestServer(mux)
	defer httpServer.Close()

	serverURL, _ := url.Parse(httpServer.URL)

	path := GetSensorDetailsEndpoint
	u := fmt.Sprintf("%v/%v", serverURL, path)
	timeout := int64(10000)
	_, _, err := getHTTPBody(u, timeout)
	if err == nil {
		t.Errorf("Erwarteter Fehler für Unauthorized Access: %v", err)
	}
}
// setupTestServer erstellt einen Testserver.
func setupTestServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

//ok 