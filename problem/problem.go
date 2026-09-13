package problem

import "fmt"

type Problem struct {
	Platforms       Platforms
	ID              string
	Title           string
	Slug            string
	Classification  Classification
	Collections     []string
	Solutions       []Solution
	Implementations []Implementation
	SchemaVersion   int
	Difficulty      Difficulty
}

type Difficulty int

const (
	DifficultyEasy Difficulty = iota
	DifficultyMedium
	DifficultyHard
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyEasy:
		return "Easy"
	case DifficultyMedium:
		return "Medium"
	case DifficultyHard:
		return "Hard"
	default:
		return "Unknown"
	}
}

func (d Difficulty) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

var ErrInvalidDifficulty = fmt.Errorf("invalid difficulty")

func (d *Difficulty) UnmarshalText(text []byte) error {
	switch string(text) {
	case "Easy":
		*d = DifficultyEasy
	case "Medium":
		*d = DifficultyMedium
	case "Hard":
		*d = DifficultyHard
	default:
		return fmt.Errorf("%w: %s", ErrInvalidDifficulty, text)
	}
	return nil
}


type Platforms map[PlatformId]PlatformFields

type PlatformId string

type PlatformFields map[string]any

type Classification struct {
	Patterns       []string
	DataStructures []string
	Algorithms     []string
	Categories     []string
	Tags           []string
}

type Solution struct {
	Id          string
	Name        string
	Description string
	Complexity  Complexity
}

type Complexity struct {
	Time  string
	Space string
}

type Implementation struct {
	SolutionID string
	Language   string
	Source     string
	Tags       []string
}
