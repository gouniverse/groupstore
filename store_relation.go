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

func (store *store) RelationCount(ctx context.Context, options RelationQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.relationSelectQuery(options)

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

func (store *store) RelationCreate(ctx context.Context, relation RelationInterface) error {
	if relation == nil {
		return errors.New("groupstore > RelationCreate. relation is nil")
	}

	if relation.GroupID() == "" {
		return errors.New("groupstore > RelationCreate. relation groupID is empty")
	}

	if relation.EntityID() == "" {
		return errors.New("groupstore > RelationCreate. relation entityID is empty")
	}

	if relation.EntityType() == "" {
		return errors.New("groupstore > RelationCreate. relation entityType is empty")
	}

	relationExists, err := store.RelationFindByEntityAndGroup(
		ctx,
		relation.EntityType(),
		relation.EntityID(),
		relation.GroupID(),
	)

	if err != nil {
		return err
	}

	if relationExists != nil {
		return errors.New("groupstore > RelationCreate. relation with the same entityType-entityID-groupID combination already exists")
	}

	relation.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	relation.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := relation.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.groupEntityRelationTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("insert", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityGroupstore: database is nil")
	}

	_, err = database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	if err != nil {
		return err
	}

	relation.MarkAsNotDirty()

	return nil
}

func (store *store) RelationDelete(ctx context.Context, relation RelationInterface) error {
	if relation == nil {
		return errors.New("relation is nil")
	}

	return store.RelationDeleteByID(ctx, relation.ID())
}

func (store *store) RelationDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("relation id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.groupEntityRelationTableName).
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

func (store *store) RelationFindByEntityAndGroup(
	ctx context.Context,
	entityType string,
	entityID string,
	groupID string,
) (relation RelationInterface, err error) {
	if entityType == "" {
		return nil, errors.New("relation findBy entity and group > entityType is empty")
	}

	if entityID == "" {
		return nil, errors.New("relation findBy entity and group > entityID is empty")
	}

	if groupID == "" {
		return nil, errors.New("relation findBy entity and group > groupID is empty")
	}

	query := NewRelationQuery().
		SetEntityType(entityType).
		SetEntityID(entityID).
		SetGroupID(groupID).
		SetLimit(1)

	list, err := store.RelationList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) RelationFindByID(ctx context.Context, id string) (relation RelationInterface, err error) {
	if id == "" {
		return nil, errors.New("relation id is empty")
	}

	query := NewRelationQuery().SetID(id).SetLimit(1)

	list, err := store.RelationList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) RelationList(ctx context.Context, query RelationQueryInterface) ([]RelationInterface, error) {
	if query == nil {
		return []RelationInterface{}, errors.New("at relation list > relation query is nil")
	}

	q, columns, err := store.relationSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []RelationInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []RelationInterface{}, errors.New("entityGroupstore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []RelationInterface{}, err
	}

	list := []RelationInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewGroupEntityRelationFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) RelationSoftDelete(ctx context.Context, relation RelationInterface) error {
	if relation == nil {
		return errors.New("at relation soft delete > relation is nil")
	}

	relation.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.RelationUpdate(ctx, relation)
}

func (store *store) RelationSoftDeleteByID(ctx context.Context, id string) error {
	relation, err := store.RelationFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.RelationSoftDelete(ctx, relation)
}

func (store *store) RelationUpdate(ctx context.Context, relation RelationInterface) error {
	if relation == nil {
		return errors.New("at relation update > relation is nil")
	}

	relation.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := relation.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.groupEntityRelationTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(relation.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityGroupstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	relation.MarkAsNotDirty()

	return err
}

func (store *store) relationSelectQuery(options RelationQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("relation options is nil")
	}

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.groupEntityRelationTableName)

	if options.HasEntityID() {
		q = q.Where(goqu.C(COLUMN_ENTITY_ID).Eq(options.EntityID()))
	}

	if options.HasEntityType() {
		q = q.Where(goqu.C(COLUMN_ENTITY_TYPE).Eq(options.EntityType()))
	}

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasGroupID() {
		q = q.Where(goqu.C(COLUMN_GROUP_ID).Eq(options.GroupID()))
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
		return q, columns, nil // soft deleted entityGroups requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
