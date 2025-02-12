package main

import (
	"net/http"

	"github.com/jancewicz/shellspy"
)

func (app *application) terminalHandler(w http.ResponseWriter, r *http.Request) {
	//  read users data ex. "ls"
	var userInput string
	//  extract json and apply it for exec cmd logic
	cmd, err := shellspy.CommandFromInput(userInput)
	if err != nil {
		app.internalSeverError(w, r, err)
		return
	}
	shellspy.HandleCommand(cmd, file)
	//  after reading input send response
}
