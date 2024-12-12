package groupstore

import (
	"context"
	"strings"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
)

func TestStoreEntityGroupCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.EntityGroupCount(context.Background(), NewEntityGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.EntityGroupCount(context.Background(), NewEntityGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	entityGroup2 := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_02").
		SetGroupID("PERMISSION_02")

	err = store.EntityGroupCreate(context.Background(), entityGroup2)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.EntityGroupCount(context.Background(), NewEntityGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreEntityGroupCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreEntityGroupCreate_Duplicate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err == nil {
		t.Fatal("must return error as duplicated entity to group relationship")
	}
}

func TestStoreEntityGroupDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityGroupDelete(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	entityGroupFound, err := store.EntityGroupFindByID(context.Background(), entityGroup.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityGroupFound != nil {
		t.Fatal("EntityGroup MUST be nil")
	}

	entityGroupFindWithDeleted, err := store.EntityGroupList(context.Background(), NewEntityGroupQuery().
		SetID(entityGroup.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityGroupFindWithDeleted) != 0 {
		t.Fatal("EntityGroup MUST be nil")
	}
}

func TestStoreEntityGroupDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityGroupDeleteByID(context.Background(), entityGroup.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	entityGroupFound, err := store.EntityGroupFindByID(context.Background(), entityGroup.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityGroupFound != nil {
		t.Fatal("EntityGroup MUST be nil")
	}

	entityGroupFindWithDeleted, err := store.EntityGroupList(context.Background(), NewEntityGroupQuery().
		SetID(entityGroup.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityGroupFindWithDeleted) != 0 {
		t.Fatal("EntityGroup MUST NOT be found")
	}
}

func TestStoreEntityGroupFindByEntityAndGroup(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = entityGroup.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityGroupCreate(database.Context(context.Background(), store.DB()), entityGroup)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	entityGroupFound, errFind := store.EntityGroupFindByEntityAndGroup(database.Context(context.Background(), store.DB()), entityGroup.EntityType(), entityGroup.EntityID(), entityGroup.GroupID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityGroupFound == nil {
		t.Fatal("EntityGroup MUST NOT be nil")
	}

	if entityGroupFound.ID() != entityGroup.ID() {
		t.Fatal("IDs do not match")
	}

	if entityGroupFound.EntityID() != entityGroup.EntityID() {
		t.Fatal("EntityIDs do not match")
	}

	if entityGroupFound.EntityType() != entityGroup.EntityType() {
		t.Fatal("EntityTypes do not match")
	}

	if entityGroupFound.GroupID() != entityGroup.GroupID() {
		t.Fatal("GroupIDs do not match")
	}

	if entityGroupFound.Meta("education_1") != entityGroup.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if entityGroupFound.Meta("education_2") != entityGroup.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if entityGroupFound.Meta("education_3") != entityGroup.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreEntityGroupFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = entityGroup.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ctx := database.Context(context.Background(), store.DB())
	err = store.EntityGroupCreate(ctx, entityGroup)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	entityGroupFound, errFind := store.EntityGroupFindByID(ctx, entityGroup.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityGroupFound == nil {
		t.Fatal("EntityGroup MUST NOT be nil")
	}

	if entityGroupFound.ID() != entityGroup.ID() {
		t.Fatal("IDs do not match")
	}

	if entityGroupFound.EntityID() != entityGroup.EntityID() {
		t.Fatal("EntityIDs do not match")
	}

	if entityGroupFound.EntityType() != entityGroup.EntityType() {
		t.Fatal("EntityTypes do not match")
	}

	if entityGroupFound.GroupID() != entityGroup.GroupID() {
		t.Fatal("GroupIDs do not match")
	}

	if entityGroupFound.Meta("education_1") != entityGroup.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if entityGroupFound.Meta("education_2") != entityGroup.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if entityGroupFound.Meta("education_3") != entityGroup.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreEntityGroupList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup1 := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	entityGroup2 := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_02").
		SetGroupID("PERMISSION_02")

	entityGroups := []EntityGroupInterface{
		entityGroup1,
		entityGroup2,
	}

	for _, entityGroup := range entityGroups {
		err = store.EntityGroupCreate(context.Background(), entityGroup)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	list1, err := store.EntityGroupList(context.Background(), NewEntityGroupQuery().SetGroupID("PERMISSION_01"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list1) != 1 {
		t.Fatal("unexpected list length:", len(list1))
	}

	list2, err := store.EntityGroupList(context.Background(), NewEntityGroupQuery().SetEntityType("USER").SetEntityID("USER_02"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list2) != 1 {
		t.Fatal("unexpected list length:", len(list2))
	}
}

func TestStoreEntityGroupSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityGroupSoftDelete(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityGroup.SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("EntityGroup MUST be soft deleted")
	}

	entityGroupFound, errFind := store.EntityGroupFindByID(context.Background(), entityGroup.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityGroupFound != nil {
		t.Fatal("EntityGroup MUST be soft deleted, so MUST be nil")
	}

	entityGroupFindWithDeleted, err := store.EntityGroupList(context.Background(), NewEntityGroupQuery().
		SetSoftDeletedIncluded(true).
		SetID(entityGroup.ID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityGroupFindWithDeleted) == 0 {
		t.Fatal("EntityGroup MUST be soft deleted")
	}

	if strings.Contains(entityGroupFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("EntityGroup MUST be soft deleted", entityGroup.SoftDeletedAt())
	}

	if !entityGroupFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("EntityGroup MUST be soft deleted")
	}
}

func TestStoreEntityGroupSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityGroup := NewEntityGroup().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetGroupID("PERMISSION_01")

	err = store.EntityGroupCreate(context.Background(), entityGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityGroupSoftDeleteByID(context.Background(), entityGroup.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityGroup.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("EntityGroup MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	entityGroupFound, errFind := store.EntityGroupFindByID(context.Background(), entityGroup.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityGroupFound != nil {
		t.Fatal("EntityGroup MUST be nil")
	}
	query := NewEntityGroupQuery().
		SetSoftDeletedIncluded(true).
		SetID(entityGroup.ID()).
		SetLimit(1)

	entityGroupFindWithDeleted, err := store.EntityGroupList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityGroupFindWithDeleted) == 0 {
		t.Fatal("EntityGroup MUST be soft deleted")
	}

	if strings.Contains(entityGroupFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("EntityGroup MUST be soft deleted", entityGroup.SoftDeletedAt())
	}

	if !entityGroupFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("EntityGroup MUST be soft deleted")
	}
}
