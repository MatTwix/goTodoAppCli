package models

import (
	"io"
)

func LoadTodoList(filename string) (*TodoList, error) {
	todoList := &TodoList{}
	err := todoList.ReadFromFile(filename)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return todoList, err
}
