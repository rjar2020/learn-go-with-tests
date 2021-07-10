package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

//Racer uses ConfigurableRacer to return the URL of the faster server or an error if the HTTP call takes more than 10 seconds
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

//ConfigurableRacer returns the URL of the faster server or an error if the HTTP call takes more than "timeout" seconds
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
