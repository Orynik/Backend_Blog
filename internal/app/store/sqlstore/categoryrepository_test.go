package sqlstore_test

import (
	"testing"

	"github.com/Orynik/Backend_Blog/internal/app/model"
	"github.com/Orynik/Backend_Blog/internal/app/store/sqlstore"
	"github.com/Orynik/Backend_Blog/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("category")

	s := sqlstore.New(db)
	a := model.TestCategory(t)

	assert.NoError(t, s.Category().Create(a))
	assert.NotNil(t, a)
	assert.Equal(t, a.ID, 1)
}

func TestCategoryRepository_GetComments(t *testing.T) {
	s := teststore.New()
	c := model.TestCategory(t)
	_ = s.Category().Create(c)

	c = model.TestCategory(t)
	_ = s.Category().Create(c)

	result, err := s.Category().GetCategories()
	assert.NoError(t, err)

	assert.Equal(t, result[0].ID, 1)
	assert.Equal(t, result[0].Category, "Backend")
}
