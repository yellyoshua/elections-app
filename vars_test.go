package main

import (
	"testing"
)

func TestUsersCollectionName(t *testing.T) {
	var expected string = "users"
	if UserCollectionName != expected {
		t.Errorf("user collection got %v want %v",
			UserCollectionName, expected)
	}
}

func TestSessionsCollectionName(t *testing.T) {
	var expected string = "sessions"
	if SessionCollectionName != expected {
		t.Errorf("sessions collection got %v want %v",
			SessionCollectionName, expected)
	}
}

func TestProfilesCollectionName(t *testing.T) {
	var expected string = "profiles"
	if ProfileCollectionName != expected {
		t.Errorf("profiles collection got %v want %v",
			ProfileCollectionName, expected)
	}
}
