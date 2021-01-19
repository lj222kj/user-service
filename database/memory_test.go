package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemory_getUserSummary(t *testing.T) {
	user, err := db.getUserSummary("1")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Jazmine_1", user.Name)
}

func TestMemory_GetUserSummary(t *testing.T) {
	user := db.GetUserSummary([]string{"1", "2"})
	assert.NotNil(t, user)
	assert.Len(t, user, 2)
}
