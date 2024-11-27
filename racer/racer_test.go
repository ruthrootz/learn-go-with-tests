package racer

import (
  "testing"
  "time"
  "net/http"
  "net/http/httptest"
)

func TestRacer(t *testing.T) {

  t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
    slowServer := createTestServer(20 * time.Millisecond)
    defer slowServer.Close()
    fastServer := createTestServer(0 * time.Millisecond)
    defer fastServer.Close()

    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got, _ := Racer(slowURL, fastURL)

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }  
  })

  t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
    serverA := createTestServer(11 * time.Second)
    defer serverA.Close()
    serverB := createTestServer(12 * time.Second)
    defer serverB.Close()

    _, err := Racer(serverA.URL, serverB.URL)

    if err == nil {
      t.Errorf("expected error but didn't get one")
    }
  })

}

func createTestServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}

