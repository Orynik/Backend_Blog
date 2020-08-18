package sqlstore_test

import (
	"testing"

	"github.com/Kentabr5427/http-rest-api/internal/app/model"
	"github.com/Kentabr5427/http-rest-api/internal/app/store/sqlstore"
	"github.com/Kentabr5427/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestCommentsRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("comments")

	s := sqlstore.New(db)
	a := model.TestComment(t)

	assert.NoError(t, s.Comment().Create(a))
	assert.NotNil(t, a)
	assert.Equal(t, a.ID, 1)
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
