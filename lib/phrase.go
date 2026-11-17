package lib

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
)

func ReadData() []string {
	data, err := os.ReadFile("data/bomdia.json")
	if err != nil {
		fmt.Printf("Unable to read file\n")
		return nil
	}

	var phrases []string
	if err := json.Unmarshal(data, &phrases); err != nil {
		fmt.Printf("Unable to parse file\n")
		return nil
	}

	return phrases
}

func GetRandomPhrase() string {
	phrases := ReadData()
	if phrases == nil {
		return "Não foi possível encontrar dados :c"
	}
	return phrases[rand.IntN(len(phrases))]
}
