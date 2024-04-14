package auth

import (
	"fmt"
	"log"
	"testing"
)

type Test struct {
	name      string
	headers   map[string][]string
	expected  string
	errString string
}

func TestGetAPIKey(t *testing.T) {
	tests := []Test{
		{
			name:      "include auth header",
			headers:   map[string][]string{},
			expected:  "",
			errString: "no authorization header included",
		},
		{
			name: "auth is not api key",
			headers: map[string][]string{
				"Authorization": {"Bearer 134234121312312312312"},
			},
			expected:  "",
			errString: "malformed authorization header",
		},
		{
			name: "malformed auth key",
			headers: map[string][]string{
				"Authorization": {"ApiKey"},
			},
			expected:  "",
			errString: "malformed authorization header",
		},
		{
			name: "correct auth",
			headers: map[string][]string{
				"Authorization": {"ApiKey helloworld"},
			},
			expected:  "helloworld",
			errString: "",
		},
	}
	for _, test := range tests {
		key, err := GetAPIKey(test.headers)
		if test.expected != "" && key == test.expected {
			continue
		}
		if key != test.expected {
			fmt.Println("Test: " + test.name)
			log.Fatalf("did not get expected key\nExpected key: %s, Actual key: %s\n", test.expected, key)
		}
		if test.errString != "" && err == nil {
			fmt.Println("Test: " + test.name)
			log.Fatalf("expected error and got nothing\nExpected error: %s\n", test.errString)
		} else if test.errString == "" && err != nil {
			fmt.Println("Test: " + test.name)
			log.Fatalf("expected no error and got one\nActual error: %s\n", err.Error())
		} else if err != nil && err.Error() != test.errString {
			fmt.Println("Test: " + test.name)
			log.Fatalf("Something did not go as expected\nExpected key: %s, Actual key: %s\n", test.errString, err.Error())
		}
	}
}
