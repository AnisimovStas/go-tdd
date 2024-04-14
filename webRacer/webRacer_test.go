package webRacer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayerServer(time.Millisecond * 20)
		fastServer := makeDelayerServer(time.Millisecond * 0)
		defer slowServer.Close()
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
		slowServer := makeDelayerServer(time.Second * 12)
		defer slowServer.Close()

		slowURL := slowServer.URL

		_, err := ConfigurableRacer(slowURL, slowURL, time.Millisecond*10)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayerServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
