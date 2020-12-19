package graphql

import (
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	gql := Initialize()

	var expected string = "*services.Service"

	if reflect.TypeOf(gql).String() != expected {
		t.Error("Should be a pointer of Service interface")
	}
}

func TestHandler(t *testing.T) {
	Initialize()

	handler := Handler()

	var expected string = "*handler.Handler"

	if reflect.TypeOf(handler).String() != expected {
		t.Error("Should be a http handler")
	}
}
