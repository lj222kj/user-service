package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	db Database
)

func TestMain(m *testing.M) {
	db = New()
	os.Exit(m.Run())
}
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
