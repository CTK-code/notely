package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		apiKey string
		error  error
	}{
		{
			apiKey: "testKey",
			error:  nil,
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			r, err := http.NewRequest("GET", "http://notAWebsite.com", nil)
			if err != nil {
				t.Error("Failed to create request")
			}
			r.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", c.apiKey))
			key, err := GetAPIKey(r.Header)
			if err != nil {
				t.Error("Could not get ApiKey")
			}

			if c.apiKey != key {
				t.Errorf("Keys %s != %s", c.apiKey, key)
			}
		})
	}
}
