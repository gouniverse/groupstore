package groupstore

import "errors"

type GroupQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) GroupQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) GroupQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) GroupQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) GroupQueryInterface

	HasHandle() bool
	Handle() string
	SetHandle(handle string) GroupQueryInterface

	HasID() bool
	ID() string
	SetID(id string) GroupQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) GroupQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) GroupQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) GroupQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) GroupQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) GroupQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) GroupQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) GroupQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) GroupQueryInterface

	HasTitleLike() bool
	TitleLike() string
	SetTitleLike(titleLike string) GroupQueryInterface

	hasProperty(name string) bool
}

func NewGroupQuery() GroupQueryInterface {
	return &groupQueryImplementation{
		properties: make(map[string]any),
	}
}

type groupQueryImplementation struct {
	properties map[string]any
}

func (c *groupQueryImplementation) Validate() error {
	if c.HasID() && c.ID() == "" {
		return errors.New("group query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("group query. id_in cannot be empty")
	}

	if c.HasStatus() && c.Status() == "" {
		return errors.New("group query. status cannot be empty")
	}

	if c.HasTitleLike() && c.TitleLike() == "" {
		return errors.New("group query. title_like cannot be empty")
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

func (c *groupQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *groupQueryImplementation) SetColumns(columns []string) GroupQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *groupQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *groupQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *groupQueryImplementation) SetCountOnly(countOnly bool) GroupQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *groupQueryImplementation) HasCreatedAtGte() bool {
	return c.hasProperty("created_at_gte")
}

func (c *groupQueryImplementation) CreatedAtGte() string {
	if !c.HasCreatedAtGte() {
		return ""
	}

	return c.properties["created_at_gte"].(string)
}

func (c *groupQueryImplementation) SetCreatedAtGte(createdAtGte string) GroupQueryInterface {
	c.properties["created_at_gte"] = createdAtGte

	return c
}

func (c *groupQueryImplementation) HasCreatedAtLte() bool {
	return c.hasProperty("created_at_lte")
}

func (c *groupQueryImplementation) CreatedAtLte() string {
	if !c.HasCreatedAtLte() {
		return ""
	}

	return c.properties["created_at_lte"].(string)
}

func (c *groupQueryImplementation) SetCreatedAtLte(createdAtLte string) GroupQueryInterface {
	c.properties["created_at_lte"] = createdAtLte

	return c
}

func (c *groupQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *groupQueryImplementation) HasHandle() bool {
	return c.hasProperty("handle")
}

func (c *groupQueryImplementation) Handle() string {
	if !c.HasHandle() {
		return ""
	}

	return c.properties["handle"].(string)
}

func (c *groupQueryImplementation) SetHandle(handle string) GroupQueryInterface {
	c.properties["handle"] = handle

	return c
}

func (c *groupQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *groupQueryImplementation) SetID(id string) GroupQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *groupQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *groupQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *groupQueryImplementation) SetIDIn(idIn []string) GroupQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *groupQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *groupQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *groupQueryImplementation) SetLimit(limit int) GroupQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *groupQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *groupQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *groupQueryImplementation) SetOffset(offset int) GroupQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *groupQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *groupQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *groupQueryImplementation) SetOrderBy(orderBy string) GroupQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *groupQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *groupQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *groupQueryImplementation) SetSortDirection(sortDirection string) GroupQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *groupQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *groupQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *groupQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) GroupQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *groupQueryImplementation) HasStatus() bool {
	return c.hasProperty("status")
}

func (c *groupQueryImplementation) Status() string {
	if !c.HasStatus() {
		return ""
	}

	return c.properties["status"].(string)
}

func (c *groupQueryImplementation) SetStatus(status string) GroupQueryInterface {
	c.properties["status"] = status

	return c
}

func (c *groupQueryImplementation) HasStatusIn() bool {
	return c.hasProperty("status_in")
}

func (c *groupQueryImplementation) StatusIn() []string {
	if !c.HasStatusIn() {
		return []string{}
	}

	return c.properties["status_in"].([]string)
}

func (c *groupQueryImplementation) SetStatusIn(statusIn []string) GroupQueryInterface {
	c.properties["status_in"] = statusIn

	return c
}

func (c *groupQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *groupQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *groupQueryImplementation) SetTitleLike(titleLike string) GroupQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *groupQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
