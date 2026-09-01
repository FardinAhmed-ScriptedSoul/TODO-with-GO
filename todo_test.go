package main


import (
	"testing"
	"time"
)

func TestToggleResetsCompletedAt(t *testing.T) {
	todos := Todos{{
		Title:     "Write release notes",
		CreatedAt: time.Now(),
	}}

	if err := todos.toggle(0); err != nil {
		t.Fatalf("toggle on first call failed: %v", err)
	}
	if !todos[0].Completed {
		t.Fatal("expected task to become completed")
	}
	if todos[0].CompletedAt == nil {
		t.Fatal("expected completedAt to be set")
	}

	if err := todos.toggle(0); err != nil {
		t.Fatalf("toggle on second call failed: %v", err)
	}
	if todos[0].Completed {
		t.Fatal("expected task to become incomplete")
	}
	if todos[0].CompletedAt != nil {
		t.Fatal("expected completedAt to be cleared")
	}
}
