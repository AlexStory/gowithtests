package racer

import (
	"fmt"
	"net/http"
	"time"
)

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out after waiting for %s and %s", a, b)
	}
}

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, time.Second*10)
}

func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
