package godb_test

import (
	"os"
	"testing"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
	"github.com/samonzeweb/godb/dbtests/common"

	. "github.com/smartystreets/goconvey/convey"
)

func fixturesSetupPostgreSQL(t *testing.T) (*godb.DB, func()) {
	if os.Getenv("GODB_POSTGRESQL") == "" {
		t.Skip("Don't run PostgreSQL test, GODB_POSTGRESQL not set")
	}

	db, err := godb.Open(postgresql.Adapter, os.Getenv("GODB_POSTGRESQL"))
	if err != nil {
		t.Fatal(err)
	}

	// Enable logger if needed
	//db.SetLogger(log.New(os.Stderr, "", 0))

	createTable :=
		`create temporary table if not exists books (
		id 						serial primary key,
		title     		varchar(128) not null,
		author    	  varchar(128) not null,
		published			date not null);
	`
	_, err = db.CurrentDB().Exec(createTable)
	if err != nil {
		t.Fatal(err)
	}

	fixturesTeardown := func() {
		dropTable := "drop table if exists books"
		_, err := db.CurrentDB().Exec(dropTable)
		if err != nil {
			t.Fatal(err)
		}
	}

	return db, fixturesTeardown
}

func TestStatementsPostgreSQL(t *testing.T) {
	Convey("A DB for a PostgreSQL database", t, func() {
		db, teardown := fixturesSetupPostgreSQL(t)
		defer teardown()

		Convey("The common tests must pass", func() {
			common.StatementsTests(db, t)
		})
	})
}

func TestStructsPostgreSQL(t *testing.T) {
	Convey("A DB for a PostgreSQL database", t, func() {
		db, teardown := fixturesSetupPostgreSQL(t)
		defer teardown()

		Convey("The common tests must pass", func() {
			common.StructsTests(db, t)
		})
	})
}
