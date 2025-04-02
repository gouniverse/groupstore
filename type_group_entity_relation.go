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

type groupEntityRelation struct {
	dataobject.DataObject
}

var _ GroupEntityRelationInterface = (*groupEntityRelation)(nil)

// == CONSTRUCTORS ============================================================

func NewGroupEntityRelation() GroupEntityRelationInterface {
	o := (&groupEntityRelation{}).
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

func NewGroupEntityRelationFromExistingData(data map[string]string) GroupEntityRelationInterface {
	o := &groupEntityRelation{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func (o *groupEntityRelation) IsSoftDeleted() bool {
	return o.SoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

// == SETTERS AND GETTERS =====================================================

func (o *groupEntityRelation) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *groupEntityRelation) CreatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *groupEntityRelation) SetCreatedAt(createdAt string) GroupEntityRelationInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *groupEntityRelation) EntityType() string {
	return o.Get(COLUMN_ENTITY_TYPE)
}

func (o *groupEntityRelation) SetEntityType(entityType string) GroupEntityRelationInterface {
	o.Set(COLUMN_ENTITY_TYPE, entityType)
	return o
}

func (o *groupEntityRelation) EntityID() string {
	return o.Get(COLUMN_ENTITY_ID)
}

func (o *groupEntityRelation) SetEntityID(entityID string) GroupEntityRelationInterface {
	o.Set(COLUMN_ENTITY_ID, entityID)
	return o
}

func (o *groupEntityRelation) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *groupEntityRelation) SetID(id string) GroupEntityRelationInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *groupEntityRelation) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *groupEntityRelation) SetMemo(memo string) GroupEntityRelationInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *groupEntityRelation) Metas() (map[string]string, error) {
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

func (o *groupEntityRelation) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *groupEntityRelation) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *groupEntityRelation) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, mapString)
	return nil
}

func (o *groupEntityRelation) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *groupEntityRelation) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *groupEntityRelation) SoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *groupEntityRelation) SetSoftDeletedAt(deletedAt string) GroupEntityRelationInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *groupEntityRelation) GroupID() string {
	return o.Get(COLUMN_GROUP_ID)
}

func (o *groupEntityRelation) SetGroupID(roleID string) GroupEntityRelationInterface {
	o.Set(COLUMN_GROUP_ID, roleID)
	return o
}

func (o *groupEntityRelation) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *groupEntityRelation) UpdatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *groupEntityRelation) SetUpdatedAt(updatedAt string) GroupEntityRelationInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
