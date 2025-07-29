package pokemon

import (
	"encoding/json"
	"errors"

	"github.com/pararang/emcp/apis"
	"github.com/pararang/emcp/claude"
)

type GetDetailInput struct {
	KeyIdentifier   string `json:"key_identifier" jsonschema_description:"The identifier of the Pokemon to get details for. This can be a name or an ID."`
	ValueIdentifier string `json:"value_identifier" jsonschema_description:"The value of the identifier to get details for. This can be a value of name or an ID."`
}

var GetDetailDefinition = claude.ToolDefinition{
	Name:        "get_pokemon_detail",
	Description: "Get details of a Pokemon by its name or ID.",
	InputSchema: claude.GenerateSchema[GetDetailInput](),
	Function:    GetDetail,
}

func GetDetail(input json.RawMessage) (string, error) {
	var getDetailInput GetDetailInput
	if err := json.Unmarshal(input, &getDetailInput); err != nil {
		return "", err
	}

	if getDetailInput.KeyIdentifier == "" || getDetailInput.ValueIdentifier == "" {
		return "", errors.New("key_identifier and value_identifier must be provided")
	}

	detail, err := apis.GetPokeDetails(getDetailInput.ValueIdentifier)
	if err != nil {
		return "", err
	}

	result := struct {
		Param  GetDetailInput `json:"param"`
		Detail any            `json:"detail"`
	}{
		Param:  getDetailInput,
		Detail: detail,
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(resultJSON), nil
}
