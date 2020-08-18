package teststore_test

import (
	"testing"

	"github.com/Orynik/Backend_Blog/internal/app/model"
	"github.com/Orynik/Backend_Blog/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesRepository_Create(t *testing.T) {
	s := teststore.New()
	a := model.TestCategory(t)
	err := s.Category().Create(a)
	assert.NoError(t, err)
}

func TestCategoryRepository_GetCategory(t *testing.T) {
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
