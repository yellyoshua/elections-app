package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersCollectionName(t *testing.T) {
	var expected string = "users"

	assert.Equal(t, expected, CollectionUsers)
}

func TestSessionsCollectionName(t *testing.T) {
	var expected string = "sessions"

	assert.Equal(t, expected, CollectionSessions)
}

func TestProfilesCollectionName(t *testing.T) {
	var expected string = "profiles"

	assert.Equal(t, expected, CollectionProfiles)
}
