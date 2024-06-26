package test

import (
	"os"
	"testing"

	"github.com/neygun/friend-management/pkg/db"
	"github.com/stretchr/testify/require"
)

// LoadTestSQLFile loads test sql data from a file
func LoadTestSQLFile(t *testing.T, tx db.ContextExecutor, filename string) {
	body, err := os.ReadFile(filename)
	require.NoError(t, err)

	_, err = tx.Exec(string(body))
	require.NoError(t, err)
}
