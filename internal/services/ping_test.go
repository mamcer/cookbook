package services

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetMessage(t *testing.T) {
	t.Run("testing ping pong message", func(t *testing.T) {
		got := (&PingService{}).GetMessage()
		want := "pong"

		assertCorrectMessage(t, got, want)
	})
}
