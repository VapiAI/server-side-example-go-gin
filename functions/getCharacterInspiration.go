package functions

import (
	"fmt"
	"log"
)

type GetCharacterInspirationParams struct {
	Inspiration string `json:"inspiration"`
}

type Response struct {
	Result                 string `json:"result"`
	ForwardToClientEnabled bool   `json:"forwardToClientEnabled"`
}

func GetCharacterInspiration(params GetCharacterInspirationParams) Response {
	fallbackResponse := Response{
		Result: "Sorry, I am dealing with a technical issue at the moment, perhaps because of heightened user traffic. Come back later and we can try this again. Apologies for that.",
	}

	if params.Inspiration != "" {
		// Here we should load data and query it, but for now we will just return a placeholder response
		response := Response{
			Result:                 "This is a placeholder response for the getCharacterInspiration function. It should be replaced with the actual implementation.",
			ForwardToClientEnabled: true,
		}

		return response
	} else {
		return fallbackResponse
	}
}

func HandleGetCharacterInspiration(params GetCharacterInspirationParams) {
	response := GetCharacterInspiration(params)
	if response.ForwardToClientEnabled {
		fmt.Println(response.Result)
	} else {
		log.Println("Error: ", response.Result)
	}
}
