package proto

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type contentLengthState struct{ state interface{} }

func (s *contentLengthState) ProtocolState() interface{}         { return s.state }
func (s *contentLengthState) SetProtocolState(state interface{}) { s.state = state }

func TestHasFullPayloadContentLengthPrefixes(t *testing.T) {
	for length := 0; length <= 12; length++ {
		t.Run(strconv.Itoa(length), func(t *testing.T) {
			headers := []byte(fmt.Sprintf("POST / HTTP/1.1\r\nHost: example.test\r\nContent-Length: %d\r\n\r\n", length))
			state := &contentLengthState{}
			for received := 0; received <= length; received++ {
				body := []byte(strings.Repeat("a", received))
				want := received == length
				if got := HasFullPayload(state, headers, body); got != want {
					t.Errorf("declared=%d received=%d: got %v, want %v", length, received, got, want)
				}
				joined := append(append([]byte(nil), headers...), body...)
				if got := HasFullPayload(nil, joined); got != want {
					t.Errorf("joined declared=%d received=%d: got %v, want %v", length, received, got, want)
				}
			}
		})
	}
}

func TestHasFullPayloadRejectsInvalidContentLength(t *testing.T) {
	for _, tc := range []struct{ name, value, body string }{
		{"empty", "", ""},
		{"non_digit", "x", ""},
		{"negative", "-1", ""},
		{"plus", "+1", ""},
		{"negative_zero", "-0", ""},
		{"invalid_suffix", "1x", "a"},
		{"overflow", "18446744073709551616", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			headers := []byte("POST / HTTP/1.1\r\nHost: example.test\r\nContent-Length: " + tc.value + "\r\n\r\n")
			payload := append(append([]byte(nil), headers...), []byte(tc.body)...)
			if HasFullPayload(nil, payload) {
				t.Error("invalid Content-Length accepted")
			}
			state := &contentLengthState{}
			for attempt := 0; attempt < 2; attempt++ {
				if HasFullPayload(state, headers, []byte(tc.body)) {
					t.Errorf("invalid Content-Length accepted with cached state on attempt %d", attempt)
				}
			}
		})
	}
}

func TestHasFullPayloadContentLengthCompatibility(t *testing.T) {
	for _, headers := range []string{
		"Content-Length: 0\r\n",
		"Content-Length: 00\r\n",
		"Content-Length: \t0 \t\r\n",
		"",
	} {
		payload := []byte("GET / HTTP/1.1\r\nHost: example.test\r\n" + headers + "\r\n")
		if !HasFullPayload(nil, payload) {
			t.Errorf("valid headers rejected: %q", headers)
		}
	}
}
