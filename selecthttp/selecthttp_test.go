package selecthttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("returns an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
		server := makeDelayedServer(11 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Second)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("returns the server slow server URL", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expext an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
