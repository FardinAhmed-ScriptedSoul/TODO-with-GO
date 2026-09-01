package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return fmt.Errorf("invalid index: %d", index)
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	if t[index].Completed {
		t[index].Completed = false
		t[index].CompletedAt = nil
		return nil
	}

	now := time.Now()
	t[index].Completed = true
	t[index].CompletedAt = &now
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for idx, todo := range *todos {
		completed := "❌"
		completedAt := ""
		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
			}
		}

		tbl.AddRow(strconv.Itoa(idx), todo.Title, completed, todo.CreatedAt.Format("2006-01-02 15:04:05"), completedAt)
	}

	tbl.Render()
}
