package model_test

import (
	"testing"

	"github.com/Orynik/Backend_Blog/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	// u := model.TestUser(t)
	// assert.NoError(t, u.Validate())
	//Табличный тест...
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			//Валидный кейс ...
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			//Если у Usera есть EncryptedPassword в этом сдучае он будет тоже волиден ...
			name: "with encrypt password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"

				return u
			},
			isValid: true,
		},
		{
			//Валидация на Email ...
			name: "emty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			//Валидация на пустой Email ...
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			//Валидация на Password...
			name: "emty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			//Валидация на пустой Password ...
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
