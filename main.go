package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type InfoObj struct {
	ResourceType         string          `json:"resourceType"`
	ID                   string          `json:"id"`
	Text                 TextObj         `json:"text"`
	Identifier           IdentifierObj   `json:"identifier"`
	Active               bool            `json:"active"`
	Name                 []NameObj       `json:"name"`
	Gender               string          `json:"gender"`
	BirthDate            string          `json:"birthDate"`
	ManagingOrganization OrganizationObj `json:"managingOrganization"`
	Conditions           []string        `json:"conditions"`
}

type TextObj struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

type IdentifierObj struct {
	Use    string  `json:"use"`
	Type   TypeObj `json:"type"`
	System string  `json:"system"`
	Value  string  `json:"value"`
}

type TypeObj struct {
	Coding []CodingObj `json:"coding"`
}

type CodingObj struct {
	System string `json:"system"`
	Code   string `json:"MR"`
}

type NameObj struct {
	Family []string `json:"family"`
	Given  []string `json:"given"`
}

type OrganizationObj struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

func main() {
	http.HandleFunc("/", viewIndex)
	http.HandleFunc("/getData", getData)
	http.ListenAndServe(":3000", nil)
}

func viewIndex(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("index.html")

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var obj InfoObj
	json.Unmarshal(raw, &obj)

	bytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := json.NewEncoder(w).Encode(string(bytes)); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
