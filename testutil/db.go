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
	host := "todo-db"
	port := 3306
	if _, defined := os.LookupEnv("CI"); defined {
		host = "127.0.0.1"
		port = 3306
	}
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("todo:todo@tcp(%s:%d)/todo?parseTime=true", host, port),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "mysql")
}
