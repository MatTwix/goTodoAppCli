package models

import (
	"encoding/csv"
	"os"
	"slices"
	"strconv"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

type TodoList struct {
	Tasks []Todo
}

func (t *TodoList) AddTask(title string) {
	id := len(t.Tasks) + 1
	t.Tasks = append(t.Tasks, Todo{ID: id, Title: title, Done: false})
}

func (t *TodoList) MarkDone(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Done = true
			break
		}
	}
}

func (t *TodoList) DeleteTask(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = slices.Delete(t.Tasks, i, i+1)
			break
		}
	}
}

func (t *TodoList) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write([]string{"ID", "Title", "Done"}); err != nil {
		return err
	}

	for _, task := range t.Tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Title,
			strconv.FormatBool(task.Done),
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func (t *TodoList) ReadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	t.Tasks = []Todo{}
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		done, err := strconv.ParseBool(record[2])
		if err != nil {
			return err
		}

		t.Tasks = append(t.Tasks, Todo{
			ID:    id,
			Title: record[1],
			Done:  done,
		})
	}

	return nil
}
