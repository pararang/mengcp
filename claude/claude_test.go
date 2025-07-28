package claude

import (
	"testing"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestClaudeAgent_addToolUnionParam(t *testing.T) {
	ca := NewClaudeAgent(nil, nil)

	ca.addToolUnionParam(anthropic.ToolUnionParam{
		OfTool: &anthropic.ToolParam{
			Name: "test_tool",
		},
	})

	assert.Equal(t, "test_tool", ca.toolsUnionParam[0].OfTool.Name)
}

func TestClaudeAgent_addToolDefinition(t *testing.T) {
	ca := NewClaudeAgent(nil, nil)

	ca.addToolDefinition(ToolDefinition{
		Name: "test_tool",
	})

	assert.NotNil(t, ca.toolsDefinition["test_tool"])
	assert.Equal(t, "test_tool", ca.toolsDefinition["test_tool"].Name)
}
