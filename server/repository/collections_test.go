package repository

import (
	"testing"
)

func TestUsersCollectionName(t *testing.T) {
	var expected string = "users"
	if CollectionUsers != expected {
		t.Errorf("user collection got %v want %v",
			CollectionUsers, expected)
	}
}

func TestSessionsCollectionName(t *testing.T) {
	var expected string = "sessions"
	if CollectionSessions != expected {
		t.Errorf("sessions collection got %v want %v",
			CollectionSessions, expected)
	}
}

func TestProfilesCollectionName(t *testing.T) {
	var expected string = "profiles"
	if CollectionProfiles != expected {
		t.Errorf("profiles collection got %v want %v",
			CollectionProfiles, expected)
	}
}
