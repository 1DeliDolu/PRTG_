package plugin

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const prtgUserAgent = "maxmarkusprogram-prtg-datasource"

// isContentXML prüft, ob der Content-Type XML oder HTML mit UTF-8 ist.
func isContentXML(header http.Header) bool {
	ct := header.Get("Content-Type")
	return strings.HasPrefix(ct, "text/xml") || strings.HasPrefix(ct, "text/html")
}

// getHTTPBody führt eine HTTP GET-Anfrage durch und gibt den Body und Header zurück.
func getHTTPBody(url string, timeoutMillis int64) ([]byte, http.Header, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeoutMillis) * time.Millisecond,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("fehler beim Erstellen der GET-Anfrage: %w", err)
	}
	req.Header.Set("User-Agent", prtgUserAgent)

	ctx, cancel := context.WithTimeout(req.Context(), time.Duration(timeoutMillis)*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("fehler bei der HTTP-Anfrage: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return nil, nil, fmt.Errorf("falscher Benutzername und/oder Passwort | HTTP-Status: %d", res.StatusCode)
	}
	if res.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("HTTP-Status nicht OK: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("fehler beim Lesen des Response-Bodys: %w", err)
	}
	return body, res.Header, nil
}

// getPrtgResponse ruft eine URL auf und unmarshalt den Response-Body entweder als XML oder JSON.
func getPrtgResponse(url string, timeoutMillis int64, v interface{}) error {
	body, header, err := getHTTPBody(url, timeoutMillis)
	if err != nil {
		return err
	}

	if isContentXML(header) {
		decoder := xml.NewDecoder(strings.NewReader(string(body)))
		decoder.Strict = false
		if err := decoder.Decode(v); err != nil {
			return fmt.Errorf("fehler beim Unmarshaln der XML-Antwort: %w", err)
		}
		return nil
	}

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("fehler beim Unmarshaln der JSON-Antwort: %w", err)
	}
	return nil
}
// ok