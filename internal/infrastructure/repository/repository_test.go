package repository

import (
	"database/sql"
	"os"
	"path"
	"testing"

	"github.com/koh-yoshimoto/go-testable-project/internal/infrastructure/db"
	"github.com/koh-yoshimoto/go-testable-project/pkg/config"
	"github.com/koh-yoshimoto/go-testable-project/pkg/testutils"
)

var testDB *sql.DB

func TestMain(m *testing.M) {

	config.InitConfig()

	err := db.InitDB()
	if err != nil {
		panic(err)
	}
	testDB = db.DB

	defer func() {
		// Truncate all tables
		_, err := db.DB.Exec("TRUNCATE TABLE tasks")
		if err != nil {
			panic(err)
		}
		db.DB.Close()
	}()

	os.Exit(m.Run())
}

func prepareTestDB(t *testing.T) {
	fixturesPath := path.Join(config.GetProjectRoot(), "/fixtures/testdata")
	err := testutils.LoadFixtures(testDB, fixturesPath)
	if err != nil {
		t.Fatalf("failed to load fixtures: %v", err)
	}
}
