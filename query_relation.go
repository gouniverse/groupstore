package groupstore

import "errors"

type RelationQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) RelationQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) RelationQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) RelationQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) RelationQueryInterface

	HasEntityID() bool
	EntityID() string
	SetEntityID(entityID string) RelationQueryInterface

	HasEntityType() bool
	EntityType() string
	SetEntityType(entityType string) RelationQueryInterface

	HasID() bool
	ID() string
	SetID(id string) RelationQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) RelationQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) RelationQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) RelationQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) RelationQueryInterface

	HasGroupID() bool
	GroupID() string
	SetGroupID(groupID string) RelationQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) RelationQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) RelationQueryInterface

	hasProperty(name string) bool
}

func NewRelationQuery() RelationQueryInterface {
	return &groupEntityQueryImplementation{
		properties: make(map[string]any),
	}
}

type groupEntityQueryImplementation struct {
	properties map[string]any
}

func (c *groupEntityQueryImplementation) Validate() error {
	if c.HasCreatedAtGte() && c.CreatedAtGte() == "" {
		return errors.New("group query. created_at_gte cannot be empty")
	}

	if c.HasCreatedAtLte() && c.CreatedAtLte() == "" {
		return errors.New("group query. created_at_lte cannot be empty")
	}

	if c.HasEntityID() && c.EntityID() == "" {
		return errors.New("group query. entity_id cannot be empty")
	}

	if c.HasEntityType() && c.EntityType() == "" {
		return errors.New("group query. entity_type cannot be empty")
	}

	if c.HasID() && c.ID() == "" {
		return errors.New("group query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("group query. id_in cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("group query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("group query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("group query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("group query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *groupEntityQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *groupEntityQueryImplementation) SetColumns(columns []string) RelationQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *groupEntityQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *groupEntityQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *groupEntityQueryImplementation) SetCountOnly(countOnly bool) RelationQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *groupEntityQueryImplementation) HasCreatedAtGte() bool {
	return c.hasProperty("created_at_gte")
}

func (c *groupEntityQueryImplementation) CreatedAtGte() string {
	if !c.HasCreatedAtGte() {
		return ""
	}

	return c.properties["created_at_gte"].(string)
}

func (c *groupEntityQueryImplementation) SetCreatedAtGte(createdAtGte string) RelationQueryInterface {
	c.properties["created_at_gte"] = createdAtGte

	return c
}

func (c *groupEntityQueryImplementation) HasCreatedAtLte() bool {
	return c.hasProperty("created_at_lte")
}

func (c *groupEntityQueryImplementation) CreatedAtLte() string {
	if !c.HasCreatedAtLte() {
		return ""
	}

	return c.properties["created_at_lte"].(string)
}

func (c *groupEntityQueryImplementation) SetCreatedAtLte(createdAtLte string) RelationQueryInterface {
	c.properties["created_at_lte"] = createdAtLte

	return c
}

func (c *groupEntityQueryImplementation) HasEntityType() bool {
	return c.hasProperty("entity_type")
}

func (c *groupEntityQueryImplementation) EntityType() string {
	if !c.HasEntityType() {
		return ""
	}

	return c.properties["entity_type"].(string)
}

func (c *groupEntityQueryImplementation) SetEntityType(entityType string) RelationQueryInterface {
	c.properties["entity_type"] = entityType

	return c
}

func (c *groupEntityQueryImplementation) HasEntityID() bool {
	return c.hasProperty("entity_id")
}

func (c *groupEntityQueryImplementation) EntityID() string {
	if !c.HasEntityID() {
		return ""
	}

	return c.properties["entity_id"].(string)
}

func (c *groupEntityQueryImplementation) SetEntityID(entityID string) RelationQueryInterface {
	c.properties["entity_id"] = entityID

	return c
}

func (c *groupEntityQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *groupEntityQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *groupEntityQueryImplementation) SetID(id string) RelationQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *groupEntityQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *groupEntityQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *groupEntityQueryImplementation) SetIDIn(idIn []string) RelationQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *groupEntityQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *groupEntityQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *groupEntityQueryImplementation) SetLimit(limit int) RelationQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *groupEntityQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *groupEntityQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *groupEntityQueryImplementation) SetOffset(offset int) RelationQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *groupEntityQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *groupEntityQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *groupEntityQueryImplementation) SetOrderBy(orderBy string) RelationQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *groupEntityQueryImplementation) HasGroupID() bool {
	return c.hasProperty("group_id")
}

func (c *groupEntityQueryImplementation) GroupID() string {
	if !c.HasGroupID() {
		return ""
	}

	return c.properties["group_id"].(string)
}

func (c *groupEntityQueryImplementation) SetGroupID(groupID string) RelationQueryInterface {
	c.properties["group_id"] = groupID

	return c
}

func (c *groupEntityQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *groupEntityQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *groupEntityQueryImplementation) SetSortDirection(sortDirection string) RelationQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *groupEntityQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *groupEntityQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *groupEntityQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) RelationQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *groupEntityQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *groupEntityQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *groupEntityQueryImplementation) SetTitleLike(titleLike string) RelationQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *groupEntityQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
