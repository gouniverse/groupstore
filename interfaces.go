package groupstore

import (
	"context"
	"database/sql"

	"github.com/dromara/carbon/v2"
)

type StoreInterface interface {
	// AutoMigrate auto migrates the database schema
	AutoMigrate() error

	// EnableDebug enables or disables the debug mode
	EnableDebug(debug bool)

	// DB returns the underlying database connection
	DB() *sql.DB

	// == Group Methods =======================================================//

	// GroupCount returns the number of groups based on the given query options
	GroupCount(ctx context.Context, options GroupQueryInterface) (int64, error)

	// GroupCreate creates a new group
	GroupCreate(ctx context.Context, group GroupInterface) error

	// GroupDelete deletes a group
	GroupDelete(ctx context.Context, group GroupInterface) error

	// GroupDeleteByID deletes a group by its ID
	GroupDeleteByID(ctx context.Context, id string) error

	// GroupFindByHandle returns a group by its handle
	GroupFindByHandle(ctx context.Context, handle string) (GroupInterface, error)

	// GroupFindByID returns a group by its ID
	GroupFindByID(ctx context.Context, id string) (GroupInterface, error)

	// GroupList returns a list of groups based on the given query options
	GroupList(ctx context.Context, query GroupQueryInterface) ([]GroupInterface, error)

	// GroupSoftDelete soft deletes a group
	GroupSoftDelete(ctx context.Context, group GroupInterface) error

	// GroupSoftDeleteByID soft deletes a group by its ID
	GroupSoftDeleteByID(ctx context.Context, id string) error

	// GroupUpdate updates a group
	GroupUpdate(ctx context.Context, group GroupInterface) error

	// == EntityGroup Methods =================================================//

	// EntityGroupCount returns the number of group entities mappings based on the given query options
	EntityGroupCount(ctx context.Context, options EntityGroupQueryInterface) (int64, error)

	// EntityGroupCreate creates a new group entity mapping
	EntityGroupCreate(ctx context.Context, entityGroup EntityGroupInterface) error

	// EntityGroupDelete deletes a group entity mapping
	EntityGroupDelete(ctx context.Context, entityGroup EntityGroupInterface) error

	// EntityGroupDeleteByID deletes a group entity mapping by its ID
	EntityGroupDeleteByID(ctx context.Context, id string) error

	// EntityGroupFindByEntityAndGroup returns a group entity mapping by its entity type, entity ID and group ID
	EntityGroupFindByEntityAndGroup(ctx context.Context, entityType string, entityID string, groupID string) (EntityGroupInterface, error)

	// EntityGroupFindByID returns a group entity mapping by its ID
	EntityGroupFindByID(ctx context.Context, id string) (EntityGroupInterface, error)

	// EntityGroupList returns a list of group entity mappings based on the given query options
	EntityGroupList(ctx context.Context, query EntityGroupQueryInterface) ([]EntityGroupInterface, error)

	// EntityGroupSoftDelete soft deletes a group entity mapping
	EntityGroupSoftDelete(ctx context.Context, entityGroup EntityGroupInterface) error

	// EntityGroupSoftDeleteByID soft deletes a group entity mapping by its ID
	EntityGroupSoftDeleteByID(ctx context.Context, id string) error

	// EntityGroupUpdate updates a group entity mapping
	EntityGroupUpdate(ctx context.Context, entityGroup EntityGroupInterface) error
}

type GroupInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) GroupInterface

	Handle() string
	SetHandle(handle string) GroupInterface

	ID() string
	SetID(id string) GroupInterface

	Memo() string
	SetMemo(memo string) GroupInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	Status() string
	SetStatus(status string) GroupInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) GroupInterface

	Title() string
	SetTitle(title string) GroupInterface

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) GroupInterface
}

type EntityGroupInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) EntityGroupInterface

	EntityType() string
	SetEntityType(entityType string) EntityGroupInterface

	EntityID() string
	SetEntityID(entityID string) EntityGroupInterface

	ID() string
	SetID(id string) EntityGroupInterface

	Memo() string
	SetMemo(memo string) EntityGroupInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	GroupID() string
	SetGroupID(groupID string) EntityGroupInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) EntityGroupInterface

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) EntityGroupInterface
}

type UserInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	Get(columnName string) string
	Set(columnName string, value string)

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool
	IsUnverified() bool

	IsAdministrator() bool
	IsManager() bool
	IsSuperuser() bool

	IsRegistrationCompleted() bool

	// setters and getters

	BusinessName() string
	SetBusinessName(businessName string) UserInterface

	Country() string
	SetCountry(country string) UserInterface

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) UserInterface

	Email() string
	SetEmail(email string) UserInterface

	ID() string
	SetID(id string) UserInterface

	FirstName() string
	SetFirstName(firstName string) UserInterface

	LastName() string
	SetLastName(lastName string) UserInterface

	Memo() string
	SetMemo(memo string) UserInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error
	UpsertMetas(metas map[string]string) error

	MiddleNames() string
	SetMiddleNames(middleNames string) UserInterface

	Password() string
	SetPassword(password string) UserInterface

	Phone() string
	SetPhone(phone string) UserInterface

	ProfileImageUrl() string
	SetProfileImageUrl(profileImageUrl string) UserInterface

	Group() string
	SetGroup(group string) UserInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(deletedAt string) UserInterface

	Timezone() string
	SetTimezone(timezone string) UserInterface

	Status() string
	SetStatus(status string) UserInterface

	PasswordCompare(password string) bool

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}
