package users

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestService_GetUserSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := NewMockDatabase(ctrl)
	userIds := []string{"1", "2"}
	db.EXPECT().GetUserSummary(userIds).Return([]*User{
		{
			ID:     "1",
			Name:   "Jazmine_1",
			Avatar: "",
		},
		{
			ID:     "2",
			Name:   "Jazmine_2",
			Avatar: "",
		},
	})

	svc := New(db)
	users := svc.GetUserSummary(userIds)
	assert.Equal(t, userIds[0], users[0].ID)
	assert.Equal(t, userIds[1], users[1].ID)
}
