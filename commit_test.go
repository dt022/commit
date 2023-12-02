package main

import (
	"strings"
	"testing"
)

func TestCommitMsg(t *testing.T) {
	args := []string{"\"hello boi\""}
	expected := "hello boi"

	commitMsg := strings.Join(args, " ")

	if commitMsg != expected {
		t.Errorf("Expected commitMsg to be %s, but got %s", expected, commitMsg)
	}
}
