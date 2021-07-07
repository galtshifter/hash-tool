package main

import (
	"testing"
)

func TestPrepareURLs(t *testing.T) {
	testCases := []struct {
		name            string
		endpoints       []string
		noErrorExpected bool
	}{
		{
			name:            "empty endpoints",
			endpoints:       []string{},
			noErrorExpected: true,
		},
		{
			name:            "empty string",
			endpoints:       []string{""},
			noErrorExpected: false,
		},
		{
			name:            "valid link",
			endpoints:       []string{"http://google.com/"},
			noErrorExpected: true,
		},
		{
			name:            "valid link no slash in the end",
			endpoints:       []string{"http://google.com"},
			noErrorExpected: true,
		},
		{
			name:            "valid link no scheme",
			endpoints:       []string{"google.com"},
			noErrorExpected: true,
		},
		{
			name:            "not valid link no tld",
			endpoints:       []string{"googlecom"},
			noErrorExpected: false,
		},
		{
			name:            "not valid link space",
			endpoints:       []string{"google com"},
			noErrorExpected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, actual := prepareURLs(tC.endpoints)
			if (actual == nil) != tC.noErrorExpected {
				t.Error()
			}
		})
	}
}
