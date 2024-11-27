package racer

import (
  "testing"
  "time"
  "net/http"
  "net/http/httptest"
)

func TestRacer(t *testing.T) {
  slowServer := createTestServer(20 * time.Millisecond)
  defer slowServer.Close()
  fastServer := createTestServer(0 * time.Millisecond)
  defer fastServer.Close()

  slowURL := slowServer.URL
  fastURL := fastServer.URL

  want := fastURL
  got := Racer(slowURL, fastURL)

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func createTestServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}

