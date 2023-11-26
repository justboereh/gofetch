package gofetch_test

import (
	"testing"

	"github.com/justboereh/gofetch"
)

func TestBasicFunctions(t *testing.T) {
	gofetch.Fetch("https://randomuser.me/api/", gofetch.FetchOptions{})
}
