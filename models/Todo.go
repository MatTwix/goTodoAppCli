package models

import (
	"encoding/json"
	"os"
	"slices"
	"time"

	"github.com/MatTwix/GoTodoAppCli/iternal"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Done      bool      `json:"done"`
}

type TodoList struct {
	Tasks []Todo
}

func (t *TodoList) AddTask(title string) {
	id := len(t.Tasks) + 1
	t.Tasks = append(t.Tasks, Todo{ID: id, Title: title, CreatedAt: time.Now(), Done: false})
}

func (t *TodoList) MarkDone(id int) error {
	flag := false
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Done = true
			flag = true
			break
		}
	}

	if !flag {
		return &iternal.DoesNotExistErr{}
	}
	return nil
}

func (t *TodoList) DeleteTask(id int) error {
	flag := false
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = slices.Delete(t.Tasks, i, i+1)
			flag = true
			break
		}
	}

	if !flag {
		return &iternal.DoesNotExistErr{}
	}
	return nil
}

func (t *TodoList) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(t.Tasks)
}

func (t *TodoList) ReadFromFile(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		newFile, createErr := os.Create(filename)
		if createErr != nil {
			return createErr
		}
		defer newFile.Close()
		t.Tasks = []Todo{}
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&t.Tasks)
}
