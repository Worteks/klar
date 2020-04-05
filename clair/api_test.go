package clair

import (
	"testing"
	"time"
)

func TestNewAPIV1(t *testing.T) {
	cases := []struct {
		url      string
		expected string
	}{
		{
			url:      "http://localhost:6060",
			expected: "http://localhost:6060",
		},
		{
			url:      "http://localhost",
			expected: "http://localhost:6060",
		},
		{
			url:      "localhost",
			expected: "http://localhost:6060",
		},
		{
			url:      "https://localhost:6060",
			expected: "https://localhost:6060",
		},
		{
			url:      "https://localhost",
			expected: "https://localhost:6060",
		},
	}
	for _, tc := range cases {
		api := newAPIV1(tc.url, time.Minute)
		if api.url != tc.expected {
			t.Errorf("expected %s got %s", api.url, tc.expected)
		}
	}
}
