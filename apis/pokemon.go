package apis

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const pokeApiHost = "https://pokeapi.co/api/v2/"


type PokemonDetails struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Types          []string `json:"types"`
	Abilities      []string `json:"abilities"`
	BaseExperience int      `json:"base_experience"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
}

func GetPokeDetails(pokemonNameOrID string) (PokemonDetails, error) {
	url := fmt.Sprintf("%spokemon/%s", pokeApiHost, pokemonNameOrID)
	data, err := fetcher(url)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("couldn't get details - '%s': %v", pokemonNameOrID, err)
	}

	var caserTitle = cases.Title(language.Indonesian)

	var types []string
	if typesData, ok := data["types"].([]any); ok {
		for _, t := range typesData {
			if typeMap, ok := t.(map[string]any); ok {
				if typeInfo, ok := typeMap["type"].(map[string]any); ok {
					if name, ok := typeInfo["name"].(string); ok {
						types = append(types, caserTitle.String(name))
					}
				}
			}
		}
	}

	var abilities []string
	if abilitiesData, ok := data["abilities"].([]any); ok {
		for _, a := range abilitiesData {
			if abilityMap, ok := a.(map[string]any); ok {
				if abilityInfo, ok := abilityMap["ability"].(map[string]any); ok {
					if name, ok := abilityInfo["name"].(string); ok {
						abilities = append(abilities, name)
					}
				}
			}
		}
	}

	result := PokemonDetails{
		ID:             int(data["id"].(float64)),
		Name:           caserTitle.String(data["name"].(string)),
		Types:          types,
		Abilities:      abilities,
		BaseExperience: int(data["base_experience"].(float64)),
		Height:         int(data["height"].(float64)),
		Weight:         int(data["weight"].(float64)),
	}

	return result, nil
}

func GetPokeAbilityDetails(abilityNameOrID string) (map[string]any, error) {
	url := fmt.Sprintf("%sability/%s", pokeApiHost, abilityNameOrID)
	data, err := fetcher(url)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve detailed information for '%s': %v", abilityNameOrID, err)
	}
	return data, nil
}
