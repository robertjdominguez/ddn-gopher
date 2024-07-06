package utils

import (
	"errors"
	"testing"
)

func TestCheckForUser(t *testing.T) {
	t.Run("return an error if no username is passed", func(t *testing.T) {
		_, got := CheckForUser("")
		want := errors.New("username cannot be an empty string")
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
