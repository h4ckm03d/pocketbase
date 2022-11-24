package core

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

func connectDB(dbPath string) (*dbx.DB, error) {
	return dbx.MustOpen("postgres", os.Getenv("DATABASE_URL"))
}
