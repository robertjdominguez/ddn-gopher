package utils

import (
	"testing"
)

func TestCheckForUser(t *testing.T) {
	t.Run("return an error if no username is passed", func(t *testing.T) {
		_, got := CheckForUser("")
		want := ErrEmptyUsername
		if got != want {
			t.Errorf("got %v, want %v", got, ErrEmptyUsername)
		}
	})
}
