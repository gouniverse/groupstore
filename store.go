package groupstore

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/gouniverse/base/database"
)

// == TYPE ====================================================================

type store struct {
	// groupTableName is the name of the group table
	groupTableName string

	// groupEntityRelationTableName is the name of the group entity relation table
	groupEntityRelationTableName string

	// db is the underlying database connection
	db *sql.DB

	// dbDriverName is the database driver name/type
	dbDriverName string

	// automigrateEnabled enables or disables automigration
	automigrateEnabled bool

	// debugEnabled enables or disables debug mode
	debugEnabled bool

	// sqlLogger is the sql logger used when debug mode is enabled
	sqlLogger *slog.Logger
}

// == INTERFACE ===============================================================

var _ StoreInterface = (*store)(nil) // verify it extends the interface

// PUBLIC METHODS ============================================================

// AutoMigrate auto-migrates the database schema
func (store *store) AutoMigrate() error {
	if store.db == nil {
		return errors.New("groupstore: database is nil")
	}

	sqlStr := store.sqlGroupTableCreate()

	if sqlStr == "" {
		return errors.New("groupstore: group table create sql is empty")
	}

	_, err := store.db.Exec(sqlStr)

	if err != nil {
		return err
	}

	sqlStr = store.sqlGroupEntityRelationTableCreate()

	if sqlStr == "" {
		return errors.New("groupstore: group entity relation table create sql is empty")
	}

	_, err = store.db.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil
}

// DB returns the underlying database connection
func (store *store) DB() *sql.DB {
	return store.db
}

// EnableDebug - enables or disables the debug mode
func (st *store) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

// logSql logs sql to the sql logger, if debug mode is enabled
func (store *store) logSql(sqlOperationType string, sql string, params ...interface{}) {
	if !store.debugEnabled {
		return
	}

	if store.sqlLogger != nil {
		store.sqlLogger.Debug("sql: "+sqlOperationType, slog.String("sql", sql), slog.Any("params", params))
	}
}

// toQuerableContext converts the context to a QueryableContext
func (store *store) toQuerableContext(ctx context.Context) database.QueryableContext {
	if database.IsQueryableContext(ctx) {
		return ctx.(database.QueryableContext)
	}

	return database.Context(ctx, store.db)
}
