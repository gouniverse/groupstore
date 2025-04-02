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

	// == Relation Methods ====================================================//

	// RelationCount returns the number of group entities mappings based on the given query options
	RelationCount(ctx context.Context, options RelationQueryInterface) (int64, error)

	// RelationCreate creates a new group entity mapping
	RelationCreate(ctx context.Context, relation RelationInterface) error

	// RelationDelete deletes a group entity mapping
	RelationDelete(ctx context.Context, relation RelationInterface) error

	// RelationDeleteByID deletes a group entity mapping by its ID
	RelationDeleteByID(ctx context.Context, id string) error

	// RelationFindByEntityAndGroup returns a group entity mapping by its entity type, entity ID and group ID
	RelationFindByEntityAndGroup(ctx context.Context, entityType string, entityID string, groupID string) (RelationInterface, error)

	// RelationFindByID returns a group entity mapping by its ID
	RelationFindByID(ctx context.Context, id string) (RelationInterface, error)

	// RelationList returns a list of group entity mappings based on the given query options
	RelationList(ctx context.Context, query RelationQueryInterface) ([]RelationInterface, error)

	// RelationSoftDelete soft deletes a group entity mapping
	RelationSoftDelete(ctx context.Context, relation RelationInterface) error

	// RelationSoftDeleteByID soft deletes a group entity mapping by its ID
	RelationSoftDeleteByID(ctx context.Context, id string) error

	// RelationUpdate updates a group entity mapping
	RelationUpdate(ctx context.Context, relation RelationInterface) error
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
	CreatedAtCarbon() *carbon.Carbon
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
	SoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) GroupInterface

	Title() string
	SetTitle(title string) GroupInterface

	UpdatedAt() string
	UpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) GroupInterface
}

type RelationInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) RelationInterface

	EntityType() string
	SetEntityType(entityType string) RelationInterface

	EntityID() string
	SetEntityID(entityID string) RelationInterface

	ID() string
	SetID(id string) RelationInterface

	Memo() string
	SetMemo(memo string) RelationInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	GroupID() string
	SetGroupID(groupID string) RelationInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) RelationInterface

	UpdatedAt() string
	UpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) RelationInterface
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
	CreatedAtCarbon() *carbon.Carbon
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
	SoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(deletedAt string) UserInterface

	Timezone() string
	SetTimezone(timezone string) UserInterface

	Status() string
	SetStatus(status string) UserInterface

	PasswordCompare(password string) bool

	UpdatedAt() string
	UpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}
