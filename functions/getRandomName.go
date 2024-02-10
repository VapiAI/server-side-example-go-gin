package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var nats = []string{
	"AU",
	"CA",
	"FR",
	"IN",
	"IR",
	"MX",
	"NL",
	"NO",
	"NZ",
	"RS",
	"TR",
	"US",
}

type NameParams struct {
	Gender string `json:"gender,omitempty"`
	Nat    string `json:"nat,omitempty"`
}

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Result struct {
	Name Name `json:"name"`
}

type ResponseData struct {
	Results []Result `json:"results"`
}

func GetRandomName(params NameParams) (string, error) {
	nat := params.Nat
	if nat != "" && !contains(nats, strings.ToUpper(nat)) {
		nat = nats[rand.Intn(len(nats))]
	}

	queryParams := fmt.Sprintf("gender=%s&nat=%s", params.Gender, nat)
	response, err := http.Get(fmt.Sprintf("https://randomuser.me/api/?%s", queryParams))
	if err != nil {
		return "", errors.New("Error fetching random name")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error! status: %d", response.StatusCode)
	}

	data, _ := io.ReadAll(response.Body)
	var responseData ResponseData
	json.Unmarshal(data, &responseData)

	name := responseData.Results[0].Name
	return name.First + " " + name.Last, nil
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
