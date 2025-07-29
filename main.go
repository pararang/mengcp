package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/pararang/code-editing-agent/claude"
	"github.com/pararang/code-editing-agent/tools"
	pokeTools "github.com/pararang/code-editing-agent/tools/pokemon"
)

func main() {
	client := anthropic.NewClient()

	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	claude := claude.NewClaudeAgent(&client, getUserMessage)

	claude.RegisterTools(
		tools.ReadFileDefinition, 
		tools.ListFilesDefinition,
		tools.EditFileDefinition,
		pokeTools.GetDetailDefinition,
	)

	err := claude.Run(context.TODO())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

