package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

type ToolDefinition struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	InputSchema anthropic.ToolInputSchemaParam `json:"input_schema"`
	Function    func(input json.RawMessage) (string, error)
}

type ClaudeAgent struct {
	client          *anthropic.Client
	getUserMessage  func() (string, bool)
	toolsDefinition map[string]ToolDefinition
	toolsUnionParam []anthropic.ToolUnionParam
}

func NewClaudeAgent(client *anthropic.Client, getUserMessage func() (string, bool)) *ClaudeAgent {
	return &ClaudeAgent{
		client:          client,
		getUserMessage:  getUserMessage,
		toolsDefinition: make(map[string]ToolDefinition, 0),
		toolsUnionParam: []anthropic.ToolUnionParam{},
	}
}

func (ca *ClaudeAgent) RegisterTool(tool ToolDefinition) {
	ca.addToolDefinition(tool)
	ca.addToolUnionParam(anthropic.ToolUnionParam{
			OfTool: &anthropic.ToolParam{
				Name:        tool.Name,
				Description: anthropic.String(tool.Description),
				InputSchema: tool.InputSchema,
			},
		})
}

func (ca *ClaudeAgent) addToolDefinition(def ToolDefinition) {
	if _, exists := ca.toolsDefinition[def.Name]; exists {
		fmt.Printf("Warning: Tool %s is already registered, overwriting.\n", def.Name)
	}

	ca.toolsDefinition[def.Name] = def
}

func (ca *ClaudeAgent) addToolUnionParam(param anthropic.ToolUnionParam) {
	ca.toolsUnionParam = append(ca.toolsUnionParam, param)
}

func (ca *ClaudeAgent) Run(ctx context.Context) error {
	conversation := []anthropic.MessageParam{}

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	readUserInput := true
	for {
		if readUserInput {
			fmt.Print("\u001b[94mYou\u001b[0m: ")
			userInput, ok := ca.getUserMessage()
			if !ok {
				break
			}

			userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
			conversation = append(conversation, userMessage)
		}

		message, err := ca.runInference(ctx, conversation)
		if err != nil {
			return err
		}
		conversation = append(conversation, message.ToParam())

		toolResults := []anthropic.ContentBlockParamUnion{}
		for _, content := range message.Content {
			switch content.Type {
			case "text":
				fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", content.Text)
			case "tool_use":
				result := ca.executeTool(content.ID, content.Name, content.Input)
				toolResults = append(toolResults, result)
			}
		}
		if len(toolResults) == 0 {
			readUserInput = true
			continue
		}
		readUserInput = false
		conversation = append(conversation, anthropic.NewUserMessage(toolResults...))
	}

	return nil
}

func (ca *ClaudeAgent) runInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error) {
	message, err := ca.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_7SonnetLatest,
		MaxTokens: int64(1024),
		Messages:  conversation,
		Tools:     ca.toolsUnionParam,
	})

	return message, err
}

func (ca *ClaudeAgent) executeTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {
	toolDef, exist := ca.toolsDefinition[name]; 
	if !exist {
		return anthropic.NewToolResultBlock(id, fmt.Sprintf("tool %s not found", name), true)
	}

	fmt.Printf("\u001b[92mtool\u001b[0m: %s(%s)\n", name, input)
	response, err := toolDef.Function(input)
	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}
	return anthropic.NewToolResultBlock(id, response, false)
}
