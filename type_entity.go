package groupstore

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
)

// == CLASS ===================================================================

type entityGroup struct {
	dataobject.DataObject
}

var _ EntityGroupInterface = (*entityGroup)(nil)

// == CONSTRUCTORS ============================================================

func NewEntityGroup() EntityGroupInterface {
	o := (&entityGroup{}).
		SetID(uid.HumanUid()).
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	err := o.SetMetas(map[string]string{})

	if err != nil {
		return o
	}

	return o
}

func NewEntityGroupFromExistingData(data map[string]string) EntityGroupInterface {
	o := &entityGroup{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func (o *entityGroup) IsSoftDeleted() bool {
	return o.SoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

// == SETTERS AND GETTERS =====================================================

func (o *entityGroup) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *entityGroup) CreatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *entityGroup) SetCreatedAt(createdAt string) EntityGroupInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *entityGroup) EntityType() string {
	return o.Get(COLUMN_ENTITY_TYPE)
}

func (o *entityGroup) SetEntityType(entityType string) EntityGroupInterface {
	o.Set(COLUMN_ENTITY_TYPE, entityType)
	return o
}

func (o *entityGroup) EntityID() string {
	return o.Get(COLUMN_ENTITY_ID)
}

func (o *entityGroup) SetEntityID(entityID string) EntityGroupInterface {
	o.Set(COLUMN_ENTITY_ID, entityID)
	return o
}

func (o *entityGroup) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *entityGroup) SetID(id string) EntityGroupInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *entityGroup) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *entityGroup) SetMemo(memo string) EntityGroupInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *entityGroup) Metas() (map[string]string, error) {
	metasStr := o.Get(COLUMN_METAS)

	if metasStr == "" {
		metasStr = "{}"
	}

	metasJson, errJson := utils.FromJSON(metasStr, map[string]string{})
	if errJson != nil {
		return map[string]string{}, errJson
	}

	return maputils.MapStringAnyToMapStringString(metasJson.(map[string]any)), nil
}

func (o *entityGroup) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *entityGroup) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *entityGroup) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, mapString)
	return nil
}

func (o *entityGroup) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *entityGroup) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *entityGroup) SoftDeletedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *entityGroup) SetSoftDeletedAt(deletedAt string) EntityGroupInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *entityGroup) GroupID() string {
	return o.Get(COLUMN_GROUP_ID)
}

func (o *entityGroup) SetGroupID(roleID string) EntityGroupInterface {
	o.Set(COLUMN_GROUP_ID, roleID)
	return o
}

func (o *entityGroup) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *entityGroup) UpdatedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *entityGroup) SetUpdatedAt(updatedAt string) EntityGroupInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
