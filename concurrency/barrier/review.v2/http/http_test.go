package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/User-Agent",
		}
		got := captureBarrierOutput(endpoints...)
		assert.Contains(t, got, "Accept-Encoding")
		t.Log(got)
	})

	t.Run("Once endpoint incorrect", func(t *testing.T) {
		endpoints := []string{
			"http://malformedurl.com",
			"http://httpbin.org/User-Agent",
		}
		_ = endpoints
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/User-Agent",
		}
		_ = endpoints
	})
}
