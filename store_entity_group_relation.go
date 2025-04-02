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

func (store *store) EntityGroupCount(ctx context.Context, options EntityGroupQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.entityGroupSelectQuery(options)

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

func (store *store) EntityGroupCreate(ctx context.Context, entityGroup GroupEntityRelationInterface) error {
	if entityGroup == nil {
		return errors.New("groupstore > EntityGroupCreate. entityGroup is nil")
	}

	if entityGroup.GroupID() == "" {
		return errors.New("groupstore > EntityGroupCreate. entityGroup groupID is empty")
	}

	if entityGroup.EntityID() == "" {
		return errors.New("groupstore > EntityGroupCreate. entityGroup entityID is empty")
	}

	if entityGroup.EntityType() == "" {
		return errors.New("groupstore > EntityGroupCreate. entityGroup entityType is empty")
	}

	entityGroupExists, err := store.EntityGroupFindByEntityAndGroup(
		ctx,
		entityGroup.EntityType(),
		entityGroup.EntityID(),
		entityGroup.GroupID(),
	)

	if err != nil {
		return err
	}

	if entityGroupExists != nil {
		return errors.New("groupstore > EntityGroupCreate. entityGroup with the same entityType-entityID-groupID combination already exists")
	}

	entityGroup.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	entityGroup.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := entityGroup.Data()

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

	entityGroup.MarkAsNotDirty()

	return nil
}

func (store *store) EntityGroupDelete(ctx context.Context, entityGroup GroupEntityRelationInterface) error {
	if entityGroup == nil {
		return errors.New("entityGroup is nil")
	}

	return store.EntityGroupDeleteByID(ctx, entityGroup.ID())
}

func (store *store) EntityGroupDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("entityGroup id is empty")
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

func (store *store) EntityGroupFindByEntityAndGroup(
	ctx context.Context,
	entityType string,
	entityID string,
	groupID string,
) (entityGroup GroupEntityRelationInterface, err error) {
	if entityType == "" {
		return nil, errors.New("EntityGroupFindByEntityAndGroup entityType is empty")
	}

	if entityID == "" {
		return nil, errors.New("EntityGroupFindByEntityAndGroup entityID is empty")
	}

	if groupID == "" {
		return nil, errors.New("EntityGroupFindByEntityAndGroup groupID is empty")
	}

	query := NewEntityGroupQuery().
		SetEntityType(entityType).
		SetEntityID(entityID).
		SetGroupID(groupID).
		SetLimit(1)

	list, err := store.EntityGroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EntityGroupFindByID(ctx context.Context, id string) (entityGroup GroupEntityRelationInterface, err error) {
	if id == "" {
		return nil, errors.New("entityGroup id is empty")
	}

	query := NewEntityGroupQuery().SetID(id).SetLimit(1)

	list, err := store.EntityGroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EntityGroupList(ctx context.Context, query EntityGroupQueryInterface) ([]GroupEntityRelationInterface, error) {
	if query == nil {
		return []GroupEntityRelationInterface{}, errors.New("at entityGroup list > entityGroup query is nil")
	}

	q, columns, err := store.entityGroupSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []GroupEntityRelationInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []GroupEntityRelationInterface{}, errors.New("entityGroupstore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []GroupEntityRelationInterface{}, err
	}

	list := []GroupEntityRelationInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewGroupEntityRelationFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) EntityGroupSoftDelete(ctx context.Context, entityGroup GroupEntityRelationInterface) error {
	if entityGroup == nil {
		return errors.New("at entityGroup soft delete > entityGroup is nil")
	}

	entityGroup.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.EntityGroupUpdate(ctx, entityGroup)
}

func (store *store) EntityGroupSoftDeleteByID(ctx context.Context, id string) error {
	entityGroup, err := store.EntityGroupFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.EntityGroupSoftDelete(ctx, entityGroup)
}

func (store *store) EntityGroupUpdate(ctx context.Context, entityGroup GroupEntityRelationInterface) error {
	if entityGroup == nil {
		return errors.New("at entityGroup update > entityGroup is nil")
	}

	entityGroup.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := entityGroup.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.groupEntityRelationTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(entityGroup.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityGroupstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	entityGroup.MarkAsNotDirty()

	return err
}

func (store *store) entityGroupSelectQuery(options EntityGroupQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("entityGroup options is nil")
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
