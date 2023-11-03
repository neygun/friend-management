package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// LoadTestJSONFile loads test json data from a file
func LoadTestJSONFile(t *testing.T, filename string) string {
	body, err := os.ReadFile(filename)
	require.NoError(t, err)
	return string(body)
}
