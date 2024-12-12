package groupstore

import (
	"github.com/gouniverse/sb"
)

// sqlGroupTableCreate returns a SQL string for creating the group table
func (st *store) sqlGroupTableCreate() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(st.db)).
		Table(st.groupTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			PrimaryKey: true,
			Length:     40,
		}).
		Column(sb.Column{
			Name:   COLUMN_STATUS,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   COLUMN_HANDLE,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   COLUMN_TITLE,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name: COLUMN_METAS,
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: COLUMN_MEMO,
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name:   COLUMN_CREATED_AT,
			Type:   sb.COLUMN_TYPE_DATETIME,
			Length: 0,
		}).
		Column(sb.Column{
			Name:   COLUMN_UPDATED_AT,
			Type:   sb.COLUMN_TYPE_DATETIME,
			Length: 0,
		}).
		Column(sb.Column{
			Name:   COLUMN_SOFT_DELETED_AT,
			Type:   sb.COLUMN_TYPE_DATETIME,
			Length: 0,
		}).
		CreateIfNotExists()

	return sql
}

// sqlEntityGroupTableCreate returns a SQL string for creating the  entity to group relation table
func (st *store) sqlEntityGroupTableCreate() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(st.db)).
		Table(st.entityGroupTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			PrimaryKey: true,
			Length:     40,
		}).
		Column(sb.Column{
			Name:   COLUMN_ENTITY_TYPE,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 80,
		}).
		Column(sb.Column{
			Name:   COLUMN_ENTITY_ID,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   COLUMN_GROUP_ID,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name: COLUMN_METAS,
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: COLUMN_MEMO,
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name:   COLUMN_CREATED_AT,
			Type:   sb.COLUMN_TYPE_DATETIME,
			Length: 0,
		}).
		Column(sb.Column{
			Name:   COLUMN_UPDATED_AT,
			Type:   sb.COLUMN_TYPE_DATETIME,
			Length: 0,
		}).
		Column(sb.Column{
			Name:   COLUMN_SOFT_DELETED_AT,
			Type:   sb.COLUMN_TYPE_DATETIME,
			Length: 0,
		}).
		CreateIfNotExists()

	return sql
}
