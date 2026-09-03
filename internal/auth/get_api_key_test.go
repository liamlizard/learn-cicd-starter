package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input         http.Header
		want          string
		expectedError error
	}{
		"empty Auth header": {
			input:         http.Header{"Authorization": []string{""}},
			want:          "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		"short auth header": {
			input:         http.Header{"Authorization": []string{"ApiKey"}},
			want:          "",
			expectedError: ErrMalformedAuthHeader,
		},
		"Auth without 'ApiKey'": {
			input:         http.Header{"Authorization": []string{"NotApiKey 12345"}},
			want:          "",
			expectedError: ErrMalformedAuthHeader,
		},
		"Simple Correct": {
			input:         http.Header{"Authorization": []string{"ApiKey 12345"}},
			want:          "12345",
			expectedError: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, actualerr := GetAPIKey(tc.input)
			if !errors.Is(actualerr, tc.expectedError) {
				t.Fatalf("expected error: %v, got error: %v", tc.expectedError, actualerr)
			}
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
