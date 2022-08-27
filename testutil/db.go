package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	// テスト実行はdocker内で行うため
	port := 3306
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("todo:todo@tcp(todo-db:%d)/todo?parseTime=true", port),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "mysql")
}
