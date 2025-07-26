package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/pararang/code-editing-agent/apis"
	"github.com/pararang/code-editing-agent/tools"
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

	mcp := apis.NewAgent(&client, getUserMessage)

	mcp.RegisterTool(tools.ReadFileDefinition)

	err := mcp.Run(context.TODO())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
