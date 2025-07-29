package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/pararang/emcp/claude"
	"github.com/pararang/emcp/tools"
	pokeTools "github.com/pararang/emcp/tools/pokemon"
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
		pokeTools.GetAbilityDetailDefinition,
	)

	err := claude.Run(context.TODO())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

