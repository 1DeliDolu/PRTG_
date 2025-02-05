package plugin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPrtgVersion(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.HandleFunc(GetSensorDetailsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"prtgversion": "22.3.79.2108"})
	})
	httpServer := httptest.NewServer(mux)
	defer httpServer.Close()

	client := NewClient(httpServer.URL, "testuser", "testpass")
	version, err := client.GetPrtgVersion()
	if err != nil {
		t.Fatalf("Fehler beim Abrufen der PRTG-Version: %v", err)
	}

	expectedVersion := "22.3.79.2108"
	if version != expectedVersion {
		t.Errorf("Erwartete Version %v, aber erhalten %v", expectedVersion, version)
	}
}

func TestGetSensorDetail(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	mux.HandleFunc(GetSensorDetailsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"sensordata": map[string]string{
				"name":       "CPU Load",
				"sensortype": "load",
			},
		})
	})
	httpServer := httptest.NewServer(mux)
	defer httpServer.Close()

	client := NewClient(httpServer.URL, "testuser", "testpass")
	sensorData, err := client.GetSensorDetail(1234)
	if err != nil {
		t.Fatalf("Fehler beim Abrufen der Sensor-Details: %v", err)
	}

	expectedSensorName := "CPU Load"
	if sensorData.Name != expectedSensorName {
		t.Errorf("Erwarteter Sensorname %v, aber erhalten %v", expectedSensorName, sensorData.Name)
	}
}
// ok