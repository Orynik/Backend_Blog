package teststore_test

import (
	"testing"

	"github.com/Orynik/Backend_Blog/internal/app/model"
	"github.com/Orynik/Backend_Blog/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestCommentRepository_Create(t *testing.T) {
	s := teststore.New()
	c := model.TestComment(t)
	err := s.Comment().Create(c)
	assert.NoError(t, err)
}
func TestCommentRepository_GetComments(t *testing.T) {
	s := teststore.New()
	c := model.TestComment(t)
	_ = s.Comment().Create(c)

	c = model.TestComment(t)
	_ = s.Comment().Create(c)

	result, err := s.Comment().GetComments()
	assert.NoError(t, err)

	assert.Equal(t, result[0].ID, 1)
	assert.Equal(t, result[0].Author, "qwe")
	assert.Equal(t, result[0].Text, "qwe")
}
