package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {

	email := "j@j.com"
	name := "Wesley"

	user, err := model.NewUser(email, name)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Email, email)
	require.Equal(t, user.Name, name)

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}
