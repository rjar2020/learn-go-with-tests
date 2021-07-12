package racer

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRacer(t *testing.T) {

	assert := require.New(t)

	t.Run("returns the URL of the site that responds faster", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Microsecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		assert.NoError(err)
		assert.Equal(want, got)
	})

	t.Run("returns error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		assert.Error(err)
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func BenchmarkConfigurableRacer(b *testing.B)  {
	server := makeDelayedServer(2 * time.Millisecond)
	defer server.Close()

	for i := 0; i < b.N; i++ {
		_,_ = ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
	}
}

func ExampleConfigurableRacer() {
	serverA := makeDelayedUnstartedServer(0 * time.Microsecond)
	serverB := makeDelayedUnstartedServer(5 * time.Millisecond)
	
	overrideListenerAndStartServer(serverA, "127.0.0.1:8585")
	overrideListenerAndStartServer(serverB, "127.0.0.1:8586")
	
	defer serverA.Close()
	defer serverB.Close()

	response, _ := ConfigurableRacer(serverA.URL, serverB.URL, 20*time.Millisecond)
	fmt.Printf("%s", response)

	_, err := ConfigurableRacer(serverB.URL, serverB.URL, 1*time.Millisecond)
	fmt.Printf(" - %s", err)
	/* Output: http://127.0.0.1:8585 - timed out waiting for http://127.0.0.1:8586 and http://127.0.0.1:8586*/
}

func makeDelayedUnstartedServer(delay time.Duration) *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func overrideListenerAndStartServer(server *httptest.Server, url string)  {
	l, _ := net.Listen("tcp", url)
	_ = server.Listener.Close()
	server.Listener = l
	server.Start()
}
