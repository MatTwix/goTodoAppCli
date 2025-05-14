package cmd

import "log"

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}
