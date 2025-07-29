package pokemon

import (
	"encoding/json"
	"errors"

	"github.com/pararang/code-editing-agent/apis"
	"github.com/pararang/code-editing-agent/claude"
)

type GetAbilityDetailInput struct {
	KeyIdentifier   string `json:"key_identifier" jsonschema_description:"The identifier of the pokemon ability to get details for. This can be a name or an ID."`
	ValueIdentifier string `json:"value_identifier" jsonschema_description:"The value of the identifier to get ability details for. This can be value of a name or an ID."`
}

var GetAbilityDetailDefinition = claude.ToolDefinition{
	Name:        "get_ability_detail",
	Description: "Get ability details by ability name or ID.",
	InputSchema: claude.GenerateSchema[GetDetailInput](),
	Function:    GetAbilityDetail,
}

func GetAbilityDetail(input json.RawMessage) (string, error) {
	var getAbilityDetailInput GetAbilityDetailInput
	if err := json.Unmarshal(input, &getAbilityDetailInput); err != nil {
		return "", err
	}

	if getAbilityDetailInput.KeyIdentifier == "" || getAbilityDetailInput.ValueIdentifier == "" {
		return "", errors.New("key_identifier and value_identifier must be provided")
	}

	detail, err := apis.GetPokeAbilityDetails(getAbilityDetailInput.ValueIdentifier)
	if err != nil {
		return "", err
	}

	result := struct {
		Param  GetAbilityDetailInput `json:"param"`
		Detail any                   `json:"detail"`
	}{
		Param:  getAbilityDetailInput,
		Detail: detail,
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(resultJSON), nil
}
