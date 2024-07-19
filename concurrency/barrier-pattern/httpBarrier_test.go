package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	// go test -run=TestBarrier/Correct_endpoints -v
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "user-agent") {
			t.Fail()
		}
		t.Log(result)
	})

	// go test -run=TestBarrier/One_endpoint_incorrect -v
	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Error") {
			t.Fail()
		}
		t.Log(result)
	})
	// go test -run=TestBarrier/Verify_short_timeout -v
	t.Run("Verify short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	barrier(endpoints...)
	writer.Close()
	temp := <-out
	return temp
}
