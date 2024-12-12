package groupstore

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func (store *store) GroupCount(ctx context.Context, options GroupQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.groupSelectQuery(options)

	sqlStr, params, errSql := q.Prepared(true).
		Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	store.logSql("select", sqlStr, params...)

	mapped, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, params...)
	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
}

func (store *store) GroupCreate(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("group is nil")
	}

	group.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	group.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := group.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.groupTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("insert", sqlStr, params...)

	if store.db == nil {
		return errors.New("groupstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	if err != nil {
		return err
	}

	group.MarkAsNotDirty()

	return nil
}

func (store *store) GroupDelete(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("group is nil")
	}

	return store.GroupDeleteByID(ctx, group.ID())
}

func (store *store) GroupDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("group id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.groupTableName).
		Prepared(true).
		Where(goqu.C(COLUMN_ID).Eq(id)).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("delete", sqlStr, params...)

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	return err
}

func (store *store) GroupFindByHandle(ctx context.Context, handle string) (group GroupInterface, err error) {
	if handle == "" {
		return nil, errors.New("group handle is empty")
	}

	query := NewGroupQuery().SetHandle(handle).SetLimit(1)

	list, err := store.GroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) GroupFindByID(ctx context.Context, id string) (group GroupInterface, err error) {
	if id == "" {
		return nil, errors.New("group id is empty")
	}

	query := NewGroupQuery().SetID(id).SetLimit(1)

	list, err := store.GroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) GroupList(ctx context.Context, query GroupQueryInterface) ([]GroupInterface, error) {
	if query == nil {
		return []GroupInterface{}, errors.New("at group list > group query is nil")
	}

	q, columns, err := store.groupSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []GroupInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []GroupInterface{}, errors.New("groupstore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []GroupInterface{}, err
	}

	list := []GroupInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewGroupFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) GroupSoftDelete(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("at group soft delete > group is nil")
	}

	group.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.GroupUpdate(ctx, group)
}

func (store *store) GroupSoftDeleteByID(ctx context.Context, id string) error {
	group, err := store.GroupFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.GroupSoftDelete(ctx, group)
}

func (store *store) GroupUpdate(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("at group update > group is nil")
	}

	group.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := group.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.groupTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(group.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("groupstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	group.MarkAsNotDirty()

	return err
}

func (store *store) groupSelectQuery(options GroupQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("group options is nil")
	}

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.groupTableName)

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasStatus() {
		q = q.Where(goqu.C(COLUMN_STATUS).Eq(options.Status()))
	}

	if options.HasStatusIn() {
		q = q.Where(goqu.C(COLUMN_STATUS).In(options.StatusIn()))
	}

	if options.HasHandle() {
		q = q.Where(goqu.C(COLUMN_HANDLE).Eq(options.Handle()))
	}

	if options.HasTitleLike() {
		q = q.Where(goqu.C(COLUMN_TITLE).ILike(`%` + options.TitleLike() + `%`))
	}

	if options.HasCreatedAtGte() && options.HasCreatedAtLte() {
		q = q.Where(
			goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()),
			goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()),
		)
	} else if options.HasCreatedAtGte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()))
	} else if options.HasCreatedAtLte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()))
	}

	if !options.IsCountOnly() {
		if options.HasLimit() {
			q = q.Limit(cast.ToUint(options.Limit()))
		}

		if options.HasOffset() {
			q = q.Offset(cast.ToUint(options.Offset()))
		}
	}

	if options.HasOrderBy() {
		sort := lo.Ternary(options.HasSortDirection(), options.SortDirection(), sb.DESC)
		if strings.EqualFold(sort, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy()).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy()).Desc())
		}
	}

	columns = []any{}

	for _, column := range options.Columns() {
		columns = append(columns, column)
	}

	if options.SoftDeletedIncluded() {
		return q, columns, nil // soft deleted groups requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
