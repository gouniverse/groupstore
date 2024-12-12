package groupstore

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/utils"
	_ "modernc.org/sqlite"
)

func initDB(filepath string) (*sql.DB, error) {
	if filepath != ":memory:" && utils.FileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			return nil, err
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initStore(filepath string) (StoreInterface, error) {
	db, err := initDB(filepath)

	if err != nil {
		return nil, err
	}

	store, err := NewStore(NewStoreOptions{
		DB:                   db,
		GroupTableName:       "groups_group_table",
		EntityGroupTableName: "groups_entity_group_table",
		AutomigrateEnabled:   true,
		DebugEnabled:         true,
		SqlLogger:            slog.New(slog.NewTextHandler(os.Stdout, nil)),
	})

	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("unexpected nil store")
	}

	return store, nil
}

func TestStoreWithTx(t *testing.T) {
	store, err := initStore("test_store_with_tx.db")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	db := store.DB()

	if db == nil {
		t.Fatal("unexpected nil db")
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	tx, err := db.Begin()

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tx == nil {
		t.Fatal("unexpected nil tx")
	}

	txCtx := database.Context(context.Background(), tx)

	// create group
	group := NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("GROUP_HANDLE").
		SetTitle("GROUP_TITLE")

	err = store.GroupCreate(txCtx, group)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// update group
	group.SetTitle("GROUP_TITLE_2")
	err = store.GroupUpdate(txCtx, group)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check group
	groupFound, errFind := store.GroupFindByID(database.Context(context.Background(), db), group.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if groupFound != nil {
		t.Fatal("Group MUST be nil, as transaction not committed")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check group
	groupFound, errFind = store.GroupFindByID(database.Context(context.Background(), db), group.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if groupFound == nil {
		t.Fatal("Group MUST be not nil, as transaction committed")
	}

	if groupFound.Title() != "GROUP_TITLE_2" {
		t.Fatal("Group MUST be GROUP_TITLE_2, as transaction committed")
	}
}
