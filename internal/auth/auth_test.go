package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers  http.Header
		expected string
	}{
		"simple": {
			headers:  http.Header{"Authorization": {"ApiKey 12345"}},
			expected: "12345",
		},
		"missing apikey": {
			headers:  http.Header{"Authorization": {"Bearer 98765"}},
			expected: "malformed authorization header",
		},
		"missing authorization": {
			headers:  http.Header{"Content-Type": {"text/html"}},
			expected: "no authorization header",
		},
		"missing header": {
			headers:  http.Header{},
			expected: "no authorization header",
		},
		"missing header value": {
			headers:  http.Header{"Authorization": {}},
			expected: "no authorization header",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := GetAPIKey(tc.headers)
			if err != nil {
				if strings.Contains(err.Error(), tc.expected) {
					return
				} else {
					t.Error("received error:", err)
				}
			} else if result != tc.expected {
				t.Errorf("expected %s, found %s", tc.expected, result)
			}
		})
	}
}
