package groupstore

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/gouniverse/sb"
)

// NewStoreOptions define the options for creating a new block store
type NewStoreOptions struct {
	// GroupTableName is the name of the group table
	GroupTableName string

	// EntityGroupTableName is the name of the entity to group relation table
	EntityGroupTableName string

	// DB is the underlying database connection
	DB *sql.DB

	// DbDriverName is the database driver name/type
	DbDriverName string

	// AutomigrateEnabled indicates whether to automatically migrate the database
	AutomigrateEnabled bool

	// DebugEnabled enables or disables the debug mode
	DebugEnabled bool

	// SqlLogger is the sql statement logger when debug mode is enabled, defaults to the default logger
	SqlLogger *slog.Logger
}

// NewStore creates a new block store
func NewStore(opts NewStoreOptions) (StoreInterface, error) {
	if opts.GroupTableName == "" {
		return nil, errors.New("group store: GroupTableName is required")
	}

	if opts.EntityGroupTableName == "" {
		return nil, errors.New("group store: EntityGroupTableName is required")
	}

	if opts.DB == nil {
		return nil, errors.New("shop store: DB is required")
	}

	if opts.DbDriverName == "" {
		opts.DbDriverName = sb.DatabaseDriverName(opts.DB)
	}

	if opts.SqlLogger == nil {
		opts.SqlLogger = slog.Default()
	}

	store := &store{
		groupTableName:       opts.GroupTableName,
		entityGroupTableName: opts.EntityGroupTableName,
		automigrateEnabled:   opts.AutomigrateEnabled,
		db:                   opts.DB,
		dbDriverName:         opts.DbDriverName,
		debugEnabled:         opts.DebugEnabled,
		sqlLogger:            opts.SqlLogger,
	}

	if store.automigrateEnabled {
		err := store.AutoMigrate()

		if err != nil {
			return nil, err
		}
	}

	return store, nil
}
