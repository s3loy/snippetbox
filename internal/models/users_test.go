package models

import (
	"errors"
	"testing"
	"time"

	"snippetbox.s3loy.tech/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skiping integration test")
	}

	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)
			m := UserModel{db}
			exists, err := m.Exists(tt.userID)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}

func TestUserModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	t.Run("Valid ID", func(t *testing.T) {
		db := newTestDB(t)
		m := UserModel{db}

		user, err := m.Get(1)
		assert.NilError(t, err)
		assert.Equal(t, user.ID, 1)
		assert.Equal(t, user.Name, "Alice Jones")
		assert.Equal(t, user.Email, "alice@example.com")
		assert.Equal(t, user.Created, time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC))
	})

	t.Run("Non-existent ID", func(t *testing.T) {
		db := newTestDB(t)
		m := UserModel{db}

		_, err := m.Get(999)
		if !errors.Is(err, ErrNoRecord) {
			t.Errorf("got: %v; want: %v", err, ErrNoRecord)
		}
	})
}
