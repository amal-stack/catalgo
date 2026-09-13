package main

import (
	"fmt"
	"log"

	"github.com/amal-stack/catalgo/cmd"
	"github.com/amal-stack/catalgo/problem"
	"github.com/goccy/go-yaml"
)

func main() {
	cmd.Execute()

	model := problem.Problem{
		SchemaVersion: 1,
		ID:            "two-sum",
		Title:         "Two Sum",
		Slug:          "two-sum",
		Difficulty:    problem.DifficultyEasy,
		Platforms: problem.Platforms{
			"leetcode": {
				"id":  "1",
			},
		},
		Collections: []string{"neetcode150"},
		Classification: problem.Classification{
			Patterns:       []string{"Two Pointers"},
			DataStructures: []string{"Array", "Hash Table"},
			Algorithms:     []string{"Brute Force", "Hashing"},
			Categories:     []string{"Array", "Hash Table"},
			Tags:           []string{"Array", "Hash Table"},
		},
	}

	data, err := yaml.Marshal(&model)

	if err != nil {
		log.Fatalf("Failed to marshal model: %v", err)
	}

	fmt.Println("====YAML====")
	fmt.Println(string(data))

	var hydrated problem.Problem

	err = yaml.Unmarshal(data, &hydrated)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== HYDRATED ===")
	fmt.Printf("%+v\n", hydrated)
}
