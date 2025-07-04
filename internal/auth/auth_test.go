package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
		err   error
	}
	tests := []test{
		{input: http.Header{"Authorization": []string{"ApiKey 12345"}}, want: "12345", err: nil},
		{input: http.Header{"": []string{""}}, want: "", err: ErrNoAuthHeaderIncluded},
		{input: http.Header{"Authorization": []string{"ApiKey"}}, want: "", err: errors.New("malformed authorization header")},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil && tc.err != nil && err.Error() == tc.err.Error() {
			t.Fatalf("wrong error type, expected %v, got %v", tc.err, err)

		}
		if got != tc.want {
			t.Fatalf("expected: %v - %v, got %v - %v", tc.want, tc.err, got, err)
		}
	}
}
