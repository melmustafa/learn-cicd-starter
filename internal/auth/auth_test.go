package auth

import (
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
				"Authorization": {"ApiKey  "},
			},
			expected:  "",
			errString: "malformed authorization header",
		},
		{
			name: "include auth header",
			headers: map[string][]string{
				"Authorization": {"ApiKey helloworld"},
			},
			expected:  "helloworld",
			errString: "",
		},
	}
	for _, test := range tests {
		key, err := GetAPIKey(test.headers)
		if key == test.expected {
			continue
		}
		if key != test.expected {
			log.Fatalln("did not get expected key")
			log.Fatalf("Expected key: %s, Actual key: %s\n", test.expected, key)
		}
		if test.errString != "" && err == nil {
			log.Fatalln("expected error and got nothing")
			log.Fatalf("Expected error: %s\n", test.errString)
		} else if test.errString == "" && err != nil {
			log.Fatalln("expected no error and got one")
			log.Fatalf("Actual error: %s\n", err.Error())
		} else if err != nil && err.Error() != test.errString {
			log.Fatalln("Something did not go as expected")
			log.Fatalf("Expected key: %s, Actual key: %s\n", test.expected, key)
		}
	}
}
