package iternal

import "log"

type DoesNotExistErr struct{}

func (d *DoesNotExistErr) Error() string {
	return "Task does not exist in todo list"
}

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}
