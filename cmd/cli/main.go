package main

import (
	"fmt"
	"github.com/rozdolsky33/audiofile/cmd/cli/command"
	"github.com/rozdolsky33/audiofile/internal/interfaces"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		command.NewGetCommand(client),
		command.NewUploadCommand(client),
		command.NewListCommand(client),
	}
	parser := command.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
		os.Exit(1)
	}
}
