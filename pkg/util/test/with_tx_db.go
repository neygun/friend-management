package test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/neygun/friend-management/pkg/db"
	"github.com/neygun/friend-management/pkg/util"
	"github.com/stretchr/testify/require"
)

func WithTxDB(t *testing.T, testFunc func(tx db.ContextExecutor)) {
	ctx := context.Background()

	os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/fm-pg?sslmode=disable")
	db, err := util.Init()
	require.NoError(t, err)

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	require.NoError(t, err)
	defer tx.Rollback()

	testFunc(tx)
}
