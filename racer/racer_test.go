package racer

import (
  "testing"
  "time"
  "net/http"
  "net/http/httptest"
)

func TestRacer(t *testing.T) {
  slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(20 * time.Millisecond)
    w.WriteHeader(http.StatusOK)
  }))
  fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
  }))

  //slowUrl := "http://www.facebook.com"
  //fastUrl := "http://www.quii.dev"
  slowURL := slowServer.URL
  fastURL := fastServer.URL

  want := fastURL
  got := Racer(slowURL, fastURL)

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }

  slowServer.Close()
  fastServer.Close()
}

